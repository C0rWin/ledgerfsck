package main

import (
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric/common/flogging"
	"github.com/hyperledger/fabric/common/ledger/blkstorage"
	"github.com/hyperledger/fabric/common/ledger/blkstorage/fsblkstorage"
	"github.com/hyperledger/fabric/common/localmsp"
	"github.com/hyperledger/fabric/core/peer"
	"github.com/hyperledger/fabric/msp"
	"github.com/hyperledger/fabric/msp/mgmt"
	common2 "github.com/hyperledger/fabric/peer/common"
	peergossip "github.com/hyperledger/fabric/peer/gossip"
	"github.com/hyperledger/fabric/protos/common"
	logging "github.com/op/go-logging"
)

var logger = flogging.MustGetLogger("common/tools/ledgerfsck")

func main() {
	logging.SetLevel(logging.INFO, "")

	//'channelID' will be a parameter
	channelID := "testorgschannel1"

	blkStoragePath := "/tmp/peer0ledger/chains"

	var mspMgrConfigDir = "/root/git/src/github.com/hyperledger/fabric/common/tools/cryptogen/crypto-config/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/msp"
	var mspID = "PeerOrg1"
	var mspType = msp.ProviderTypeToString(msp.FABRIC)

	err := common2.InitCrypto(mspMgrConfigDir, mspID, mspType)
	if err != nil { // Handle errors reading the config file
		panic(fmt.Sprintf("Cannot run client because %s", err.Error()))
	}

	attrsToIndex := []blkstorage.IndexableAttr{
		blkstorage.IndexableAttrBlockHash,
		blkstorage.IndexableAttrBlockNum,
		blkstorage.IndexableAttrTxID,
		blkstorage.IndexableAttrBlockNumTranNum,
		blkstorage.IndexableAttrBlockTxID,
		blkstorage.IndexableAttrTxValidationCode,
	}

	indexConfig := &blkstorage.IndexConfig{AttrsToIndex: attrsToIndex}

	conf := fsblkstorage.NewConf(blkStoragePath, 0)

	messageCryptoService := peergossip.NewMCS(
		peer.NewChannelPolicyManagerGetter(),
		localmsp.NewSigner(),
		mgmt.NewDeserializersManager())

	provider := fsblkstorage.NewProvider(conf, indexConfig).(*fsblkstorage.FsBlockstoreProvider)

	store, _ := provider.OpenBlockStore(channelID)
	logger.Info("store is %s", store)

	bcInfo, err := store.GetBlockchainInfo()
	if err != nil {
		return
	}

	logger.Info("blockchain height is", bcInfo.Height)

	// We may not need line 72-78, dirty approach
	/*
		gensisBlock, err := configtxtest.MakeGenesisBlock(channelID)
		ledger := ledgermgmt.GetExistingLedger(channelID)
		if err == nil {
			logger.Info("ledger is nil")
		} else {
			peer.CreateChain(channelID, ledger, gensisBlock)
		}
	*/

	var block *common.Block
	var i uint64
	for i = 0; i < bcInfo.Height; i++ {
		block, _ = store.RetrieveBlockByNumber(i)
		logger.Info(block.Header.Number)

		marshaledBlock, err := proto.Marshal(block)
		if err != nil {
			return
		}

		err = messageCryptoService.VerifyBlock([]byte(channelID), i, marshaledBlock)

		if err != nil {
			logger.Info("block ERROR", err.Error())
		} else {
			logger.Info("block correct")
		}

	}
}
