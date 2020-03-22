package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric/common/channelconfig"
	"github.com/hyperledger/fabric/common/flogging"
	"github.com/hyperledger/fabric/common/localmsp"
	"github.com/hyperledger/fabric/common/metrics/disabled"
	"github.com/hyperledger/fabric/common/policies"
	"github.com/hyperledger/fabric/core/chaincode/platforms"
	"github.com/hyperledger/fabric/core/chaincode/platforms/car"
	"github.com/hyperledger/fabric/core/chaincode/platforms/golang"
	"github.com/hyperledger/fabric/core/chaincode/platforms/java"
	"github.com/hyperledger/fabric/core/chaincode/platforms/node"
	"github.com/hyperledger/fabric/core/ledger"
	"github.com/hyperledger/fabric/core/ledger/ledgermgmt"
	"github.com/hyperledger/fabric/core/peer"
	"github.com/hyperledger/fabric/core/scc/lscc"
	gossipCommon "github.com/hyperledger/fabric/gossip/common"
	"github.com/hyperledger/fabric/msp/mgmt"
	"github.com/hyperledger/fabric/peer/common"
	"github.com/hyperledger/fabric/peer/gossip"
	pb "github.com/hyperledger/fabric/protos/common"
	"github.com/hyperledger/fabric/protos/utils"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

var logger = flogging.MustGetLogger("ledgerfsck")

type ledgerFsck struct {
	channelName   string
	mspConfigPath string
	mspID         string
	mspType       string

	ledger ledger.PeerLedger
	bundle *channelconfig.Bundle
}

func (fsck *ledgerFsck) Manager(channelID string) (policies.Manager, bool) {
	return fsck.bundle.PolicyManager(), true
}

// Initialize
func (fsck *ledgerFsck) Initialize() error {
	// Initialize viper configuration
	viper.SetEnvPrefix("core")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := common.InitConfig("core")
	if err != nil {
		logger.Errorf("failed to initialize configuration, because of %s", err)
		return err
	}
	return nil
}

// ReadConfiguration read configuration parameters
func (fsck *ledgerFsck) ReadConfiguration() error {
	// Read configuration parameters
	flag.StringVar(&fsck.channelName, "channelName", "testChannel", "channel name to check the integrity")
	flag.StringVar(&fsck.mspConfigPath, "mspPath", "", "path to the msp folder")
	flag.StringVar(&fsck.mspID, "mspID", "", "the MSP identity of the organization")
	flag.StringVar(&fsck.mspType, "mspType", "bccsp", "the type of the MSP provider, default bccsp")
	flag.Parse()

	if fsck.mspConfigPath == "" {
		errMsg := "MSP folder not configured"
		logger.Error(errMsg)
		return errors.New(errMsg)
	}

	if fsck.mspID == "" {
		errMsg := "MSPID was not provided"
		logger.Error(errMsg)
		return errors.New(errMsg)
	}

	logger.Infof("channel name = %s", fsck.channelName)
	logger.Infof("MSP folder path = %s", fsck.mspConfigPath)
	logger.Infof("MSPID = %s", fsck.mspID)
	logger.Infof("MSP type = %s", fsck.mspType)
	return nil
}

// InitCrypto
func (fsck *ledgerFsck) InitCrypto() error {
	// Next need to init MSP
	err := common.InitCrypto(fsck.mspConfigPath, fsck.mspID, fsck.mspType)
	if err != nil {
		logger.Errorf("failed to initialize MSP related configuration, failure %s", err)
		return err
	}
	return nil
}

// OpenLedger
func (fsck *ledgerFsck) OpenLedger() error {
	// Initialize ledger management
	ledgermgmt.Initialize(&ledgermgmt.Initializer{
		CustomTxProcessors: peer.ConfigTxProcessors,
		PlatformRegistry: platforms.NewRegistry(
			&golang.Platform{},
			&node.Platform{},
			&java.Platform{},
			&car.Platform{},
		),
		DeployedChaincodeInfoProvider: &lscc.DeployedCCInfoProvider{},
		MetricsProvider:               &disabled.Provider{},
	})
	ledgerIds, err := ledgermgmt.GetLedgerIDs()
	if err != nil {
		errMsg := fmt.Sprintf("failed to read ledger, because of %s", err)
		logger.Errorf(errMsg)
		return errors.New(errMsg)
	}

	// Check whenever channel name has corresponding ledger
	var found = false
	for _, name := range ledgerIds {
		if name == fsck.channelName {
			found = true
		}
	}

	if !found {
		errMsg := fmt.Sprintf("there is no ledger corresponding to the provided channel name %s. Exiting...", fsck.channelName)
		logger.Errorf(errMsg)
		return errors.New(errMsg)
	}

	if fsck.ledger, err = ledgermgmt.OpenLedger(fsck.channelName); err != nil {
		errMsg := fmt.Sprintf("failed to open ledger %s, because of the %s", fsck.channelName, err)
		logger.Errorf(errMsg)
		return errors.New(errMsg)
	}
	return nil
}

// GetLatestChannelConfigBundle
func (fsck *ledgerFsck) GetLatestChannelConfigBundle() error {
	var cb *pb.Block
	var err error
	if cb, err = getCurrConfigBlockFromLedger(fsck.ledger); err != nil {
		logger.Warningf("Failed to find config block on ledger %s(%s)", fsck.channelName, err)
		return err
	}

	qe, err := fsck.ledger.NewQueryExecutor()
	defer qe.Done()
	if err != nil {
		logger.Errorf("failed to obtain query executor, error is %s", err)
		return err
	}

	logger.Info("reading configuration from state DB")
	confBytes, err := qe.GetState("", "resourcesconfigtx.CHANNEL_CONFIG_KEY")
	if err != nil {
		logger.Errorf("failed to read channel config, error %s", err)
		return err
	}
	conf := &pb.Config{}
	err = proto.Unmarshal(confBytes, conf)
	if err != nil {
		logger.Errorf("could not read configuration, due to %s", err)
		return err
	}

	if conf != nil {
		logger.Info("initialize channel config bundle")
		fsck.bundle, err = channelconfig.NewBundle(fsck.channelName, conf)
		if err != nil {
			return err
		}
	} else {
		// Config was only stored in the statedb starting with v1.1 binaries
		// so if the config is not found there, extract it manually from the config block
		logger.Info("configuration wasn't stored in state DB retrieving config envelope from ledger")
		envelopeConfig, err := utils.ExtractEnvelope(cb, 0)
		if err != nil {
			return err
		}

		logger.Info("initialize channel config bundle from config transaction")
		fsck.bundle, err = channelconfig.NewBundleFromEnvelope(envelopeConfig)
		if err != nil {
			return err
		}
	}

	capabilitiesSupportedOrPanic(fsck.bundle)

	channelconfig.LogSanityChecks(fsck.bundle)

	return nil
}

func (fsck *ledgerFsck) Verify() {
	blockchainInfo, err := fsck.ledger.GetBlockchainInfo()
	if err != nil {
		logger.Errorf("could not obtain blockchain information "+
			"channel name %s, due to %s", fsck.channelName, err)
		os.Exit(-1)
	}

	logger.Infof("ledger height of channel %s, is %d\n", fsck.channelName, blockchainInfo.Height)

	mcs := gossip.NewMCS(
		fsck,
		localmsp.NewSigner(),
		mgmt.NewDeserializersManager())

	block, err := fsck.ledger.GetBlockByNumber(uint64(0))
	if err != nil {
		logger.Errorf("failed to read genesis block number, with error", err)
		os.Exit(-1)
	}

	// Get hash of genesis block
	prevHash := block.Header.Hash()

	// complete full scan and check over ledger blocks
	for blockIndex := uint64(1); blockIndex < blockchainInfo.Height; blockIndex++ {
		block, err := fsck.ledger.GetBlockByNumber(blockIndex)
		if err != nil {
			logger.Errorf("failed to read block number %d from ledger, with error", blockIndex, err)
			os.Exit(-1)
		}

		if !bytes.Equal(prevHash, block.Header.PreviousHash) {
			logger.Errorf("block number [%d]: hash comparison has failed, previous block hash %x doesn't"+
				" equal to hash claimed within block header %x", blockIndex, prevHash, block.Header.PreviousHash)
			os.Exit(-1)
		} else {
			logger.Infof("block number [%d]: previous hash matched", blockIndex)
		}

		signedBlock, err := proto.Marshal(block)
		if err != nil {
			logger.Errorf("failed marshaling block, due to", err)
			os.Exit(-1)
		}

		if err := mcs.VerifyBlock(gossipCommon.ChainID(fsck.channelName), block.Header.Number, signedBlock); err != nil {
			logger.Errorf("failed to verify block with sequence number %d. %s", blockIndex, err)
			os.Exit(-1)
		} else {
			logger.Infof("Block [seq = %d], hash = [%x], previous hash = [%x], VERIFICATION PASSED",
				blockIndex, block.Header.Hash(), block.Header.PreviousHash)
		}
		prevHash = block.Header.Hash()
	}
}

func main() {
	fsck := &ledgerFsck{}
	// Initialize configuration
	if err := fsck.Initialize(); err != nil {
		os.Exit(-1)
	}
	// Read configuration parameters
	if err := fsck.ReadConfiguration(); err != nil {
		os.Exit(-1)
	}
	// Init crypto & MSP
	if err := fsck.InitCrypto(); err != nil {
		os.Exit(-1)
	}
	// OpenLedger
	if err := fsck.OpenLedger(); err != nil {
		os.Exit(-1)
	}
	// GetLatestChannelConfigBundle
	if err := fsck.GetLatestChannelConfigBundle(); err != nil {
		os.Exit(-1)
	}

	fsck.Verify()
}

// getCurrConfigBlockFromLedger read latest configuratoin block from the ledger
func getCurrConfigBlockFromLedger(ledger ledger.PeerLedger) (*pb.Block, error) {
	logger.Debugf("Getting config block")

	// get last block.  Last block number is Height-1
	blockchainInfo, err := ledger.GetBlockchainInfo()
	if err != nil {
		return nil, err
	}
	lastBlock, err := ledger.GetBlockByNumber(blockchainInfo.Height - 1)
	if err != nil {
		return nil, err
	}

	// get most recent config block location from last block metadata
	configBlockIndex, err := utils.GetLastConfigIndexFromBlock(lastBlock)
	if err != nil {
		return nil, err
	}

	// get most recent config block
	configBlock, err := ledger.GetBlockByNumber(configBlockIndex)
	if err != nil {
		return nil, err
	}

	logger.Debugf("Got config block[%d]", configBlockIndex)
	return configBlock, nil
}

func capabilitiesSupportedOrPanic(res channelconfig.Resources) {
	ac, ok := res.ApplicationConfig()
	if !ok {
		logger.Panicf("[channel %s] does not have application config so is incompatible", res.ConfigtxValidator().ChainID())
	}

	if err := ac.Capabilities().Supported(); err != nil {
		logger.Panicf("[channel %s] incompatible %s", res.ConfigtxValidator(), err)
	}

	if err := res.ChannelConfig().Capabilities().Supported(); err != nil {
		logger.Panicf("[channel %s] incompatible %s", res.ConfigtxValidator(), err)
	}
}
