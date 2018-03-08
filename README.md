# ledgerfsck - Tool to verify ledger integrity

The goal of this tool is to verify Hyperledger Fabric channel ledger. It travers ledger starting from 
the very first block till the last one, `ledgerfsck` utilizes MessageCryptoService to verify blocks
integrity. It reads channel and resource configuration from the ledger to initialize MSP structures.


##### Execution.

Following is the example of `ledgerfsck` execution:

```bash
./ledgerfsck --channelName mychannel \
             --mspID Org1MSP \
             --mspPath crypto-config/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
```

where:

- `--channelName` - the name of the channel which ledger we would like to verify
- `--mspID` - the MSP ID of the organization which owns and executes the `ledgerfsck`
- `--mspPath` - path to the MSP folder with all relevant crypto material


##### Output example.

Bellow example of the output which corresponds to the execution command above:

```bash
2018-03-08 02:57:10.894 IST [ledgerfsck] ReadConfiguration -> INFO 001 channel name = mychannel
2018-03-08 02:57:10.894 IST [ledgerfsck] ReadConfiguration -> INFO 002 MSP folder path = crypto-config/peerOrganizations/org1.example.com/users/Admin@org1.example.co
m/msp
2018-03-08 02:57:10.894 IST [ledgerfsck] ReadConfiguration -> INFO 003 MSPID = Org1MSP
2018-03-08 02:57:10.894 IST [ledgerfsck] ReadConfiguration -> INFO 004 MSP type = bccsp
2018-03-08 02:57:10.915 IST [ledgermgmt] initialize -> INFO 005 Initializing ledger mgmt
2018-03-08 02:57:10.915 IST [kvledger] NewProvider -> INFO 006 Initializing ledger provider
2018-03-08 02:57:10.933 IST [kvledger] NewProvider -> INFO 007 ledger provider Initialized
2018-03-08 02:57:10.933 IST [ledgermgmt] initialize -> INFO 008 ledger mgmt initialized
2018-03-08 02:57:10.933 IST [ledgermgmt] OpenLedger -> INFO 009 Opening ledger with id = mychannel
2018-03-08 02:57:10.935 IST [ledgermgmt] OpenLedger -> INFO 00a Opened ledger with id = mychannel
2018-03-08 02:57:10.935 IST [ledgerfsck] GetLatestChannelConfigBundle -> INFO 00b reading configuration from state DB
2018-03-08 02:57:10.935 IST [ledgerfsck] GetLatestChannelConfigBundle -> INFO 00c initialize channel config bundle
2018-03-08 02:57:10.938 IST [ledgerfsck] GetLatestResourceConfigBundle -> INFO 00d check capabilities whenever there is support for resource tree
2018-03-08 02:57:10.938 IST [ledgerfsck] GetLatestResourceConfigBundle -> INFO 00e initialize resource config bundle
2018-03-08 02:57:10.938 IST [ledgerfsck] GetLatestResourceConfigBundle -> INFO 00f Initialize MSP Manager
2018-03-08 02:57:10.938 IST [ledgerfsck] Verify -> INFO 010 ledger height of channel mychannel, is 120
2018-03-08 02:57:10.938 IST [ledgerfsck] Verify -> INFO 011 block number [1]: previous hash matched
2018-03-08 02:57:10.938 IST [msp] DeserializeIdentity -> INFO 012 Obtaining identity
2018-03-08 02:57:10.938 IST [ledgerfsck] Verify -> INFO 013 Block [seq = 1], hash = [b4f21d966705efdade2be5823d2ab8196b50bac471ef2d03f0ea3e64d3bb3b84], previous hash = [df218993357f55c6dee14ed29d43815c36bdd06a3187cd609eccc0d
0fd3624b6], VERIFICATION PASSED
2018-03-08 02:57:10.939 IST [ledgerfsck] Verify -> INFO 014 block number [2]: previous hash matched
2018-03-08 02:57:10.939 IST [ledgerfsck] Verify -> INFO 015 Block [seq = 2], hash = [bf5ad11abbd84ec32388774895e99ed0bc2d81b8917667dd0f191ab732f529b4], previous hash = [b4f21d966705efdade2be5823d2ab8196b50bac471ef2d03f0ea3e6
4d3bb3b84], VERIFICATION PASSED
2018-03-08 02:57:10.939 IST [ledgerfsck] Verify -> INFO 016 block number [3]: previous hash matched
2018-03-08 02:57:10.940 IST [ledgerfsck] Verify -> INFO 017 Block [seq = 3], hash = [3a2b1576cc7eb2f766597b43fcf00d76bed411aa9f05337b2456af7e8bc88f9a], previous hash = [bf5ad11abbd84ec32388774895e99ed0bc2d81b8917667dd0f191ab
732f529b4], VERIFICATION PASSED
2018-03-08 02:57:10.940 IST [ledgerfsck] Verify -> INFO 018 block number [4]: previous hash matched
2018-03-08 02:57:10.941 IST [ledgerfsck] Verify -> INFO 019 Block [seq = 4], hash = [8a6d7e73dc4e1825f633f157ae44bab380ef941a27e4f500b0221ac1920c0eb8], previous hash = [3a2b1576cc7eb2f766597b43fcf00d76bed411aa9f05337b2456af7
e8bc88f9a], VERIFICATION PASSED
2018-03-08 02:57:10.941 IST [ledgerfsck] Verify -> INFO 01a block number [5]: previous hash matched
2018-03-08 02:57:10.941 IST [ledgerfsck] Verify -> INFO 01b Block [seq = 5], hash = [f3dfba89c04b304ea257543613fb5c1748693da7d946827f4cc9afb6374e9c12], previous hash = [8a6d7e73dc4e1825f633f157ae44bab380ef941a27e4f500b0221ac
1920c0eb8], VERIFICATION PASSED
2018-03-08 02:57:10.941 IST [ledgerfsck] Verify -> INFO 01c block number [6]: previous hash matched
2018-03-08 02:57:10.942 IST [ledgerfsck] Verify -> INFO 01d Block [seq = 6], hash = [feedbedca119c339466bdb4ba84ec59b34926bd265599cf244b28255b62a9c4e], previous hash = [f3dfba89c04b304ea257543613fb5c1748693da7d946827f4cc9afb
6374e9c12], VERIFICATION PASSED
2018-03-08 02:57:10.942 IST [ledgerfsck] Verify -> INFO 01e block number [7]: previous hash matched
2018-03-08 02:57:10.943 IST [ledgerfsck] Verify -> INFO 01f Block [seq = 7], hash = [5e09a3da3f8e23015b1986e71ca349a6c2379b8fa0f1233e448a5696046d0e7c], previous hash = [feedbedca119c339466bdb4ba84ec59b34926bd265599cf244b2825
5b62a9c4e], VERIFICATION PASSED
2018-03-08 02:57:10.943 IST [ledgerfsck] Verify -> INFO 020 block number [8]: previous hash matched
2018-03-08 02:57:10.943 IST [ledgerfsck] Verify -> INFO 021 Block [seq = 8], hash = [c1b34f888f3caead21a59c7096d903515f7d1cfbb70274243cf1fa03440140c5], previous hash = [5e09a3da3f8e23015b1986e71ca349a6c2379b8fa0f1233e448a569
6046d0e7c], VERIFICATION PASSED
2018-03-08 02:57:10.943 IST [ledgerfsck] Verify -> INFO 022 block number [9]: previous hash matched
2018-03-08 02:57:10.944 IST [ledgerfsck] Verify -> INFO 023 Block [seq = 9], hash = [fa785f9935a7c56ad8f52d1c1cdbce080cdb469d90eb9959cdb1be8a38f74b87], previous hash = [c1b34f888f3caead21a59c7096d903515f7d1cfbb70274243cf1fa0
3440140c5], VERIFICATION PASSED
2018-03-08 02:57:10.944 IST [ledgerfsck] Verify -> INFO 024 block number [10]: previous hash matched
2018-03-08 02:57:10.945 IST [ledgerfsck] Verify -> INFO 025 Block [seq = 10], hash = [19c14185e9d1dfc781d7c4a8b1df4f5edf16181780d9e3569680a7ef4cbeb76c], previous hash = [fa785f9935a7c56ad8f52d1c1cdbce080cdb469d90eb9959cdb1be
8a38f74b87], VERIFICATION PASSED
2018-03-08 02:57:10.945 IST [ledgerfsck] Verify -> INFO 026 block number [11]: previous hash matched
2018-03-08 02:57:10.946 IST [ledgerfsck] Verify -> INFO 027 Block [seq = 11], hash = [97e529ddf2ced3cd76f67f78dafb1c524e7a75d63cdefd8449c981bbfee86233], previous hash = [19c14185e9d1dfc781d7c4a8b1df4f5edf16181780d9e3569680a7
ef4cbeb76c], VERIFICATION PASSED
2018-03-08 02:57:10.946 IST [ledgerfsck] Verify -> INFO 028 block number [12]: previous hash matched
2018-03-08 02:57:10.946 IST [ledgerfsck] Verify -> INFO 029 Block [seq = 12], hash = [b95e40bc9a4796b9035a1463f6359c7da601e0a8874374e94c7e794c13b8681f], previous hash = [97e529ddf2ced3cd76f67f78dafb1c524e7a75d63cdefd8449c981
bbfee86233], VERIFICATION PASSED
2018-03-08 02:57:10.946 IST [ledgerfsck] Verify -> INFO 02a block number [13]: previous hash matched
2018-03-08 02:57:10.947 IST [ledgerfsck] Verify -> INFO 02b Block [seq = 13], hash = [01a5ea9354af47d76ba60f03e563d34c55178c0a62d798271e01f44c99ce6585], previous hash = [b95e40bc9a4796b9035a1463f6359c7da601e0a8874374e94c7e79
4c13b8681f], VERIFICATION PASSED
2018-03-08 02:57:10.947 IST [ledgerfsck] Verify -> INFO 02c block number [14]: previous hash matched
2018-03-08 02:57:10.947 IST [ledgerfsck] Verify -> INFO 02d Block [seq = 14], hash = [385d0a5abda35f5d8abcb143d84ec55765ff5f67f7f25a87be479ad340c1b8b2], previous hash = [01a5ea9354af47d76ba60f03e563d34c55178c0a62d798271e01f4
4c99ce6585], VERIFICATION PASSED
2018-03-08 02:57:10.947 IST [ledgerfsck] Verify -> INFO 02e block number [15]: previous hash matched
2018-03-08 02:57:10.948 IST [ledgerfsck] Verify -> INFO 02f Block [seq = 15], hash = [dbd0ee5ad4e4371a362c4853c57e06a2b7f57e49672f4ec848093cb17b7c340c], previous hash = [385d0a5abda35f5d8abcb143d84ec55765ff5f67f7f25a87be479a
d340c1b8b2], VERIFICATION PASSED
2018-03-08 02:57:10.948 IST [ledgerfsck] Verify -> INFO 030 block number [16]: previous hash matched
2018-03-08 02:57:10.948 IST [ledgerfsck] Verify -> INFO 031 Block [seq = 16], hash = [6943e1c4212972b24ed015d6adec176c7095365bd18f7ecdef5d1d211c8994d0], previous hash = [dbd0ee5ad4e4371a362c4853c57e06a2b7f57e49672f4ec848093c
b17b7c340c], VERIFICATION PASSED
2018-03-08 02:57:10.949 IST [ledgerfsck] Verify -> INFO 032 block number [17]: previous hash matched
2018-03-08 02:57:10.949 IST [ledgerfsck] Verify -> INFO 033 Block [seq = 17], hash = [094600cc256a2f534362619ac05b6b58b6b7f6680fee81e5c378d89a8e3da2b5], previous hash = [6943e1c4212972b24ed015d6adec176c7095365bd18f7ecdef5d1d
211c8994d0], VERIFICATION PASSED
2018-03-08 02:57:10.949 IST [ledgerfsck] Verify -> INFO 034 block number [18]: previous hash matched
2018-03-08 02:57:10.950 IST [ledgerfsck] Verify -> INFO 035 Block [seq = 18], hash = [fa1f93ad56c761d08c7fce71f2087db4e6fd4de3648fe02bf1accf53ecfc6da8], previous hash = [094600cc256a2f534362619ac05b6b58b6b7f6680fee81e5c378d8
9a8e3da2b5], VERIFICATION PASSED
2018-03-08 02:57:10.950 IST [ledgerfsck] Verify -> INFO 036 block number [19]: previous hash matched
2018-03-08 02:57:10.951 IST [ledgerfsck] Verify -> INFO 037 Block [seq = 19], hash = [c2a2ef321f439a7420168c54d9adf250e820870a837db1e64eb095041f8135d3], previous hash = [fa1f93ad56c761d08c7fce71f2087db4e6fd4de3648fe02bf1accf
53ecfc6da8], VERIFICATION PASSED
2018-03-08 02:57:10.951 IST [ledgerfsck] Verify -> INFO 038 block number [20]: previous hash matched
2018-03-08 02:57:10.951 IST [ledgerfsck] Verify -> INFO 039 Block [seq = 20], hash = [f4a47b61832427e8b8e9e35990d91c0596d61c38138c2de105ef9620df2f2a19], previous hash = [c2a2ef321f439a7420168c54d9adf250e820870a837db1e64eb095
041f8135d3], VERIFICATION PASSED
2018-03-08 02:57:10.952 IST [ledgerfsck] Verify -> INFO 03a block number [21]: previous hash matched
2018-03-08 02:57:10.952 IST [ledgerfsck] Verify -> INFO 03b Block [seq = 21], hash = [2ac5024423155f68bc598486594a981e8863a152fd3ce1ee88d80bf3e9f024d5], previous hash = [f4a47b61832427e8b8e9e35990d91c0596d61c38138c2de105ef96
20df2f2a19], VERIFICATION PASSED
2018-03-08 02:57:10.954 IST [ledgerfsck] Verify -> INFO 03c block number [22]: previous hash matched
2018-03-08 02:57:10.955 IST [ledgerfsck] Verify -> INFO 03d Block [seq = 22], hash = [7b042659c419b19d68671e7db6e7e6614f41e07fc875293454aab4daf83872f2], previous hash = [2ac5024423155f68bc598486594a981e8863a152fd3ce1ee88d80b
f3e9f024d5], VERIFICATION PASSED
2018-03-08 02:57:10.955 IST [ledgerfsck] Verify -> INFO 03e block number [23]: previous hash matched
2018-03-08 02:57:10.956 IST [ledgerfsck] Verify -> INFO 03f Block [seq = 23], hash = [f91a792dbc76307aa156ce37dffdbc0dbb27492aa234bba14b54602b68b65539], previous hash = [7b042659c419b19d68671e7db6e7e6614f41e07fc875293454aab4
daf83872f2], VERIFICATION PASSED
2018-03-08 02:57:10.956 IST [ledgerfsck] Verify -> INFO 040 block number [24]: previous hash matched
2018-03-08 02:57:10.956 IST [ledgerfsck] Verify -> INFO 041 Block [seq = 24], hash = [c162fb69dc09d05ff2273a8a4871e36a6894f6dd4557935abc51997232136484], previous hash = [f91a792dbc76307aa156ce37dffdbc0dbb27492aa234bba14b5460
2b68b65539], VERIFICATION PASSED
2018-03-08 02:57:10.956 IST [ledgerfsck] Verify -> INFO 042 block number [25]: previous hash matched
2018-03-08 02:57:10.957 IST [ledgerfsck] Verify -> INFO 043 Block [seq = 25], hash = [a1efda1a7aed65110acb19c9a5636b6ce7f0015b13e97e9310402516cf5689af], previous hash = [c162fb69dc09d05ff2273a8a4871e36a6894f6dd4557935abc5199
7232136484], VERIFICATION PASSED
2018-03-08 02:57:10.957 IST [ledgerfsck] Verify -> INFO 044 block number [26]: previous hash matched
2018-03-08 02:57:10.957 IST [ledgerfsck] Verify -> INFO 045 Block [seq = 26], hash = [1194b7539337159c1c9cfdeb970b9675f27ff25b165bdba6775c5323e9c37cf4], previous hash = [a1efda1a7aed65110acb19c9a5636b6ce7f0015b13e97e93104025
16cf5689af], VERIFICATION PASSED
2018-03-08 02:57:10.957 IST [ledgerfsck] Verify -> INFO 046 block number [27]: previous hash matched
2018-03-08 02:57:10.958 IST [ledgerfsck] Verify -> INFO 047 Block [seq = 27], hash = [abdb2998540a0e9ae4d3dde998bf52f0bdbe117827964f67825b27c969b46010], previous hash = [1194b7539337159c1c9cfdeb970b9675f27ff25b165bdba6775c53
23e9c37cf4], VERIFICATION PASSED
2018-03-08 02:57:10.958 IST [ledgerfsck] Verify -> INFO 048 block number [28]: previous hash matched
2018-03-08 02:57:10.958 IST [ledgerfsck] Verify -> INFO 049 Block [seq = 28], hash = [98ece8b963b9fdf81f5a8eb5f59eafdf39398ef67f99f1e6a07fd1b2df52af28], previous hash = [abdb2998540a0e9ae4d3dde998bf52f0bdbe117827964f67825b27
c969b46010], VERIFICATION PASSED
2018-03-08 02:57:10.959 IST [ledgerfsck] Verify -> INFO 04a block number [29]: previous hash matched
2018-03-08 02:57:10.959 IST [ledgerfsck] Verify -> INFO 04b Block [seq = 29], hash = [4409cfbe09ffcd748ffe921d9523c157157531e86e9a3f6a59ec65887d91bac4], previous hash = [98ece8b963b9fdf81f5a8eb5f59eafdf39398ef67f99f1e6a07fd1
b2df52af28], VERIFICATION PASSED
2018-03-08 02:57:10.959 IST [ledgerfsck] Verify -> INFO 04c block number [30]: previous hash matched
2018-03-08 02:57:10.959 IST [ledgerfsck] Verify -> INFO 04d Block [seq = 30], hash = [5a06b4a7d31bf0c61f7ff267a28ad9ce919990ad244d700e1e34ac8d876e51d8], previous hash = [4409cfbe09ffcd748ffe921d9523c157157531e86e9a3f6a59ec65
887d91bac4], VERIFICATION PASSED
2018-03-08 02:57:10.960 IST [ledgerfsck] Verify -> INFO 04e block number [31]: previous hash matched
2018-03-08 02:57:10.960 IST [ledgerfsck] Verify -> INFO 04f Block [seq = 31], hash = [c5e4c5d056f48f129c252259dec99b1e579693ac8b9db674b5a3a060e6b363a6], previous hash = [5a06b4a7d31bf0c61f7ff267a28ad9ce919990ad244d700e1e34ac
8d876e51d8], VERIFICATION PASSED
2018-03-08 02:57:10.960 IST [ledgerfsck] Verify -> INFO 050 block number [32]: previous hash matched
2018-03-08 02:57:10.960 IST [ledgerfsck] Verify -> INFO 051 Block [seq = 32], hash = [4c191392c7b8b10533cd0c945864a8d95248ae70e9427d63f5842ec472e75925], previous hash = [c5e4c5d056f48f129c252259dec99b1e579693ac8b9db674b5a3a0
60e6b363a6], VERIFICATION PASSED
2018-03-08 02:57:10.961 IST [ledgerfsck] Verify -> INFO 052 block number [33]: previous hash matched
2018-03-08 02:57:10.961 IST [ledgerfsck] Verify -> INFO 053 Block [seq = 33], hash = [ca5cac7626777dd33808a7930e5a39daa7f6fb4ba27e8e808461e81fa1b02395], previous hash = [4c191392c7b8b10533cd0c945864a8d95248ae70e9427d63f5842e
c472e75925], VERIFICATION PASSED
2018-03-08 02:57:10.961 IST [ledgerfsck] Verify -> INFO 054 block number [34]: previous hash matched
2018-03-08 02:57:10.962 IST [ledgerfsck] Verify -> INFO 055 Block [seq = 34], hash = [4253704839936e0123c63de81f05b1a8e4016cb685f6896a13985c439839a126], previous hash = [ca5cac7626777dd33808a7930e5a39daa7f6fb4ba27e8e808461e8
1fa1b02395], VERIFICATION PASSED
2018-03-08 02:57:10.962 IST [ledgerfsck] Verify -> INFO 056 block number [35]: previous hash matched
2018-03-08 02:57:10.962 IST [ledgerfsck] Verify -> INFO 057 Block [seq = 35], hash = [59f404222f34a31a5fe213e600e0e90f6d617d7436c9b95690f58b6b6853369c], previous hash = [4253704839936e0123c63de81f05b1a8e4016cb685f6896a13985c
439839a126], VERIFICATION PASSED
2018-03-08 02:57:10.962 IST [ledgerfsck] Verify -> INFO 058 block number [36]: previous hash matched
2018-03-08 02:57:10.963 IST [ledgerfsck] Verify -> INFO 059 Block [seq = 36], hash = [0e7b67732ed6a65c6b20d1cd74115a05c5086557bf3f21ed071f8527afef75d1], previous hash = [59f404222f34a31a5fe213e600e0e90f6d617d7436c9b95690f58b
6b6853369c], VERIFICATION PASSED
2018-03-08 02:57:10.963 IST [ledgerfsck] Verify -> INFO 05a block number [37]: previous hash matched
2018-03-08 02:57:10.963 IST [ledgerfsck] Verify -> INFO 05b Block [seq = 37], hash = [bee9e5c76bc5908625be30dd53246cadfd8d158db9235bb08cc065f824caf79d], previous hash = [0e7b67732ed6a65c6b20d1cd74115a05c5086557bf3f21ed071f85
27afef75d1], VERIFICATION PASSED
2018-03-08 02:57:10.963 IST [ledgerfsck] Verify -> INFO 05c block number [38]: previous hash matched
2018-03-08 02:57:10.964 IST [ledgerfsck] Verify -> INFO 05d Block [seq = 38], hash = [f576123a523b3dcd5f0d391bbd0273318c52473e91fbec2a9baca2d77eb41418], previous hash = [bee9e5c76bc5908625be30dd53246cadfd8d158db9235bb08cc065
f824caf79d], VERIFICATION PASSED
2018-03-08 02:57:10.964 IST [ledgerfsck] Verify -> INFO 05e block number [39]: previous hash matched
2018-03-08 02:57:10.964 IST [ledgerfsck] Verify -> INFO 05f Block [seq = 39], hash = [4a58fddca6999703d49430245413377a757f1b3c07a8c62b9e2ea80781ad5635], previous hash = [f576123a523b3dcd5f0d391bbd0273318c52473e91fbec2a9baca2
d77eb41418], VERIFICATION PASSED
2018-03-08 02:57:10.964 IST [ledgerfsck] Verify -> INFO 060 block number [40]: previous hash matched
2018-03-08 02:57:10.965 IST [ledgerfsck] Verify -> INFO 061 Block [seq = 40], hash = [f35433595162bd964d0a833ccbe01f34f7f7adf78eba504c95dbe62fa845d787], previous hash = [4a58fddca6999703d49430245413377a757f1b3c07a8c62b9e2ea8
0781ad5635], VERIFICATION PASSED
2018-03-08 02:57:10.965 IST [ledgerfsck] Verify -> INFO 062 block number [41]: previous hash matched
2018-03-08 02:57:10.965 IST [ledgerfsck] Verify -> INFO 063 Block [seq = 41], hash = [bbe86863af87883b9de6d4300d900b25d38de47b92709cc301a616f06d33c23e], previous hash = [f35433595162bd964d0a833ccbe01f34f7f7adf78eba504c95dbe6
2fa845d787], VERIFICATION PASSED
2018-03-08 02:57:10.966 IST [ledgerfsck] Verify -> INFO 064 block number [42]: previous hash matched
2018-03-08 02:57:10.966 IST [ledgerfsck] Verify -> INFO 065 Block [seq = 42], hash = [0feb604e2602ef3e7539125747d0a3aef456cdaade9ca82f5c49e283a8719108], previous hash = [bbe86863af87883b9de6d4300d900b25d38de47b92709cc301a616
f06d33c23e], VERIFICATION PASSED
2018-03-08 02:57:10.966 IST [ledgerfsck] Verify -> INFO 066 block number [43]: previous hash matched
2018-03-08 02:57:10.967 IST [ledgerfsck] Verify -> INFO 067 Block [seq = 43], hash = [19adc48355397f6b3308f628683e0f68658a9446208a735a8eb5b6aeef402932], previous hash = [0feb604e2602ef3e7539125747d0a3aef456cdaade9ca82f5c49e2
83a8719108], VERIFICATION PASSED
2018-03-08 02:57:10.967 IST [ledgerfsck] Verify -> INFO 068 block number [44]: previous hash matched
2018-03-08 02:57:10.967 IST [ledgerfsck] Verify -> INFO 069 Block [seq = 44], hash = [fabcc4caafc9abcd62cfff7677983781d1d0263592e173fc3b815c0fcd170f55], previous hash = [19adc48355397f6b3308f628683e0f68658a9446208a735a8eb5b6
aeef402932], VERIFICATION PASSED
2018-03-08 02:57:10.968 IST [ledgerfsck] Verify -> INFO 06a block number [45]: previous hash matched
2018-03-08 02:57:10.968 IST [ledgerfsck] Verify -> INFO 06b Block [seq = 45], hash = [88f30451aaed29328bce6951de37eb4146c81133ad78897d9fe95a5181a3c57b], previous hash = [fabcc4caafc9abcd62cfff7677983781d1d0263592e173fc3b815c
0fcd170f55], VERIFICATION PASSED
2018-03-08 02:57:10.968 IST [ledgerfsck] Verify -> INFO 06c block number [46]: previous hash matched
2018-03-08 02:57:10.969 IST [ledgerfsck] Verify -> INFO 06d Block [seq = 46], hash = [3bc94de05e93d12bd98e5a7abb21028b6b2fbcffb21930996bbce5f4553f473f], previous hash = [88f30451aaed29328bce6951de37eb4146c81133ad78897d9fe95a
5181a3c57b], VERIFICATION PASSED
2018-03-08 02:57:10.969 IST [ledgerfsck] Verify -> INFO 06e block number [47]: previous hash matched
2018-03-08 02:57:10.969 IST [ledgerfsck] Verify -> INFO 06f Block [seq = 47], hash = [4410697aff88b18680e10180e7a9569326d38df63e5a917e815e540ba89a246e], previous hash = [3bc94de05e93d12bd98e5a7abb21028b6b2fbcffb21930996bbce5
f4553f473f], VERIFICATION PASSED
2018-03-08 02:57:10.970 IST [ledgerfsck] Verify -> INFO 070 block number [48]: previous hash matched
2018-03-08 02:57:10.970 IST [ledgerfsck] Verify -> INFO 071 Block [seq = 48], hash = [c70ec26d513eab674aa3f39399a011743636af09a011e953d99d785862f6ce59], previous hash = [4410697aff88b18680e10180e7a9569326d38df63e5a917e815e54
0ba89a246e], VERIFICATION PASSED
2018-03-08 02:57:10.970 IST [ledgerfsck] Verify -> INFO 072 block number [49]: previous hash matched
2018-03-08 02:57:10.971 IST [ledgerfsck] Verify -> INFO 073 Block [seq = 49], hash = [d58a633faf364550b7cd39be3752bd8ee314915f2b81c8211c5f683a65e8783b], previous hash = [c70ec26d513eab674aa3f39399a011743636af09a011e953d99d78
5862f6ce59], VERIFICATION PASSED
2018-03-08 02:57:10.971 IST [ledgerfsck] Verify -> INFO 074 block number [50]: previous hash matched
2018-03-08 02:57:10.972 IST [ledgerfsck] Verify -> INFO 075 Block [seq = 50], hash = [46c90c450fd52d96dc873b3de9e7d6c22ce34ed72befbc527db2b44013aabcb8], previous hash = [d58a633faf364550b7cd39be3752bd8ee314915f2b81c8211c5f68
3a65e8783b], VERIFICATION PASSED
2018-03-08 02:57:10.972 IST [ledgerfsck] Verify -> INFO 076 block number [51]: previous hash matched
2018-03-08 02:57:10.972 IST [ledgerfsck] Verify -> INFO 077 Block [seq = 51], hash = [46e2921b6415b18f469181495a48fceeebdd3b88db2382d527c3a57ede4f2759], previous hash = [46c90c450fd52d96dc873b3de9e7d6c22ce34ed72befbc527db2b4
4013aabcb8], VERIFICATION PASSED
2018-03-08 02:57:10.972 IST [ledgerfsck] Verify -> INFO 078 block number [52]: previous hash matched
2018-03-08 02:57:10.973 IST [ledgerfsck] Verify -> INFO 079 Block [seq = 52], hash = [69eee6df67f2bfc634093b290cf520d3691834c6b50d749e65ff8db4f164c416], previous hash = [46e2921b6415b18f469181495a48fceeebdd3b88db2382d527c3a5
7ede4f2759], VERIFICATION PASSED
2018-03-08 02:57:10.973 IST [ledgerfsck] Verify -> INFO 07a block number [53]: previous hash matched
2018-03-08 02:57:10.974 IST [ledgerfsck] Verify -> INFO 07b Block [seq = 53], hash = [4dab4a702708fef3bf50319ae2a1f54447bc24a1f6f0e3225a942837b907433b], previous hash = [69eee6df67f2bfc634093b290cf520d3691834c6b50d749e65ff8d
b4f164c416], VERIFICATION PASSED
2018-03-08 02:57:10.974 IST [ledgerfsck] Verify -> INFO 07c block number [54]: previous hash matched
2018-03-08 02:57:10.974 IST [ledgerfsck] Verify -> INFO 07d Block [seq = 54], hash = [14861fbf6c9c68ce775e494ef0f6fefea0bc4c733444c548f6940177a92f5ee2], previous hash = [4dab4a702708fef3bf50319ae2a1f54447bc24a1f6f0e3225a9428
37b907433b], VERIFICATION PASSED
2018-03-08 02:57:10.974 IST [ledgerfsck] Verify -> INFO 07e block number [55]: previous hash matched
2018-03-08 02:57:10.975 IST [ledgerfsck] Verify -> INFO 07f Block [seq = 55], hash = [d45aa2ecfe82bfa0f2dae22dc8dc77a6c640fc74cf078d7f0e039e74d4e72f27], previous hash = [14861fbf6c9c68ce775e494ef0f6fefea0bc4c733444c548f69401
77a92f5ee2], VERIFICATION PASSED
2018-03-08 02:57:10.975 IST [ledgerfsck] Verify -> INFO 080 block number [56]: previous hash matched
2018-03-08 02:57:10.976 IST [ledgerfsck] Verify -> INFO 081 Block [seq = 56], hash = [d20eae3dce351bfcc185589242c9aecfb0a32787d613f61f1fd3cade4f819056], previous hash = [d45aa2ecfe82bfa0f2dae22dc8dc77a6c640fc74cf078d7f0e039e
74d4e72f27], VERIFICATION PASSED
2018-03-08 02:57:10.976 IST [ledgerfsck] Verify -> INFO 082 block number [57]: previous hash matched
2018-03-08 02:57:10.976 IST [ledgerfsck] Verify -> INFO 083 Block [seq = 57], hash = [15bc739a87823dbf36ed65b8d7925dc87a75c4725d758c7f6e608e6c24732cf3], previous hash = [d20eae3dce351bfcc185589242c9aecfb0a32787d613f61f1fd3ca
de4f819056], VERIFICATION PASSED
2018-03-08 02:57:10.978 IST [ledgerfsck] Verify -> INFO 084 block number [58]: previous hash matched
2018-03-08 02:57:10.978 IST [ledgerfsck] Verify -> INFO 085 Block [seq = 58], hash = [b0b7026fcc97d9121d7ed90cb08dae56ae330bbf873110d6bf06cc6a81510c6a], previous hash = [15bc739a87823dbf36ed65b8d7925dc87a75c4725d758c7f6e608e
6c24732cf3], VERIFICATION PASSED
2018-03-08 02:57:10.979 IST [ledgerfsck] Verify -> INFO 086 block number [59]: previous hash matched
2018-03-08 02:57:10.979 IST [ledgerfsck] Verify -> INFO 087 Block [seq = 59], hash = [d7f6b8f7bb08662313e6a5ec1db43044d8083d9adddd563955e9fd520a9fc9d3], previous hash = [b0b7026fcc97d9121d7ed90cb08dae56ae330bbf873110d6bf06cc
6a81510c6a], VERIFICATION PASSED
2018-03-08 02:57:10.980 IST [ledgerfsck] Verify -> INFO 088 block number [60]: previous hash matched
2018-03-08 02:57:10.980 IST [ledgerfsck] Verify -> INFO 089 Block [seq = 60], hash = [f670027e16e475ab127bacb9356acbf865991c4d2fe1ee9aa1b3e6dc12651a03], previous hash = [d7f6b8f7bb08662313e6a5ec1db43044d8083d9adddd563955e9fd
520a9fc9d3], VERIFICATION PASSED
2018-03-08 02:57:10.980 IST [ledgerfsck] Verify -> INFO 08a block number [61]: previous hash matched
2018-03-08 02:57:10.981 IST [ledgerfsck] Verify -> INFO 08b Block [seq = 61], hash = [babba33bb149814e26fe67b217fa5cd177d9057745a84570180b987e5bc40910], previous hash = [f670027e16e475ab127bacb9356acbf865991c4d2fe1ee9aa1b3e6
dc12651a03], VERIFICATION PASSED
2018-03-08 02:57:10.981 IST [ledgerfsck] Verify -> INFO 08c block number [62]: previous hash matched
2018-03-08 02:57:10.982 IST [ledgerfsck] Verify -> INFO 08d Block [seq = 62], hash = [ea6add9c9a0c29e43133bc333a7d88f82065f225f20efbf125e33d4e5060dbf9], previous hash = [babba33bb149814e26fe67b217fa5cd177d9057745a84570180b98
7e5bc40910], VERIFICATION PASSED
2018-03-08 02:57:10.982 IST [ledgerfsck] Verify -> INFO 08e block number [63]: previous hash matched
2018-03-08 02:57:10.982 IST [ledgerfsck] Verify -> INFO 08f Block [seq = 63], hash = [b3143d138f6b4e89dbc8729027bf8b57f8da9cbaecee47424a7e573ce16c7074], previous hash = [ea6add9c9a0c29e43133bc333a7d88f82065f225f20efbf125e33d
4e5060dbf9], VERIFICATION PASSED
2018-03-08 02:57:10.983 IST [ledgerfsck] Verify -> INFO 090 block number [64]: previous hash matched
2018-03-08 02:57:10.983 IST [ledgerfsck] Verify -> INFO 091 Block [seq = 64], hash = [21335a1ab74682754041e06332fcb6157a99608399626e4405d301aae85c2381], previous hash = [b3143d138f6b4e89dbc8729027bf8b57f8da9cbaecee47424a7e57
3ce16c7074], VERIFICATION PASSED
2018-03-08 02:57:10.983 IST [ledgerfsck] Verify -> INFO 092 block number [65]: previous hash matched
2018-03-08 02:57:10.984 IST [ledgerfsck] Verify -> INFO 093 Block [seq = 65], hash = [1018e603e242ef78efcb72c26d420b1a271989e2e223771d8b2c69ddab52e65b], previous hash = [21335a1ab74682754041e06332fcb6157a99608399626e4405d301
aae85c2381], VERIFICATION PASSED
2018-03-08 02:57:10.984 IST [ledgerfsck] Verify -> INFO 094 block number [66]: previous hash matched
2018-03-08 02:57:10.985 IST [ledgerfsck] Verify -> INFO 095 Block [seq = 66], hash = [4208c154d60e97ae8b980f96eae19564e430eb1464f07c62358a593ac44017b6], previous hash = [1018e603e242ef78efcb72c26d420b1a271989e2e223771d8b2c69
ddab52e65b], VERIFICATION PASSED
2018-03-08 02:57:10.985 IST [ledgerfsck] Verify -> INFO 096 block number [67]: previous hash matched
2018-03-08 02:57:10.985 IST [ledgerfsck] Verify -> INFO 097 Block [seq = 67], hash = [06172077eb62b87f9e43c8c333c8232bb4a5da1a1ca9c6b2fdb5f68239f591e8], previous hash = [4208c154d60e97ae8b980f96eae19564e430eb1464f07c62358a59
3ac44017b6], VERIFICATION PASSED
2018-03-08 02:57:10.986 IST [ledgerfsck] Verify -> INFO 098 block number [68]: previous hash matched
2018-03-08 02:57:10.986 IST [ledgerfsck] Verify -> INFO 099 Block [seq = 68], hash = [1c1d0a4553301078ca72d7ddf95add36be1134fc94d55fdc48ad37c6c3aca4e7], previous hash = [06172077eb62b87f9e43c8c333c8232bb4a5da1a1ca9c6b2fdb5f6
8239f591e8], VERIFICATION PASSED
2018-03-08 02:57:10.986 IST [ledgerfsck] Verify -> INFO 09a block number [69]: previous hash matched
2018-03-08 02:57:10.987 IST [ledgerfsck] Verify -> INFO 09b Block [seq = 69], hash = [6803ed3ae35f22a5a32aebb5387e256bf006ea74dd8b53291dd60b32a2e66081], previous hash = [1c1d0a4553301078ca72d7ddf95add36be1134fc94d55fdc48ad37
c6c3aca4e7], VERIFICATION PASSED
2018-03-08 02:57:10.987 IST [ledgerfsck] Verify -> INFO 09c block number [70]: previous hash matched
2018-03-08 02:57:10.987 IST [ledgerfsck] Verify -> INFO 09d Block [seq = 70], hash = [9f02d9d57c16d14809871f5007c0fe6b5a55ae7ec740747ce6f1666e95438f18], previous hash = [6803ed3ae35f22a5a32aebb5387e256bf006ea74dd8b53291dd60b
32a2e66081], VERIFICATION PASSED
2018-03-08 02:57:10.988 IST [ledgerfsck] Verify -> INFO 09e block number [71]: previous hash matched
2018-03-08 02:57:10.988 IST [ledgerfsck] Verify -> INFO 09f Block [seq = 71], hash = [8ce9932595a2a5949189a0dc18845a04ca24aacaf00508e5e69a7a6580de4981], previous hash = [9f02d9d57c16d14809871f5007c0fe6b5a55ae7ec740747ce6f166
6e95438f18], VERIFICATION PASSED
2018-03-08 02:57:10.989 IST [ledgerfsck] Verify -> INFO 0a0 block number [72]: previous hash matched
2018-03-08 02:57:10.989 IST [ledgerfsck] Verify -> INFO 0a1 Block [seq = 72], hash = [d6e70f0bb41501aff0f4430480db76dd445bc6fb3bc588acc44721462e6cd750], previous hash = [8ce9932595a2a5949189a0dc18845a04ca24aacaf00508e5e69a7a
6580de4981], VERIFICATION PASSED
2018-03-08 02:57:10.989 IST [ledgerfsck] Verify -> INFO 0a2 block number [73]: previous hash matched
2018-03-08 02:57:10.989 IST [ledgerfsck] Verify -> INFO 0a3 Block [seq = 73], hash = [2f05da54cca8503362da2249b7b13d39eade2a437d3a3d5b156c0aa80efc8d2b], previous hash = [d6e70f0bb41501aff0f4430480db76dd445bc6fb3bc588acc44721
462e6cd750], VERIFICATION PASSED
2018-03-08 02:57:10.990 IST [ledgerfsck] Verify -> INFO 0a4 block number [74]: previous hash matched
2018-03-08 02:57:10.990 IST [ledgerfsck] Verify -> INFO 0a5 Block [seq = 74], hash = [6b6fc2a2a3ede5d573469c399ecb5a43aea44bdd4d24912bfe94c692b45a10cd], previous hash = [2f05da54cca8503362da2249b7b13d39eade2a437d3a3d5b156c0a
a80efc8d2b], VERIFICATION PASSED
2018-03-08 02:57:10.990 IST [ledgerfsck] Verify -> INFO 0a6 block number [75]: previous hash matched
2018-03-08 02:57:10.990 IST [ledgerfsck] Verify -> INFO 0a7 Block [seq = 75], hash = [cba48915a58ab5e951a1785c390d5b6a79667f3ef59126d7955e5617f8c6fdb5], previous hash = [6b6fc2a2a3ede5d573469c399ecb5a43aea44bdd4d24912bfe94c6
92b45a10cd], VERIFICATION PASSED
2018-03-08 02:57:10.990 IST [ledgerfsck] Verify -> INFO 0a8 block number [76]: previous hash matched
2018-03-08 02:57:10.991 IST [ledgerfsck] Verify -> INFO 0a9 Block [seq = 76], hash = [cb83bd8430262fc1072a37ad800582d1880c2744ee7f7ee5546ff9d44e05df6e], previous hash = [cba48915a58ab5e951a1785c390d5b6a79667f3ef59126d7955e56
17f8c6fdb5], VERIFICATION PASSED
2018-03-08 02:57:10.991 IST [ledgerfsck] Verify -> INFO 0aa block number [77]: previous hash matched
2018-03-08 02:57:10.991 IST [ledgerfsck] Verify -> INFO 0ab Block [seq = 77], hash = [a1092ea1fc5461205194628d0c62bff32533a00d90723df273da5f014d54da04], previous hash = [cb83bd8430262fc1072a37ad800582d1880c2744ee7f7ee5546ff9
d44e05df6e], VERIFICATION PASSED
2018-03-08 02:57:10.991 IST [ledgerfsck] Verify -> INFO 0ac block number [78]: previous hash matched
2018-03-08 02:57:10.992 IST [ledgerfsck] Verify -> INFO 0ad Block [seq = 78], hash = [274ff3dbcf13e9b47efabbdc8efe8cfc45ba5166e260d479ce72d4e9d059c74d], previous hash = [a1092ea1fc5461205194628d0c62bff32533a00d90723df273da5f
014d54da04], VERIFICATION PASSED
2018-03-08 02:57:10.992 IST [ledgerfsck] Verify -> INFO 0ae block number [79]: previous hash matched
2018-03-08 02:57:10.992 IST [ledgerfsck] Verify -> INFO 0af Block [seq = 79], hash = [75bdd3965699b0c9be52a1811b676e27c2b15d6eed2c97db07adc563df1c951e], previous hash = [274ff3dbcf13e9b47efabbdc8efe8cfc45ba5166e260d479ce72d4
e9d059c74d], VERIFICATION PASSED
2018-03-08 02:57:10.992 IST [ledgerfsck] Verify -> INFO 0b0 block number [80]: previous hash matched
2018-03-08 02:57:10.993 IST [ledgerfsck] Verify -> INFO 0b1 Block [seq = 80], hash = [f8c2813bfc91ff294a10c966d35722566ad6ddfc7981b3ff3cc7826362f6637d], previous hash = [75bdd3965699b0c9be52a1811b676e27c2b15d6eed2c97db07adc5
63df1c951e], VERIFICATION PASSED
2018-03-08 02:57:10.993 IST [ledgerfsck] Verify -> INFO 0b2 block number [81]: previous hash matched
2018-03-08 02:57:10.993 IST [ledgerfsck] Verify -> INFO 0b3 Block [seq = 81], hash = [b5d90a413fd32ad45cc7759476cf1358a86eb0a3dea40a75efa9ee67bb9b7b3e], previous hash = [f8c2813bfc91ff294a10c966d35722566ad6ddfc7981b3ff3cc782
6362f6637d], VERIFICATION PASSED
2018-03-08 02:57:10.994 IST [ledgerfsck] Verify -> INFO 0b4 block number [82]: previous hash matched
2018-03-08 02:57:10.994 IST [ledgerfsck] Verify -> INFO 0b5 Block [seq = 82], hash = [339a3c8f5287b9c9f83612b96203619d4571c98175df8542a22c2f93afa6f81c], previous hash = [b5d90a413fd32ad45cc7759476cf1358a86eb0a3dea40a75efa9ee
67bb9b7b3e], VERIFICATION PASSED
2018-03-08 02:57:10.994 IST [ledgerfsck] Verify -> INFO 0b6 block number [83]: previous hash matched
2018-03-08 02:57:10.994 IST [ledgerfsck] Verify -> INFO 0b7 Block [seq = 83], hash = [c7d09ee566e1bdcfd6adc1654c9b9f5d8a434f059a3dcb61db98853bdcfeca73], previous hash = [339a3c8f5287b9c9f83612b96203619d4571c98175df8542a22c2f
93afa6f81c], VERIFICATION PASSED
2018-03-08 02:57:10.995 IST [ledgerfsck] Verify -> INFO 0b8 block number [84]: previous hash matched
2018-03-08 02:57:10.995 IST [ledgerfsck] Verify -> INFO 0b9 Block [seq = 84], hash = [65bb44fe8bdf1a95e20e061b939f18d84aff35348a3896afe6cd6efb55d807a5], previous hash = [c7d09ee566e1bdcfd6adc1654c9b9f5d8a434f059a3dcb61db9885
3bdcfeca73], VERIFICATION PASSED
2018-03-08 02:57:10.995 IST [ledgerfsck] Verify -> INFO 0ba block number [85]: previous hash matched
2018-03-08 02:57:10.995 IST [ledgerfsck] Verify -> INFO 0bb Block [seq = 85], hash = [dd85438b27465b63e66d490a65eb86cb17d1db8cd0e130b5e0d1e5526b1aa9dc], previous hash = [65bb44fe8bdf1a95e20e061b939f18d84aff35348a3896afe6cd6e
fb55d807a5], VERIFICATION PASSED
2018-03-08 02:57:10.996 IST [ledgerfsck] Verify -> INFO 0bc block number [86]: previous hash matched
2018-03-08 02:57:10.996 IST [ledgerfsck] Verify -> INFO 0bd Block [seq = 86], hash = [70a8db81743bf9ed8553287b0063cdd4c0fdbac4723a26588174d6c5d608d23d], previous hash = [dd85438b27465b63e66d490a65eb86cb17d1db8cd0e130b5e0d1e5
526b1aa9dc], VERIFICATION PASSED
2018-03-08 02:57:10.996 IST [ledgerfsck] Verify -> INFO 0be block number [87]: previous hash matched
2018-03-08 02:57:10.996 IST [ledgerfsck] Verify -> INFO 0bf Block [seq = 87], hash = [b0fe399e33b9ca507402eef8650e8dbe35c97bb859f29e589a95c48e4299fdcb], previous hash = [70a8db81743bf9ed8553287b0063cdd4c0fdbac4723a26588174d6
c5d608d23d], VERIFICATION PASSED
2018-03-08 02:57:10.996 IST [ledgerfsck] Verify -> INFO 0c0 block number [88]: previous hash matched
2018-03-08 02:57:10.997 IST [ledgerfsck] Verify -> INFO 0c1 Block [seq = 88], hash = [251d36791c7774d716da87774d3631d2a713667de1188c5c4e4318c017e6fbe5], previous hash = [b0fe399e33b9ca507402eef8650e8dbe35c97bb859f29e589a95c4
8e4299fdcb], VERIFICATION PASSED
2018-03-08 02:57:10.997 IST [ledgerfsck] Verify -> INFO 0c2 block number [89]: previous hash matched
2018-03-08 02:57:10.997 IST [ledgerfsck] Verify -> INFO 0c3 Block [seq = 89], hash = [ddfca7e51b865dccc83d064d0f10d52c5c6d0d22b5bcd3bd8da09f3714229ec7], previous hash = [251d36791c7774d716da87774d3631d2a713667de1188c5c4e4318
c017e6fbe5], VERIFICATION PASSED
2018-03-08 02:57:10.997 IST [ledgerfsck] Verify -> INFO 0c4 block number [90]: previous hash matched
2018-03-08 02:57:10.998 IST [ledgerfsck] Verify -> INFO 0c5 Block [seq = 90], hash = [1dc9e987eae2dcf54cec130055e8f589c2836fb60e3ec6a67000ffaca0be9c1c], previous hash = [ddfca7e51b865dccc83d064d0f10d52c5c6d0d22b5bcd3bd8da09f
3714229ec7], VERIFICATION PASSED
2018-03-08 02:57:10.998 IST [ledgerfsck] Verify -> INFO 0c6 block number [91]: previous hash matched
2018-03-08 02:57:10.998 IST [ledgerfsck] Verify -> INFO 0c7 Block [seq = 91], hash = [7dac43e4fd96527fb8fe9f6ef6968e9b1b374244ecfa3e6e80cdc65b05cb044f], previous hash = [1dc9e987eae2dcf54cec130055e8f589c2836fb60e3ec6a67000ff
aca0be9c1c], VERIFICATION PASSED
2018-03-08 02:57:10.998 IST [ledgerfsck] Verify -> INFO 0c8 block number [92]: previous hash matched
2018-03-08 02:57:10.999 IST [ledgerfsck] Verify -> INFO 0c9 Block [seq = 92], hash = [cfec2b4b101684406298afffb42a4f4dfbcbbb60c035346e407a57cfd0d667d1], previous hash = [7dac43e4fd96527fb8fe9f6ef6968e9b1b374244ecfa3e6e80cdc6
5b05cb044f], VERIFICATION PASSED
2018-03-08 02:57:10.999 IST [ledgerfsck] Verify -> INFO 0ca block number [93]: previous hash matched
2018-03-08 02:57:10.999 IST [ledgerfsck] Verify -> INFO 0cb Block [seq = 93], hash = [b6f325db476e44a5836a76c484cbfd68dbfa9a648c94038bcc66d8a464994058], previous hash = [cfec2b4b101684406298afffb42a4f4dfbcbbb60c035346e407a57
cfd0d667d1], VERIFICATION PASSED
2018-03-08 02:57:11.000 IST [ledgerfsck] Verify -> INFO 0cc block number [94]: previous hash matched
2018-03-08 02:57:11.000 IST [ledgerfsck] Verify -> INFO 0cd Block [seq = 94], hash = [17d3d0590ebd93f18680574fb16189bfd90ee9c52f4ca76233b77256d98084f0], previous hash = [b6f325db476e44a5836a76c484cbfd68dbfa9a648c94038bcc66d8
a464994058], VERIFICATION PASSED
2018-03-08 02:57:11.000 IST [ledgerfsck] Verify -> INFO 0ce block number [95]: previous hash matched
2018-03-08 02:57:11.000 IST [ledgerfsck] Verify -> INFO 0cf Block [seq = 95], hash = [224d00f8609520942e69130df811df506fe3af3902c316f971bcf28304a2c234], previous hash = [17d3d0590ebd93f18680574fb16189bfd90ee9c52f4ca76233b772
56d98084f0], VERIFICATION PASSED
2018-03-08 02:57:11.003 IST [ledgerfsck] Verify -> INFO 0d0 block number [96]: previous hash matched
2018-03-08 02:57:11.004 IST [ledgerfsck] Verify -> INFO 0d1 Block [seq = 96], hash = [ae4d5cb4559070154e163473e77cc41a0350c96f0c65cdc8594ccb76d67cd06e], previous hash = [224d00f8609520942e69130df811df506fe3af3902c316f971bcf2
8304a2c234], VERIFICATION PASSED
2018-03-08 02:57:11.004 IST [ledgerfsck] Verify -> INFO 0d2 block number [97]: previous hash matched
2018-03-08 02:57:11.004 IST [ledgerfsck] Verify -> INFO 0d3 Block [seq = 97], hash = [ae0f872b350d7f01503635207dca36c5b6847e9132504867063e95a33254ac33], previous hash = [ae4d5cb4559070154e163473e77cc41a0350c96f0c65cdc8594ccb
76d67cd06e], VERIFICATION PASSED
2018-03-08 02:57:11.004 IST [ledgerfsck] Verify -> INFO 0d4 block number [98]: previous hash matched
2018-03-08 02:57:11.005 IST [ledgerfsck] Verify -> INFO 0d5 Block [seq = 98], hash = [55e36ad0daa3b829258eafce24fded2aa2b1b006237088cc2ee9a8b387c9b773], previous hash = [ae0f872b350d7f01503635207dca36c5b6847e9132504867063e95
a33254ac33], VERIFICATION PASSED
2018-03-08 02:57:11.005 IST [ledgerfsck] Verify -> INFO 0d6 block number [99]: previous hash matched
2018-03-08 02:57:11.005 IST [ledgerfsck] Verify -> INFO 0d7 Block [seq = 99], hash = [ac9da4411c2d15c340dd87519369718900903f63281dcdc372f6cc9abb16b2fa], previous hash = [55e36ad0daa3b829258eafce24fded2aa2b1b006237088cc2ee9a8
b387c9b773], VERIFICATION PASSED
2018-03-08 02:57:11.006 IST [ledgerfsck] Verify -> INFO 0d8 block number [100]: previous hash matched
2018-03-08 02:57:11.006 IST [ledgerfsck] Verify -> INFO 0d9 Block [seq = 100], hash = [f3061eddf1c7c866b02bb712dd05c4a44e7921c93dfeeb109d291a52036daf21], previous hash = [ac9da4411c2d15c340dd87519369718900903f63281dcdc372f6c
c9abb16b2fa], VERIFICATION PASSED
2018-03-08 02:57:11.006 IST [ledgerfsck] Verify -> INFO 0da block number [101]: previous hash matched
2018-03-08 02:57:11.006 IST [ledgerfsck] Verify -> INFO 0db Block [seq = 101], hash = [fcf96572a393440ad6495fea869457304da033baac93ad58677f2d8ef4a81d23], previous hash = [f3061eddf1c7c866b02bb712dd05c4a44e7921c93dfeeb109d291
a52036daf21], VERIFICATION PASSED
2018-03-08 02:57:11.007 IST [ledgerfsck] Verify -> INFO 0dc block number [102]: previous hash matched
2018-03-08 02:57:11.007 IST [ledgerfsck] Verify -> INFO 0dd Block [seq = 102], hash = [74dc3b599a463656e38a1782acbc71473a3d710c2ff250afcb5da6e1e0fb00bf], previous hash = [fcf96572a393440ad6495fea869457304da033baac93ad58677f2
d8ef4a81d23], VERIFICATION PASSED
2018-03-08 02:57:11.007 IST [ledgerfsck] Verify -> INFO 0de block number [103]: previous hash matched
2018-03-08 02:57:11.007 IST [ledgerfsck] Verify -> INFO 0df Block [seq = 103], hash = [47ccfe029b202f95e1955c5a281d6567300083afd77e9d05c68826ab3e5ebcf1], previous hash = [74dc3b599a463656e38a1782acbc71473a3d710c2ff250afcb5da
6e1e0fb00bf], VERIFICATION PASSED
2018-03-08 02:57:11.007 IST [ledgerfsck] Verify -> INFO 0e0 block number [104]: previous hash matched
2018-03-08 02:57:11.008 IST [ledgerfsck] Verify -> INFO 0e1 Block [seq = 104], hash = [25944e93cc72da73137f6f8958e85f0ed0740ca559de0c2d6856d938de947d4d], previous hash = [47ccfe029b202f95e1955c5a281d6567300083afd77e9d05c6882
6ab3e5ebcf1], VERIFICATION PASSED
2018-03-08 02:57:11.008 IST [ledgerfsck] Verify -> INFO 0e2 block number [105]: previous hash matched
2018-03-08 02:57:11.008 IST [ledgerfsck] Verify -> INFO 0e3 Block [seq = 105], hash = [c5a743d46066912e50729dc204ab2fde558209ce950273a4612548e01ebea0b4], previous hash = [25944e93cc72da73137f6f8958e85f0ed0740ca559de0c2d6856d
938de947d4d], VERIFICATION PASSED
2018-03-08 02:57:11.008 IST [ledgerfsck] Verify -> INFO 0e4 block number [106]: previous hash matched
2018-03-08 02:57:11.009 IST [ledgerfsck] Verify -> INFO 0e5 Block [seq = 106], hash = [6fce894a179ee645cace06a7b430d18f7f1eabeb154423d76bde994cb1711e20], previous hash = [c5a743d46066912e50729dc204ab2fde558209ce950273a461254
8e01ebea0b4], VERIFICATION PASSED
2018-03-08 02:57:11.009 IST [ledgerfsck] Verify -> INFO 0e6 block number [107]: previous hash matched
2018-03-08 02:57:11.009 IST [ledgerfsck] Verify -> INFO 0e7 Block [seq = 107], hash = [694695dbd9d9c4f889ada4e05939003cc9ae6b94b4870002691bbd62f9d6230e], previous hash = [6fce894a179ee645cace06a7b430d18f7f1eabeb154423d76bde9
94cb1711e20], VERIFICATION PASSED
2018-03-08 02:57:11.010 IST [ledgerfsck] Verify -> INFO 0e8 block number [108]: previous hash matched
2018-03-08 02:57:11.010 IST [ledgerfsck] Verify -> INFO 0e9 Block [seq = 108], hash = [1389635d275313cdb5928326a9671667dbeb4c825de74dd87cfeeaf1407dcdb3], previous hash = [694695dbd9d9c4f889ada4e05939003cc9ae6b94b4870002691bb
d62f9d6230e], VERIFICATION PASSED
2018-03-08 02:57:11.010 IST [ledgerfsck] Verify -> INFO 0ea block number [109]: previous hash matched
2018-03-08 02:57:11.010 IST [ledgerfsck] Verify -> INFO 0eb Block [seq = 109], hash = [4ea907eb885439b7aaf07ad397d00f34eb89a3a4b8b15cb227a5cdfba43eece4], previous hash = [1389635d275313cdb5928326a9671667dbeb4c825de74dd87cfee
af1407dcdb3], VERIFICATION PASSED
2018-03-08 02:57:11.011 IST [ledgerfsck] Verify -> INFO 0ec block number [110]: previous hash matched
2018-03-08 02:57:11.011 IST [ledgerfsck] Verify -> INFO 0ed Block [seq = 110], hash = [095199245fad8ce6cea11f3c0cfdd909af9f5d6651af44578b7aea045293d49d], previous hash = [4ea907eb885439b7aaf07ad397d00f34eb89a3a4b8b15cb227a5c
dfba43eece4], VERIFICATION PASSED
2018-03-08 02:57:11.011 IST [ledgerfsck] Verify -> INFO 0ee block number [111]: previous hash matched
2018-03-08 02:57:11.011 IST [ledgerfsck] Verify -> INFO 0ef Block [seq = 111], hash = [3e2111ccf555210d8c7bfdd90985cd157ac1af411ddcf2c2533e53bd03aafdab], previous hash = [095199245fad8ce6cea11f3c0cfdd909af9f5d6651af44578b7ae
a045293d49d], VERIFICATION PASSED
2018-03-08 02:57:11.012 IST [ledgerfsck] Verify -> INFO 0f0 block number [112]: previous hash matched
2018-03-08 02:57:11.012 IST [ledgerfsck] Verify -> INFO 0f1 Block [seq = 112], hash = [890e8fa34aa604c3d65f5e2fdc27bd4f34bed9c2c7f54abf1fdb5d7b5f809bee], previous hash = [3e2111ccf555210d8c7bfdd90985cd157ac1af411ddcf2c2533e5
3bd03aafdab], VERIFICATION PASSED
2018-03-08 02:57:11.012 IST [ledgerfsck] Verify -> INFO 0f2 block number [113]: previous hash matched
2018-03-08 02:57:11.013 IST [ledgerfsck] Verify -> INFO 0f3 Block [seq = 113], hash = [020dc88b8311050fa0580d745ffa94a974ad71eb3723b56caaa8b7528c59cd45], previous hash = [890e8fa34aa604c3d65f5e2fdc27bd4f34bed9c2c7f54abf1fdb5
d7b5f809bee], VERIFICATION PASSED
2018-03-08 02:57:11.013 IST [ledgerfsck] Verify -> INFO 0f4 block number [114]: previous hash matched
2018-03-08 02:57:11.013 IST [ledgerfsck] Verify -> INFO 0f5 Block [seq = 114], hash = [1b64d820db528ac48e2331150d2db0c65bf2b2a283c851dcd48d8bef764b20b8], previous hash = [020dc88b8311050fa0580d745ffa94a974ad71eb3723b56caaa8b
7528c59cd45], VERIFICATION PASSED
2018-03-08 02:57:11.013 IST [ledgerfsck] Verify -> INFO 0f6 block number [115]: previous hash matched
2018-03-08 02:57:11.014 IST [ledgerfsck] Verify -> INFO 0f7 Block [seq = 115], hash = [36d57ac3839e0a520362c6b986cde8171f746a725a0bf13037c6a7f5a54a4ba7], previous hash = [1b64d820db528ac48e2331150d2db0c65bf2b2a283c851dcd48d8
bef764b20b8], VERIFICATION PASSED
2018-03-08 02:57:11.014 IST [ledgerfsck] Verify -> INFO 0f8 block number [116]: previous hash matched
2018-03-08 02:57:11.014 IST [ledgerfsck] Verify -> INFO 0f9 Block [seq = 116], hash = [d3420103600e8ced3c54e6c26f947b60b7ffcfb5f6f65897638ff7f962b484ac], previous hash = [36d57ac3839e0a520362c6b986cde8171f746a725a0bf13037c6a
7f5a54a4ba7], VERIFICATION PASSED
2018-03-08 02:57:11.014 IST [ledgerfsck] Verify -> INFO 0fa block number [117]: previous hash matched
2018-03-08 02:57:11.014 IST [ledgerfsck] Verify -> INFO 0fb Block [seq = 117], hash = [c6fa496ab34af7481a4e0dacfceb7e5c1dce517880c81d3de3b4b59700e0d5e8], previous hash = [d3420103600e8ced3c54e6c26f947b60b7ffcfb5f6f65897638ff
7f962b484ac], VERIFICATION PASSED
2018-03-08 02:57:11.015 IST [ledgerfsck] Verify -> INFO 0fc block number [118]: previous hash matched
2018-03-08 02:57:11.015 IST [ledgerfsck] Verify -> INFO 0fd Block [seq = 118], hash = [0f4eb4ac5460570655881adf59ddf9984532ce4ae3455edd53a64893115910c8], previous hash = [c6fa496ab34af7481a4e0dacfceb7e5c1dce517880c81d3de3b4b
59700e0d5e8], VERIFICATION PASSED
2018-03-08 02:57:11.015 IST [ledgerfsck] Verify -> INFO 0fe block number [119]: previous hash matched
2018-03-08 02:57:11.015 IST [ledgerfsck] Verify -> INFO 0ff Block [seq = 119], hash = [33744d41f05eeda20ec96b5e4f252fcc50d72dfe87a6803e96266292e4f6e661], previous hash = [0f4eb4ac5460570655881adf59ddf9984532ce4ae3455edd53a64
893115910c8], VERIFICATION PASSED
 bartem@bartem1-mac  ~/golang/src/github.com/C0rwin/ledgerfsck   master ● 
$ ./ledgerfsck --channelName mychannel --mspID Org1MSP --mspPath ~/golang/src/github.com/hyperledger/fabric/tmp/crypto-config/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
2018-03-08 02:57:17.689 IST [ledgerfsck] ReadConfiguration -> INFO 001 channel name = mychannel
2018-03-08 02:57:17.689 IST [ledgerfsck] ReadConfiguration -> INFO 002 MSP folder path = /Users/bartem/golang/src/github.com/hyperledger/fabric/tmp/crypto-config/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
2018-03-08 02:57:17.689 IST [ledgerfsck] ReadConfiguration -> INFO 003 MSPID = Org1MSP
2018-03-08 02:57:17.689 IST [ledgerfsck] ReadConfiguration -> INFO 004 MSP type = bccsp
2018-03-08 02:57:17.709 IST [ledgermgmt] initialize -> INFO 005 Initializing ledger mgmt
2018-03-08 02:57:17.709 IST [kvledger] NewProvider -> INFO 006 Initializing ledger provider
2018-03-08 02:57:17.728 IST [kvledger] NewProvider -> INFO 007 ledger provider Initialized
2018-03-08 02:57:17.728 IST [ledgermgmt] initialize -> INFO 008 ledger mgmt initialized
2018-03-08 02:57:17.728 IST [ledgermgmt] OpenLedger -> INFO 009 Opening ledger with id = mychannel
2018-03-08 02:57:17.729 IST [ledgermgmt] OpenLedger -> INFO 00a Opened ledger with id = mychannel
2018-03-08 02:57:17.729 IST [ledgerfsck] GetLatestChannelConfigBundle -> INFO 00b reading configuration from state DB
2018-03-08 02:57:17.730 IST [ledgerfsck] GetLatestChannelConfigBundle -> INFO 00c initialize channel config bundle
2018-03-08 02:57:17.732 IST [ledgerfsck] GetLatestResourceConfigBundle -> INFO 00d check capabilities whenever there is support for resource tree
2018-03-08 02:57:17.732 IST [ledgerfsck] GetLatestResourceConfigBundle -> INFO 00e initialize resource config bundle
2018-03-08 02:57:17.732 IST [ledgerfsck] GetLatestResourceConfigBundle -> INFO 00f Initialize MSP Manager
2018-03-08 02:57:17.732 IST [ledgerfsck] Verify -> INFO 010 ledger height of channel mychannel, is 120
2018-03-08 02:57:17.732 IST [ledgerfsck] Verify -> INFO 011 block number [1]: previous hash matched
2018-03-08 02:57:17.732 IST [msp] DeserializeIdentity -> INFO 012 Obtaining identity
2018-03-08 02:57:17.733 IST [ledgerfsck] Verify -> INFO 013 Block [seq = 1], hash = [b4f21d966705efdade2be5823d2ab8196b50bac471ef2d03f0ea3e64d3bb3b84], previous hash = [df218993357f55c6dee14ed29d43815c36bdd06a3187cd609eccc0d0fd3624b6], VERIFICATION PASSED
2018-03-08 02:57:17.733 IST [ledgerfsck] Verify -> INFO 014 block number [2]: previous hash matched
2018-03-08 02:57:17.734 IST [ledgerfsck] Verify -> INFO 015 Block [seq = 2], hash = [bf5ad11abbd84ec32388774895e99ed0bc2d81b8917667dd0f191ab732f529b4], previous hash = [b4f21d966705efdade2be5823d2ab8196b50bac471ef2d03f0ea3e64d3bb3b84], VERIFICATION PASSED
2018-03-08 02:57:17.734 IST [ledgerfsck] Verify -> INFO 016 block number [3]: previous hash matched
2018-03-08 02:57:17.734 IST [ledgerfsck] Verify -> INFO 017 Block [seq = 3], hash = [3a2b1576cc7eb2f766597b43fcf00d76bed411aa9f05337b2456af7e8bc88f9a], previous hash = [bf5ad11abbd84ec32388774895e99ed0bc2d81b8917667dd0f191ab732f529b4], VERIFICATION PASSED
2018-03-08 02:57:17.735 IST [ledgerfsck] Verify -> INFO 018 block number [4]: previous hash matched
2018-03-08 02:57:17.735 IST [ledgerfsck] Verify -> INFO 019 Block [seq = 4], hash = [8a6d7e73dc4e1825f633f157ae44bab380ef941a27e4f500b0221ac1920c0eb8], previous hash = [3a2b1576cc7eb2f766597b43fcf00d76bed411aa9f05337b2456af7e8bc88f9a], VERIFICATION PASSED
2018-03-08 02:57:17.735 IST [ledgerfsck] Verify -> INFO 01a block number [5]: previous hash matched
2018-03-08 02:57:17.736 IST [ledgerfsck] Verify -> INFO 01b Block [seq = 5], hash = [f3dfba89c04b304ea257543613fb5c1748693da7d946827f4cc9afb6374e9c12], previous hash = [8a6d7e73dc4e1825f633f157ae44bab380ef941a27e4f500b0221ac1920c0eb8], VERIFICATION PASSED
2018-03-08 02:57:17.736 IST [ledgerfsck] Verify -> INFO 01c block number [6]: previous hash matched
2018-03-08 02:57:17.736 IST [ledgerfsck] Verify -> INFO 01d Block [seq = 6], hash = [feedbedca119c339466bdb4ba84ec59b34926bd265599cf244b28255b62a9c4e], previous hash = [f3dfba89c04b304ea257543613fb5c1748693da7d946827f4cc9afb6374e9c12], VERIFICATION PASSED
2018-03-08 02:57:17.737 IST [ledgerfsck] Verify -> INFO 01e block number [7]: previous hash matched
2018-03-08 02:57:17.737 IST [ledgerfsck] Verify -> INFO 01f Block [seq = 7], hash = [5e09a3da3f8e23015b1986e71ca349a6c2379b8fa0f1233e448a5696046d0e7c], previous hash = [feedbedca119c339466bdb4ba84ec59b34926bd265599cf244b28255b62a9c4e], VERIFICATION PASSED
2018-03-08 02:57:17.737 IST [ledgerfsck] Verify -> INFO 020 block number [8]: previous hash matched
2018-03-08 02:57:17.738 IST [ledgerfsck] Verify -> INFO 021 Block [seq = 8], hash = [c1b34f888f3caead21a59c7096d903515f7d1cfbb70274243cf1fa03440140c5], previous hash = [5e09a3da3f8e23015b1986e71ca349a6c2379b8fa0f1233e448a5696046d0e7c], VERIFICATION PASSED
2018-03-08 02:57:17.738 IST [ledgerfsck] Verify -> INFO 022 block number [9]: previous hash matched
2018-03-08 02:57:17.739 IST [ledgerfsck] Verify -> INFO 023 Block [seq = 9], hash = [fa785f9935a7c56ad8f52d1c1cdbce080cdb469d90eb9959cdb1be8a38f74b87], previous hash = [c1b34f888f3caead21a59c7096d903515f7d1cfbb70274243cf1fa03440140c5], VERIFICATION PASSED
2018-03-08 02:57:17.739 IST [ledgerfsck] Verify -> INFO 024 block number [10]: previous hash matched
2018-03-08 02:57:17.740 IST [ledgerfsck] Verify -> INFO 025 Block [seq = 10], hash = [19c14185e9d1dfc781d7c4a8b1df4f5edf16181780d9e3569680a7ef4cbeb76c], previous hash = [fa785f9935a7c56ad8f52d1c1cdbce080cdb469d90eb9959cdb1be8a38f74b87], VERIFICATION PASSED
2018-03-08 02:57:17.740 IST [ledgerfsck] Verify -> INFO 026 block number [11]: previous hash matched
2018-03-08 02:57:17.740 IST [ledgerfsck] Verify -> INFO 027 Block [seq = 11], hash = [97e529ddf2ced3cd76f67f78dafb1c524e7a75d63cdefd8449c981bbfee86233], previous hash = [19c14185e9d1dfc781d7c4a8b1df4f5edf16181780d9e3569680a7ef4cbeb76c], VERIFICATION PASSED
2018-03-08 02:57:17.740 IST [ledgerfsck] Verify -> INFO 028 block number [12]: previous hash matched
2018-03-08 02:57:17.741 IST [ledgerfsck] Verify -> INFO 029 Block [seq = 12], hash = [b95e40bc9a4796b9035a1463f6359c7da601e0a8874374e94c7e794c13b8681f], previous hash = [97e529ddf2ced3cd76f67f78dafb1c524e7a75d63cdefd8449c981bbfee86233], VERIFICATION PASSED
2018-03-08 02:57:17.741 IST [ledgerfsck] Verify -> INFO 02a block number [13]: previous hash matched
2018-03-08 02:57:17.741 IST [ledgerfsck] Verify -> INFO 02b Block [seq = 13], hash = [01a5ea9354af47d76ba60f03e563d34c55178c0a62d798271e01f44c99ce6585], previous hash = [b95e40bc9a4796b9035a1463f6359c7da601e0a8874374e94c7e794c13b8681f], VERIFICATION PASSED
2018-03-08 02:57:17.741 IST [ledgerfsck] Verify -> INFO 02c block number [14]: previous hash matched
2018-03-08 02:57:17.742 IST [ledgerfsck] Verify -> INFO 02d Block [seq = 14], hash = [385d0a5abda35f5d8abcb143d84ec55765ff5f67f7f25a87be479ad340c1b8b2], previous hash = [01a5ea9354af47d76ba60f03e563d34c55178c0a62d798271e01f44c99ce6585], VERIFICATION PASSED
2018-03-08 02:57:17.742 IST [ledgerfsck] Verify -> INFO 02e block number [15]: previous hash matched
2018-03-08 02:57:17.742 IST [ledgerfsck] Verify -> INFO 02f Block [seq = 15], hash = [dbd0ee5ad4e4371a362c4853c57e06a2b7f57e49672f4ec848093cb17b7c340c], previous hash = [385d0a5abda35f5d8abcb143d84ec55765ff5f67f7f25a87be479ad340c1b8b2], VERIFICATION PASSED
2018-03-08 02:57:17.743 IST [ledgerfsck] Verify -> INFO 030 block number [16]: previous hash matched
2018-03-08 02:57:17.743 IST [ledgerfsck] Verify -> INFO 031 Block [seq = 16], hash = [6943e1c4212972b24ed015d6adec176c7095365bd18f7ecdef5d1d211c8994d0], previous hash = [dbd0ee5ad4e4371a362c4853c57e06a2b7f57e49672f4ec848093cb17b7c340c], VERIFICATION PASSED
2018-03-08 02:57:17.743 IST [ledgerfsck] Verify -> INFO 032 block number [17]: previous hash matched
2018-03-08 02:57:17.744 IST [ledgerfsck] Verify -> INFO 033 Block [seq = 17], hash = [094600cc256a2f534362619ac05b6b58b6b7f6680fee81e5c378d89a8e3da2b5], previous hash = [6943e1c4212972b24ed015d6adec176c7095365bd18f7ecdef5d1d211c8994d0], VERIFICATION PASSED
2018-03-08 02:57:17.744 IST [ledgerfsck] Verify -> INFO 034 block number [18]: previous hash matched
2018-03-08 02:57:17.745 IST [ledgerfsck] Verify -> INFO 035 Block [seq = 18], hash = [fa1f93ad56c761d08c7fce71f2087db4e6fd4de3648fe02bf1accf53ecfc6da8], previous hash = [094600cc256a2f534362619ac05b6b58b6b7f6680fee81e5c378d89a8e3da2b5], VERIFICATION PASSED
2018-03-08 02:57:17.745 IST [ledgerfsck] Verify -> INFO 036 block number [19]: previous hash matched
2018-03-08 02:57:17.745 IST [ledgerfsck] Verify -> INFO 037 Block [seq = 19], hash = [c2a2ef321f439a7420168c54d9adf250e820870a837db1e64eb095041f8135d3], previous hash = [fa1f93ad56c761d08c7fce71f2087db4e6fd4de3648fe02bf1accf53ecfc6da8], VERIFICATION PASSED
2018-03-08 02:57:17.746 IST [ledgerfsck] Verify -> INFO 038 block number [20]: previous hash matched
2018-03-08 02:57:17.746 IST [ledgerfsck] Verify -> INFO 039 Block [seq = 20], hash = [f4a47b61832427e8b8e9e35990d91c0596d61c38138c2de105ef9620df2f2a19], previous hash = [c2a2ef321f439a7420168c54d9adf250e820870a837db1e64eb095041f8135d3], VERIFICATION PASSED
2018-03-08 02:57:17.746 IST [ledgerfsck] Verify -> INFO 03a block number [21]: previous hash matched
2018-03-08 02:57:17.747 IST [ledgerfsck] Verify -> INFO 03b Block [seq = 21], hash = [2ac5024423155f68bc598486594a981e8863a152fd3ce1ee88d80bf3e9f024d5], previous hash = [f4a47b61832427e8b8e9e35990d91c0596d61c38138c2de105ef9620df2f2a19], VERIFICATION PASSED
2018-03-08 02:57:17.747 IST [ledgerfsck] Verify -> INFO 03c block number [22]: previous hash matched
2018-03-08 02:57:17.747 IST [ledgerfsck] Verify -> INFO 03d Block [seq = 22], hash = [7b042659c419b19d68671e7db6e7e6614f41e07fc875293454aab4daf83872f2], previous hash = [2ac5024423155f68bc598486594a981e8863a152fd3ce1ee88d80bf3e9f024d5], VERIFICATION PASSED
2018-03-08 02:57:17.747 IST [ledgerfsck] Verify -> INFO 03e block number [23]: previous hash matched
2018-03-08 02:57:17.752 IST [ledgerfsck] Verify -> INFO 03f Block [seq = 23], hash = [f91a792dbc76307aa156ce37dffdbc0dbb27492aa234bba14b54602b68b65539], previous hash = [7b042659c419b19d68671e7db6e7e6614f41e07fc875293454aab4daf83872f2], VERIFICATION PASSED
2018-03-08 02:57:17.752 IST [ledgerfsck] Verify -> INFO 040 block number [24]: previous hash matched
2018-03-08 02:57:17.752 IST [ledgerfsck] Verify -> INFO 041 Block [seq = 24], hash = [c162fb69dc09d05ff2273a8a4871e36a6894f6dd4557935abc51997232136484], previous hash = [f91a792dbc76307aa156ce37dffdbc0dbb27492aa234bba14b54602b68b65539], VERIFICATION PASSED
2018-03-08 02:57:17.752 IST [ledgerfsck] Verify -> INFO 042 block number [25]: previous hash matched
2018-03-08 02:57:17.753 IST [ledgerfsck] Verify -> INFO 043 Block [seq = 25], hash = [a1efda1a7aed65110acb19c9a5636b6ce7f0015b13e97e9310402516cf5689af], previous hash = [c162fb69dc09d05ff2273a8a4871e36a6894f6dd4557935abc51997232136484], VERIFICATION PASSED
2018-03-08 02:57:17.753 IST [ledgerfsck] Verify -> INFO 044 block number [26]: previous hash matched
2018-03-08 02:57:17.754 IST [ledgerfsck] Verify -> INFO 045 Block [seq = 26], hash = [1194b7539337159c1c9cfdeb970b9675f27ff25b165bdba6775c5323e9c37cf4], previous hash = [a1efda1a7aed65110acb19c9a5636b6ce7f0015b13e97e9310402516cf5689af], VERIFICATION PASSED
2018-03-08 02:57:17.754 IST [ledgerfsck] Verify -> INFO 046 block number [27]: previous hash matched
2018-03-08 02:57:17.754 IST [ledgerfsck] Verify -> INFO 047 Block [seq = 27], hash = [abdb2998540a0e9ae4d3dde998bf52f0bdbe117827964f67825b27c969b46010], previous hash = [1194b7539337159c1c9cfdeb970b9675f27ff25b165bdba6775c5323e9c37cf4], VERIFICATION PASSED
2018-03-08 02:57:17.754 IST [ledgerfsck] Verify -> INFO 048 block number [28]: previous hash matched
2018-03-08 02:57:17.755 IST [ledgerfsck] Verify -> INFO 049 Block [seq = 28], hash = [98ece8b963b9fdf81f5a8eb5f59eafdf39398ef67f99f1e6a07fd1b2df52af28], previous hash = [abdb2998540a0e9ae4d3dde998bf52f0bdbe117827964f67825b27c969b46010], VERIFICATION PASSED
2018-03-08 02:57:17.755 IST [ledgerfsck] Verify -> INFO 04a block number [29]: previous hash matched
2018-03-08 02:57:17.755 IST [ledgerfsck] Verify -> INFO 04b Block [seq = 29], hash = [4409cfbe09ffcd748ffe921d9523c157157531e86e9a3f6a59ec65887d91bac4], previous hash = [98ece8b963b9fdf81f5a8eb5f59eafdf39398ef67f99f1e6a07fd1b2df52af28], VERIFICATION PASSED
2018-03-08 02:57:17.755 IST [ledgerfsck] Verify -> INFO 04c block number [30]: previous hash matched
2018-03-08 02:57:17.756 IST [ledgerfsck] Verify -> INFO 04d Block [seq = 30], hash = [5a06b4a7d31bf0c61f7ff267a28ad9ce919990ad244d700e1e34ac8d876e51d8], previous hash = [4409cfbe09ffcd748ffe921d9523c157157531e86e9a3f6a59ec65887d91bac4], VERIFICATION PASSED
2018-03-08 02:57:17.756 IST [ledgerfsck] Verify -> INFO 04e block number [31]: previous hash matched
2018-03-08 02:57:17.756 IST [ledgerfsck] Verify -> INFO 04f Block [seq = 31], hash = [c5e4c5d056f48f129c252259dec99b1e579693ac8b9db674b5a3a060e6b363a6], previous hash = [5a06b4a7d31bf0c61f7ff267a28ad9ce919990ad244d700e1e34ac8d876e51d8], VERIFICATION PASSED
2018-03-08 02:57:17.756 IST [ledgerfsck] Verify -> INFO 050 block number [32]: previous hash matched
2018-03-08 02:57:17.756 IST [ledgerfsck] Verify -> INFO 051 Block [seq = 32], hash = [4c191392c7b8b10533cd0c945864a8d95248ae70e9427d63f5842ec472e75925], previous hash = [c5e4c5d056f48f129c252259dec99b1e579693ac8b9db674b5a3a060e6b363a6], VERIFICATION PASSED
2018-03-08 02:57:17.757 IST [ledgerfsck] Verify -> INFO 052 block number [33]: previous hash matched
2018-03-08 02:57:17.757 IST [ledgerfsck] Verify -> INFO 053 Block [seq = 33], hash = [ca5cac7626777dd33808a7930e5a39daa7f6fb4ba27e8e808461e81fa1b02395], previous hash = [4c191392c7b8b10533cd0c945864a8d95248ae70e9427d63f5842ec472e75925], VERIFICATION PASSED
2018-03-08 02:57:17.757 IST [ledgerfsck] Verify -> INFO 054 block number [34]: previous hash matched
2018-03-08 02:57:17.757 IST [ledgerfsck] Verify -> INFO 055 Block [seq = 34], hash = [4253704839936e0123c63de81f05b1a8e4016cb685f6896a13985c439839a126], previous hash = [ca5cac7626777dd33808a7930e5a39daa7f6fb4ba27e8e808461e81fa1b02395], VERIFICATION PASSED
2018-03-08 02:57:17.758 IST [ledgerfsck] Verify -> INFO 056 block number [35]: previous hash matched
2018-03-08 02:57:17.758 IST [ledgerfsck] Verify -> INFO 057 Block [seq = 35], hash = [59f404222f34a31a5fe213e600e0e90f6d617d7436c9b95690f58b6b6853369c], previous hash = [4253704839936e0123c63de81f05b1a8e4016cb685f6896a13985c439839a126], VERIFICATION PASSED
2018-03-08 02:57:17.758 IST [ledgerfsck] Verify -> INFO 058 block number [36]: previous hash matched
2018-03-08 02:57:17.758 IST [ledgerfsck] Verify -> INFO 059 Block [seq = 36], hash = [0e7b67732ed6a65c6b20d1cd74115a05c5086557bf3f21ed071f8527afef75d1], previous hash = [59f404222f34a31a5fe213e600e0e90f6d617d7436c9b95690f58b6b6853369c], VERIFICATION PASSED
2018-03-08 02:57:17.758 IST [ledgerfsck] Verify -> INFO 05a block number [37]: previous hash matched
2018-03-08 02:57:17.759 IST [ledgerfsck] Verify -> INFO 05b Block [seq = 37], hash = [bee9e5c76bc5908625be30dd53246cadfd8d158db9235bb08cc065f824caf79d], previous hash = [0e7b67732ed6a65c6b20d1cd74115a05c5086557bf3f21ed071f8527afef75d1], VERIFICATION PASSED
2018-03-08 02:57:17.759 IST [ledgerfsck] Verify -> INFO 05c block number [38]: previous hash matched
2018-03-08 02:57:17.759 IST [ledgerfsck] Verify -> INFO 05d Block [seq = 38], hash = [f576123a523b3dcd5f0d391bbd0273318c52473e91fbec2a9baca2d77eb41418], previous hash = [bee9e5c76bc5908625be30dd53246cadfd8d158db9235bb08cc065f824caf79d], VERIFICATION PASSED
2018-03-08 02:57:17.759 IST [ledgerfsck] Verify -> INFO 05e block number [39]: previous hash matched
2018-03-08 02:57:17.760 IST [ledgerfsck] Verify -> INFO 05f Block [seq = 39], hash = [4a58fddca6999703d49430245413377a757f1b3c07a8c62b9e2ea80781ad5635], previous hash = [f576123a523b3dcd5f0d391bbd0273318c52473e91fbec2a9baca2d77eb41418], VERIFICATION PASSED
2018-03-08 02:57:17.760 IST [ledgerfsck] Verify -> INFO 060 block number [40]: previous hash matched
2018-03-08 02:57:17.760 IST [ledgerfsck] Verify -> INFO 061 Block [seq = 40], hash = [f35433595162bd964d0a833ccbe01f34f7f7adf78eba504c95dbe62fa845d787], previous hash = [4a58fddca6999703d49430245413377a757f1b3c07a8c62b9e2ea80781ad5635], VERIFICATION PASSED
2018-03-08 02:57:17.760 IST [ledgerfsck] Verify -> INFO 062 block number [41]: previous hash matched
2018-03-08 02:57:17.761 IST [ledgerfsck] Verify -> INFO 063 Block [seq = 41], hash = [bbe86863af87883b9de6d4300d900b25d38de47b92709cc301a616f06d33c23e], previous hash = [f35433595162bd964d0a833ccbe01f34f7f7adf78eba504c95dbe62fa845d787], VERIFICATION PASSED
2018-03-08 02:57:17.761 IST [ledgerfsck] Verify -> INFO 064 block number [42]: previous hash matched
2018-03-08 02:57:17.762 IST [ledgerfsck] Verify -> INFO 065 Block [seq = 42], hash = [0feb604e2602ef3e7539125747d0a3aef456cdaade9ca82f5c49e283a8719108], previous hash = [bbe86863af87883b9de6d4300d900b25d38de47b92709cc301a616f06d33c23e], VERIFICATION PASSED
2018-03-08 02:57:17.762 IST [ledgerfsck] Verify -> INFO 066 block number [43]: previous hash matched
2018-03-08 02:57:17.762 IST [ledgerfsck] Verify -> INFO 067 Block [seq = 43], hash = [19adc48355397f6b3308f628683e0f68658a9446208a735a8eb5b6aeef402932], previous hash = [0feb604e2602ef3e7539125747d0a3aef456cdaade9ca82f5c49e283a8719108], VERIFICATION PASSED
2018-03-08 02:57:17.763 IST [ledgerfsck] Verify -> INFO 068 block number [44]: previous hash matched
2018-03-08 02:57:17.763 IST [ledgerfsck] Verify -> INFO 069 Block [seq = 44], hash = [fabcc4caafc9abcd62cfff7677983781d1d0263592e173fc3b815c0fcd170f55], previous hash = [19adc48355397f6b3308f628683e0f68658a9446208a735a8eb5b6aeef402932], VERIFICATION PASSED
2018-03-08 02:57:17.763 IST [ledgerfsck] Verify -> INFO 06a block number [45]: previous hash matched
2018-03-08 02:57:17.764 IST [ledgerfsck] Verify -> INFO 06b Block [seq = 45], hash = [88f30451aaed29328bce6951de37eb4146c81133ad78897d9fe95a5181a3c57b], previous hash = [fabcc4caafc9abcd62cfff7677983781d1d0263592e173fc3b815c0fcd170f55], VERIFICATION PASSED
2018-03-08 02:57:17.764 IST [ledgerfsck] Verify -> INFO 06c block number [46]: previous hash matched
2018-03-08 02:57:17.764 IST [ledgerfsck] Verify -> INFO 06d Block [seq = 46], hash = [3bc94de05e93d12bd98e5a7abb21028b6b2fbcffb21930996bbce5f4553f473f], previous hash = [88f30451aaed29328bce6951de37eb4146c81133ad78897d9fe95a5181a3c57b], VERIFICATION PASSED
2018-03-08 02:57:17.765 IST [ledgerfsck] Verify -> INFO 06e block number [47]: previous hash matched
2018-03-08 02:57:17.765 IST [ledgerfsck] Verify -> INFO 06f Block [seq = 47], hash = [4410697aff88b18680e10180e7a9569326d38df63e5a917e815e540ba89a246e], previous hash = [3bc94de05e93d12bd98e5a7abb21028b6b2fbcffb21930996bbce5f4553f473f], VERIFICATION PASSED
2018-03-08 02:57:17.765 IST [ledgerfsck] Verify -> INFO 070 block number [48]: previous hash matched
2018-03-08 02:57:17.766 IST [ledgerfsck] Verify -> INFO 071 Block [seq = 48], hash = [c70ec26d513eab674aa3f39399a011743636af09a011e953d99d785862f6ce59], previous hash = [4410697aff88b18680e10180e7a9569326d38df63e5a917e815e540ba89a246e], VERIFICATION PASSED
2018-03-08 02:57:17.766 IST [ledgerfsck] Verify -> INFO 072 block number [49]: previous hash matched
2018-03-08 02:57:17.766 IST [ledgerfsck] Verify -> INFO 073 Block [seq = 49], hash = [d58a633faf364550b7cd39be3752bd8ee314915f2b81c8211c5f683a65e8783b], previous hash = [c70ec26d513eab674aa3f39399a011743636af09a011e953d99d785862f6ce59], VERIFICATION PASSED
2018-03-08 02:57:17.767 IST [ledgerfsck] Verify -> INFO 074 block number [50]: previous hash matched
2018-03-08 02:57:17.767 IST [ledgerfsck] Verify -> INFO 075 Block [seq = 50], hash = [46c90c450fd52d96dc873b3de9e7d6c22ce34ed72befbc527db2b44013aabcb8], previous hash = [d58a633faf364550b7cd39be3752bd8ee314915f2b81c8211c5f683a65e8783b], VERIFICATION PASSED
2018-03-08 02:57:17.767 IST [ledgerfsck] Verify -> INFO 076 block number [51]: previous hash matched
2018-03-08 02:57:17.768 IST [ledgerfsck] Verify -> INFO 077 Block [seq = 51], hash = [46e2921b6415b18f469181495a48fceeebdd3b88db2382d527c3a57ede4f2759], previous hash = [46c90c450fd52d96dc873b3de9e7d6c22ce34ed72befbc527db2b44013aabcb8], VERIFICATION PASSED
2018-03-08 02:57:17.768 IST [ledgerfsck] Verify -> INFO 078 block number [52]: previous hash matched
2018-03-08 02:57:17.768 IST [ledgerfsck] Verify -> INFO 079 Block [seq = 52], hash = [69eee6df67f2bfc634093b290cf520d3691834c6b50d749e65ff8db4f164c416], previous hash = [46e2921b6415b18f469181495a48fceeebdd3b88db2382d527c3a57ede4f2759], VERIFICATION PASSED
2018-03-08 02:57:17.769 IST [ledgerfsck] Verify -> INFO 07a block number [53]: previous hash matched
2018-03-08 02:57:17.769 IST [ledgerfsck] Verify -> INFO 07b Block [seq = 53], hash = [4dab4a702708fef3bf50319ae2a1f54447bc24a1f6f0e3225a942837b907433b], previous hash = [69eee6df67f2bfc634093b290cf520d3691834c6b50d749e65ff8db4f164c416], VERIFICATION PASSED
2018-03-08 02:57:17.769 IST [ledgerfsck] Verify -> INFO 07c block number [54]: previous hash matched
2018-03-08 02:57:17.770 IST [ledgerfsck] Verify -> INFO 07d Block [seq = 54], hash = [14861fbf6c9c68ce775e494ef0f6fefea0bc4c733444c548f6940177a92f5ee2], previous hash = [4dab4a702708fef3bf50319ae2a1f54447bc24a1f6f0e3225a942837b907433b], VERIFICATION PASSED
2018-03-08 02:57:17.773 IST [ledgerfsck] Verify -> INFO 07e block number [55]: previous hash matched
2018-03-08 02:57:17.773 IST [ledgerfsck] Verify -> INFO 07f Block [seq = 55], hash = [d45aa2ecfe82bfa0f2dae22dc8dc77a6c640fc74cf078d7f0e039e74d4e72f27], previous hash = [14861fbf6c9c68ce775e494ef0f6fefea0bc4c733444c548f6940177a92f5ee2], VERIFICATION PASSED
2018-03-08 02:57:17.774 IST [ledgerfsck] Verify -> INFO 080 block number [56]: previous hash matched
2018-03-08 02:57:17.774 IST [ledgerfsck] Verify -> INFO 081 Block [seq = 56], hash = [d20eae3dce351bfcc185589242c9aecfb0a32787d613f61f1fd3cade4f819056], previous hash = [d45aa2ecfe82bfa0f2dae22dc8dc77a6c640fc74cf078d7f0e039e74d4e72f27], VERIFICATION PASSED
2018-03-08 02:57:17.775 IST [ledgerfsck] Verify -> INFO 082 block number [57]: previous hash matched
2018-03-08 02:57:17.775 IST [ledgerfsck] Verify -> INFO 083 Block [seq = 57], hash = [15bc739a87823dbf36ed65b8d7925dc87a75c4725d758c7f6e608e6c24732cf3], previous hash = [d20eae3dce351bfcc185589242c9aecfb0a32787d613f61f1fd3cade4f819056], VERIFICATION PASSED
2018-03-08 02:57:17.775 IST [ledgerfsck] Verify -> INFO 084 block number [58]: previous hash matched
2018-03-08 02:57:17.776 IST [ledgerfsck] Verify -> INFO 085 Block [seq = 58], hash = [b0b7026fcc97d9121d7ed90cb08dae56ae330bbf873110d6bf06cc6a81510c6a], previous hash = [15bc739a87823dbf36ed65b8d7925dc87a75c4725d758c7f6e608e6c24732cf3], VERIFICATION PASSED
2018-03-08 02:57:17.776 IST [ledgerfsck] Verify -> INFO 086 block number [59]: previous hash matched
2018-03-08 02:57:17.776 IST [ledgerfsck] Verify -> INFO 087 Block [seq = 59], hash = [d7f6b8f7bb08662313e6a5ec1db43044d8083d9adddd563955e9fd520a9fc9d3], previous hash = [b0b7026fcc97d9121d7ed90cb08dae56ae330bbf873110d6bf06cc6a81510c6a], VERIFICATION PASSED
2018-03-08 02:57:17.777 IST [ledgerfsck] Verify -> INFO 088 block number [60]: previous hash matched
2018-03-08 02:57:17.777 IST [ledgerfsck] Verify -> INFO 089 Block [seq = 60], hash = [f670027e16e475ab127bacb9356acbf865991c4d2fe1ee9aa1b3e6dc12651a03], previous hash = [d7f6b8f7bb08662313e6a5ec1db43044d8083d9adddd563955e9fd520a9fc9d3], VERIFICATION PASSED
2018-03-08 02:57:17.777 IST [ledgerfsck] Verify -> INFO 08a block number [61]: previous hash matched
2018-03-08 02:57:17.778 IST [ledgerfsck] Verify -> INFO 08b Block [seq = 61], hash = [babba33bb149814e26fe67b217fa5cd177d9057745a84570180b987e5bc40910], previous hash = [f670027e16e475ab127bacb9356acbf865991c4d2fe1ee9aa1b3e6dc12651a03], VERIFICATION PASSED
2018-03-08 02:57:17.778 IST [ledgerfsck] Verify -> INFO 08c block number [62]: previous hash matched
2018-03-08 02:57:17.778 IST [ledgerfsck] Verify -> INFO 08d Block [seq = 62], hash = [ea6add9c9a0c29e43133bc333a7d88f82065f225f20efbf125e33d4e5060dbf9], previous hash = [babba33bb149814e26fe67b217fa5cd177d9057745a84570180b987e5bc40910], VERIFICATION PASSED
2018-03-08 02:57:17.779 IST [ledgerfsck] Verify -> INFO 08e block number [63]: previous hash matched
2018-03-08 02:57:17.779 IST [ledgerfsck] Verify -> INFO 08f Block [seq = 63], hash = [b3143d138f6b4e89dbc8729027bf8b57f8da9cbaecee47424a7e573ce16c7074], previous hash = [ea6add9c9a0c29e43133bc333a7d88f82065f225f20efbf125e33d4e5060dbf9], VERIFICATION PASSED
2018-03-08 02:57:17.779 IST [ledgerfsck] Verify -> INFO 090 block number [64]: previous hash matched
2018-03-08 02:57:17.780 IST [ledgerfsck] Verify -> INFO 091 Block [seq = 64], hash = [21335a1ab74682754041e06332fcb6157a99608399626e4405d301aae85c2381], previous hash = [b3143d138f6b4e89dbc8729027bf8b57f8da9cbaecee47424a7e573ce16c7074], VERIFICATION PASSED
2018-03-08 02:57:17.780 IST [ledgerfsck] Verify -> INFO 092 block number [65]: previous hash matched
2018-03-08 02:57:17.780 IST [ledgerfsck] Verify -> INFO 093 Block [seq = 65], hash = [1018e603e242ef78efcb72c26d420b1a271989e2e223771d8b2c69ddab52e65b], previous hash = [21335a1ab74682754041e06332fcb6157a99608399626e4405d301aae85c2381], VERIFICATION PASSED
2018-03-08 02:57:17.780 IST [ledgerfsck] Verify -> INFO 094 block number [66]: previous hash matched
2018-03-08 02:57:17.781 IST [ledgerfsck] Verify -> INFO 095 Block [seq = 66], hash = [4208c154d60e97ae8b980f96eae19564e430eb1464f07c62358a593ac44017b6], previous hash = [1018e603e242ef78efcb72c26d420b1a271989e2e223771d8b2c69ddab52e65b], VERIFICATION PASSED
2018-03-08 02:57:17.781 IST [ledgerfsck] Verify -> INFO 096 block number [67]: previous hash matched
2018-03-08 02:57:17.782 IST [ledgerfsck] Verify -> INFO 097 Block [seq = 67], hash = [06172077eb62b87f9e43c8c333c8232bb4a5da1a1ca9c6b2fdb5f68239f591e8], previous hash = [4208c154d60e97ae8b980f96eae19564e430eb1464f07c62358a593ac44017b6], VERIFICATION PASSED
2018-03-08 02:57:17.782 IST [ledgerfsck] Verify -> INFO 098 block number [68]: previous hash matched
2018-03-08 02:57:17.782 IST [ledgerfsck] Verify -> INFO 099 Block [seq = 68], hash = [1c1d0a4553301078ca72d7ddf95add36be1134fc94d55fdc48ad37c6c3aca4e7], previous hash = [06172077eb62b87f9e43c8c333c8232bb4a5da1a1ca9c6b2fdb5f68239f591e8], VERIFICATION PASSED
2018-03-08 02:57:17.783 IST [ledgerfsck] Verify -> INFO 09a block number [69]: previous hash matched
2018-03-08 02:57:17.783 IST [ledgerfsck] Verify -> INFO 09b Block [seq = 69], hash = [6803ed3ae35f22a5a32aebb5387e256bf006ea74dd8b53291dd60b32a2e66081], previous hash = [1c1d0a4553301078ca72d7ddf95add36be1134fc94d55fdc48ad37c6c3aca4e7], VERIFICATION PASSED
2018-03-08 02:57:17.783 IST [ledgerfsck] Verify -> INFO 09c block number [70]: previous hash matched
2018-03-08 02:57:17.784 IST [ledgerfsck] Verify -> INFO 09d Block [seq = 70], hash = [9f02d9d57c16d14809871f5007c0fe6b5a55ae7ec740747ce6f1666e95438f18], previous hash = [6803ed3ae35f22a5a32aebb5387e256bf006ea74dd8b53291dd60b32a2e66081], VERIFICATION PASSED
2018-03-08 02:57:17.785 IST [ledgerfsck] Verify -> INFO 09e block number [71]: previous hash matched
2018-03-08 02:57:17.785 IST [ledgerfsck] Verify -> INFO 09f Block [seq = 71], hash = [8ce9932595a2a5949189a0dc18845a04ca24aacaf00508e5e69a7a6580de4981], previous hash = [9f02d9d57c16d14809871f5007c0fe6b5a55ae7ec740747ce6f1666e95438f18], VERIFICATION PASSED
2018-03-08 02:57:17.785 IST [ledgerfsck] Verify -> INFO 0a0 block number [72]: previous hash matched
2018-03-08 02:57:17.786 IST [ledgerfsck] Verify -> INFO 0a1 Block [seq = 72], hash = [d6e70f0bb41501aff0f4430480db76dd445bc6fb3bc588acc44721462e6cd750], previous hash = [8ce9932595a2a5949189a0dc18845a04ca24aacaf00508e5e69a7a6580de4981], VERIFICATION PASSED
2018-03-08 02:57:17.786 IST [ledgerfsck] Verify -> INFO 0a2 block number [73]: previous hash matched
2018-03-08 02:57:17.786 IST [ledgerfsck] Verify -> INFO 0a3 Block [seq = 73], hash = [2f05da54cca8503362da2249b7b13d39eade2a437d3a3d5b156c0aa80efc8d2b], previous hash = [d6e70f0bb41501aff0f4430480db76dd445bc6fb3bc588acc44721462e6cd750], VERIFICATION PASSED
2018-03-08 02:57:17.786 IST [ledgerfsck] Verify -> INFO 0a4 block number [74]: previous hash matched
2018-03-08 02:57:17.787 IST [ledgerfsck] Verify -> INFO 0a5 Block [seq = 74], hash = [6b6fc2a2a3ede5d573469c399ecb5a43aea44bdd4d24912bfe94c692b45a10cd], previous hash = [2f05da54cca8503362da2249b7b13d39eade2a437d3a3d5b156c0aa80efc8d2b], VERIFICATION PASSED
2018-03-08 02:57:17.787 IST [ledgerfsck] Verify -> INFO 0a6 block number [75]: previous hash matched
2018-03-08 02:57:17.787 IST [ledgerfsck] Verify -> INFO 0a7 Block [seq = 75], hash = [cba48915a58ab5e951a1785c390d5b6a79667f3ef59126d7955e5617f8c6fdb5], previous hash = [6b6fc2a2a3ede5d573469c399ecb5a43aea44bdd4d24912bfe94c692b45a10cd], VERIFICATION PASSED
2018-03-08 02:57:17.787 IST [ledgerfsck] Verify -> INFO 0a8 block number [76]: previous hash matched
2018-03-08 02:57:17.788 IST [ledgerfsck] Verify -> INFO 0a9 Block [seq = 76], hash = [cb83bd8430262fc1072a37ad800582d1880c2744ee7f7ee5546ff9d44e05df6e], previous hash = [cba48915a58ab5e951a1785c390d5b6a79667f3ef59126d7955e5617f8c6fdb5], VERIFICATION PASSED
2018-03-08 02:57:17.788 IST [ledgerfsck] Verify -> INFO 0aa block number [77]: previous hash matched
2018-03-08 02:57:17.788 IST [ledgerfsck] Verify -> INFO 0ab Block [seq = 77], hash = [a1092ea1fc5461205194628d0c62bff32533a00d90723df273da5f014d54da04], previous hash = [cb83bd8430262fc1072a37ad800582d1880c2744ee7f7ee5546ff9d44e05df6e], VERIFICATION PASSED
2018-03-08 02:57:17.788 IST [ledgerfsck] Verify -> INFO 0ac block number [78]: previous hash matched
2018-03-08 02:57:17.789 IST [ledgerfsck] Verify -> INFO 0ad Block [seq = 78], hash = [274ff3dbcf13e9b47efabbdc8efe8cfc45ba5166e260d479ce72d4e9d059c74d], previous hash = [a1092ea1fc5461205194628d0c62bff32533a00d90723df273da5f014d54da04], VERIFICATION PASSED
2018-03-08 02:57:17.789 IST [ledgerfsck] Verify -> INFO 0ae block number [79]: previous hash matched
2018-03-08 02:57:17.789 IST [ledgerfsck] Verify -> INFO 0af Block [seq = 79], hash = [75bdd3965699b0c9be52a1811b676e27c2b15d6eed2c97db07adc563df1c951e], previous hash = [274ff3dbcf13e9b47efabbdc8efe8cfc45ba5166e260d479ce72d4e9d059c74d], VERIFICATION PASSED
2018-03-08 02:57:17.789 IST [ledgerfsck] Verify -> INFO 0b0 block number [80]: previous hash matched
2018-03-08 02:57:17.790 IST [ledgerfsck] Verify -> INFO 0b1 Block [seq = 80], hash = [f8c2813bfc91ff294a10c966d35722566ad6ddfc7981b3ff3cc7826362f6637d], previous hash = [75bdd3965699b0c9be52a1811b676e27c2b15d6eed2c97db07adc563df1c951e], VERIFICATION PASSED
2018-03-08 02:57:17.790 IST [ledgerfsck] Verify -> INFO 0b2 block number [81]: previous hash matched
2018-03-08 02:57:17.790 IST [ledgerfsck] Verify -> INFO 0b3 Block [seq = 81], hash = [b5d90a413fd32ad45cc7759476cf1358a86eb0a3dea40a75efa9ee67bb9b7b3e], previous hash = [f8c2813bfc91ff294a10c966d35722566ad6ddfc7981b3ff3cc7826362f6637d], VERIFICATION PASSED
2018-03-08 02:57:17.790 IST [ledgerfsck] Verify -> INFO 0b4 block number [82]: previous hash matched
2018-03-08 02:57:17.791 IST [ledgerfsck] Verify -> INFO 0b5 Block [seq = 82], hash = [339a3c8f5287b9c9f83612b96203619d4571c98175df8542a22c2f93afa6f81c], previous hash = [b5d90a413fd32ad45cc7759476cf1358a86eb0a3dea40a75efa9ee67bb9b7b3e], VERIFICATION PASSED
2018-03-08 02:57:17.791 IST [ledgerfsck] Verify -> INFO 0b6 block number [83]: previous hash matched
2018-03-08 02:57:17.791 IST [ledgerfsck] Verify -> INFO 0b7 Block [seq = 83], hash = [c7d09ee566e1bdcfd6adc1654c9b9f5d8a434f059a3dcb61db98853bdcfeca73], previous hash = [339a3c8f5287b9c9f83612b96203619d4571c98175df8542a22c2f93afa6f81c], VERIFICATION PASSED
2018-03-08 02:57:17.791 IST [ledgerfsck] Verify -> INFO 0b8 block number [84]: previous hash matched
2018-03-08 02:57:17.792 IST [ledgerfsck] Verify -> INFO 0b9 Block [seq = 84], hash = [65bb44fe8bdf1a95e20e061b939f18d84aff35348a3896afe6cd6efb55d807a5], previous hash = [c7d09ee566e1bdcfd6adc1654c9b9f5d8a434f059a3dcb61db98853bdcfeca73], VERIFICATION PASSED
2018-03-08 02:57:17.792 IST [ledgerfsck] Verify -> INFO 0ba block number [85]: previous hash matched
2018-03-08 02:57:17.792 IST [ledgerfsck] Verify -> INFO 0bb Block [seq = 85], hash = [dd85438b27465b63e66d490a65eb86cb17d1db8cd0e130b5e0d1e5526b1aa9dc], previous hash = [65bb44fe8bdf1a95e20e061b939f18d84aff35348a3896afe6cd6efb55d807a5], VERIFICATION PASSED
2018-03-08 02:57:17.792 IST [ledgerfsck] Verify -> INFO 0bc block number [86]: previous hash matched
2018-03-08 02:57:17.793 IST [ledgerfsck] Verify -> INFO 0bd Block [seq = 86], hash = [70a8db81743bf9ed8553287b0063cdd4c0fdbac4723a26588174d6c5d608d23d], previous hash = [dd85438b27465b63e66d490a65eb86cb17d1db8cd0e130b5e0d1e5526b1aa9dc], VERIFICATION PASSED
2018-03-08 02:57:17.793 IST [ledgerfsck] Verify -> INFO 0be block number [87]: previous hash matched
2018-03-08 02:57:17.797 IST [ledgerfsck] Verify -> INFO 0bf Block [seq = 87], hash = [b0fe399e33b9ca507402eef8650e8dbe35c97bb859f29e589a95c48e4299fdcb], previous hash = [70a8db81743bf9ed8553287b0063cdd4c0fdbac4723a26588174d6c5d608d23d], VERIFICATION PASSED
2018-03-08 02:57:17.797 IST [ledgerfsck] Verify -> INFO 0c0 block number [88]: previous hash matched
2018-03-08 02:57:17.798 IST [ledgerfsck] Verify -> INFO 0c1 Block [seq = 88], hash = [251d36791c7774d716da87774d3631d2a713667de1188c5c4e4318c017e6fbe5], previous hash = [b0fe399e33b9ca507402eef8650e8dbe35c97bb859f29e589a95c48e4299fdcb], VERIFICATION PASSED
2018-03-08 02:57:17.798 IST [ledgerfsck] Verify -> INFO 0c2 block number [89]: previous hash matched
2018-03-08 02:57:17.798 IST [ledgerfsck] Verify -> INFO 0c3 Block [seq = 89], hash = [ddfca7e51b865dccc83d064d0f10d52c5c6d0d22b5bcd3bd8da09f3714229ec7], previous hash = [251d36791c7774d716da87774d3631d2a713667de1188c5c4e4318c017e6fbe5], VERIFICATION PASSED
2018-03-08 02:57:17.799 IST [ledgerfsck] Verify -> INFO 0c4 block number [90]: previous hash matched
2018-03-08 02:57:17.799 IST [ledgerfsck] Verify -> INFO 0c5 Block [seq = 90], hash = [1dc9e987eae2dcf54cec130055e8f589c2836fb60e3ec6a67000ffaca0be9c1c], previous hash = [ddfca7e51b865dccc83d064d0f10d52c5c6d0d22b5bcd3bd8da09f3714229ec7], VERIFICATION PASSED
2018-03-08 02:57:17.799 IST [ledgerfsck] Verify -> INFO 0c6 block number [91]: previous hash matched
2018-03-08 02:57:17.799 IST [ledgerfsck] Verify -> INFO 0c7 Block [seq = 91], hash = [7dac43e4fd96527fb8fe9f6ef6968e9b1b374244ecfa3e6e80cdc65b05cb044f], previous hash = [1dc9e987eae2dcf54cec130055e8f589c2836fb60e3ec6a67000ffaca0be9c1c], VERIFICATION PASSED
2018-03-08 02:57:17.800 IST [ledgerfsck] Verify -> INFO 0c8 block number [92]: previous hash matched
2018-03-08 02:57:17.800 IST [ledgerfsck] Verify -> INFO 0c9 Block [seq = 92], hash = [cfec2b4b101684406298afffb42a4f4dfbcbbb60c035346e407a57cfd0d667d1], previous hash = [7dac43e4fd96527fb8fe9f6ef6968e9b1b374244ecfa3e6e80cdc65b05cb044f], VERIFICATION PASSED
2018-03-08 02:57:17.800 IST [ledgerfsck] Verify -> INFO 0ca block number [93]: previous hash matched
2018-03-08 02:57:17.800 IST [ledgerfsck] Verify -> INFO 0cb Block [seq = 93], hash = [b6f325db476e44a5836a76c484cbfd68dbfa9a648c94038bcc66d8a464994058], previous hash = [cfec2b4b101684406298afffb42a4f4dfbcbbb60c035346e407a57cfd0d667d1], VERIFICATION PASSED
2018-03-08 02:57:17.801 IST [ledgerfsck] Verify -> INFO 0cc block number [94]: previous hash matched
2018-03-08 02:57:17.801 IST [ledgerfsck] Verify -> INFO 0cd Block [seq = 94], hash = [17d3d0590ebd93f18680574fb16189bfd90ee9c52f4ca76233b77256d98084f0], previous hash = [b6f325db476e44a5836a76c484cbfd68dbfa9a648c94038bcc66d8a464994058], VERIFICATION PASSED
2018-03-08 02:57:17.801 IST [ledgerfsck] Verify -> INFO 0ce block number [95]: previous hash matched
2018-03-08 02:57:17.801 IST [ledgerfsck] Verify -> INFO 0cf Block [seq = 95], hash = [224d00f8609520942e69130df811df506fe3af3902c316f971bcf28304a2c234], previous hash = [17d3d0590ebd93f18680574fb16189bfd90ee9c52f4ca76233b77256d98084f0], VERIFICATION PASSED
2018-03-08 02:57:17.802 IST [ledgerfsck] Verify -> INFO 0d0 block number [96]: previous hash matched
2018-03-08 02:57:17.802 IST [ledgerfsck] Verify -> INFO 0d1 Block [seq = 96], hash = [ae4d5cb4559070154e163473e77cc41a0350c96f0c65cdc8594ccb76d67cd06e], previous hash = [224d00f8609520942e69130df811df506fe3af3902c316f971bcf28304a2c234], VERIFICATION PASSED
2018-03-08 02:57:17.802 IST [ledgerfsck] Verify -> INFO 0d2 block number [97]: previous hash matched
2018-03-08 02:57:17.802 IST [ledgerfsck] Verify -> INFO 0d3 Block [seq = 97], hash = [ae0f872b350d7f01503635207dca36c5b6847e9132504867063e95a33254ac33], previous hash = [ae4d5cb4559070154e163473e77cc41a0350c96f0c65cdc8594ccb76d67cd06e], VERIFICATION PASSED
2018-03-08 02:57:17.803 IST [ledgerfsck] Verify -> INFO 0d4 block number [98]: previous hash matched
2018-03-08 02:57:17.803 IST [ledgerfsck] Verify -> INFO 0d5 Block [seq = 98], hash = [55e36ad0daa3b829258eafce24fded2aa2b1b006237088cc2ee9a8b387c9b773], previous hash = [ae0f872b350d7f01503635207dca36c5b6847e9132504867063e95a33254ac33], VERIFICATION PASSED
2018-03-08 02:57:17.803 IST [ledgerfsck] Verify -> INFO 0d6 block number [99]: previous hash matched
2018-03-08 02:57:17.803 IST [ledgerfsck] Verify -> INFO 0d7 Block [seq = 99], hash = [ac9da4411c2d15c340dd87519369718900903f63281dcdc372f6cc9abb16b2fa], previous hash = [55e36ad0daa3b829258eafce24fded2aa2b1b006237088cc2ee9a8b387c9b773], VERIFICATION PASSED
2018-03-08 02:57:17.803 IST [ledgerfsck] Verify -> INFO 0d8 block number [100]: previous hash matched
2018-03-08 02:57:17.804 IST [ledgerfsck] Verify -> INFO 0d9 Block [seq = 100], hash = [f3061eddf1c7c866b02bb712dd05c4a44e7921c93dfeeb109d291a52036daf21], previous hash = [ac9da4411c2d15c340dd87519369718900903f63281dcdc372f6cc9abb16b2fa], VERIFICATION PASSED
2018-03-08 02:57:17.804 IST [ledgerfsck] Verify -> INFO 0da block number [101]: previous hash matched
2018-03-08 02:57:17.804 IST [ledgerfsck] Verify -> INFO 0db Block [seq = 101], hash = [fcf96572a393440ad6495fea869457304da033baac93ad58677f2d8ef4a81d23], previous hash = [f3061eddf1c7c866b02bb712dd05c4a44e7921c93dfeeb109d291a52036daf21], VERIFICATION PASSED
2018-03-08 02:57:17.804 IST [ledgerfsck] Verify -> INFO 0dc block number [102]: previous hash matched
2018-03-08 02:57:17.805 IST [ledgerfsck] Verify -> INFO 0dd Block [seq = 102], hash = [74dc3b599a463656e38a1782acbc71473a3d710c2ff250afcb5da6e1e0fb00bf], previous hash = [fcf96572a393440ad6495fea869457304da033baac93ad58677f2d8ef4a81d23], VERIFICATION PASSED
2018-03-08 02:57:17.805 IST [ledgerfsck] Verify -> INFO 0de block number [103]: previous hash matched
2018-03-08 02:57:17.805 IST [ledgerfsck] Verify -> INFO 0df Block [seq = 103], hash = [47ccfe029b202f95e1955c5a281d6567300083afd77e9d05c68826ab3e5ebcf1], previous hash = [74dc3b599a463656e38a1782acbc71473a3d710c2ff250afcb5da6e1e0fb00bf], VERIFICATION PASSED
2018-03-08 02:57:17.805 IST [ledgerfsck] Verify -> INFO 0e0 block number [104]: previous hash matched
2018-03-08 02:57:17.806 IST [ledgerfsck] Verify -> INFO 0e1 Block [seq = 104], hash = [25944e93cc72da73137f6f8958e85f0ed0740ca559de0c2d6856d938de947d4d], previous hash = [47ccfe029b202f95e1955c5a281d6567300083afd77e9d05c68826ab3e5ebcf1], VERIFICATION PASSED
2018-03-08 02:57:17.806 IST [ledgerfsck] Verify -> INFO 0e2 block number [105]: previous hash matched
2018-03-08 02:57:17.806 IST [ledgerfsck] Verify -> INFO 0e3 Block [seq = 105], hash = [c5a743d46066912e50729dc204ab2fde558209ce950273a4612548e01ebea0b4], previous hash = [25944e93cc72da73137f6f8958e85f0ed0740ca559de0c2d6856d938de947d4d], VERIFICATION PASSED
2018-03-08 02:57:17.806 IST [ledgerfsck] Verify -> INFO 0e4 block number [106]: previous hash matched
2018-03-08 02:57:17.807 IST [ledgerfsck] Verify -> INFO 0e5 Block [seq = 106], hash = [6fce894a179ee645cace06a7b430d18f7f1eabeb154423d76bde994cb1711e20], previous hash = [c5a743d46066912e50729dc204ab2fde558209ce950273a4612548e01ebea0b4], VERIFICATION PASSED
2018-03-08 02:57:17.807 IST [ledgerfsck] Verify -> INFO 0e6 block number [107]: previous hash matched
2018-03-08 02:57:17.807 IST [ledgerfsck] Verify -> INFO 0e7 Block [seq = 107], hash = [694695dbd9d9c4f889ada4e05939003cc9ae6b94b4870002691bbd62f9d6230e], previous hash = [6fce894a179ee645cace06a7b430d18f7f1eabeb154423d76bde994cb1711e20], VERIFICATION PASSED
2018-03-08 02:57:17.807 IST [ledgerfsck] Verify -> INFO 0e8 block number [108]: previous hash matched
2018-03-08 02:57:17.808 IST [ledgerfsck] Verify -> INFO 0e9 Block [seq = 108], hash = [1389635d275313cdb5928326a9671667dbeb4c825de74dd87cfeeaf1407dcdb3], previous hash = [694695dbd9d9c4f889ada4e05939003cc9ae6b94b4870002691bbd62f9d6230e], VERIFICATION PASSED
2018-03-08 02:57:17.808 IST [ledgerfsck] Verify -> INFO 0ea block number [109]: previous hash matched
2018-03-08 02:57:17.808 IST [ledgerfsck] Verify -> INFO 0eb Block [seq = 109], hash = [4ea907eb885439b7aaf07ad397d00f34eb89a3a4b8b15cb227a5cdfba43eece4], previous hash = [1389635d275313cdb5928326a9671667dbeb4c825de74dd87cfeeaf1407dcdb3], VERIFICATION PASSED
2018-03-08 02:57:17.808 IST [ledgerfsck] Verify -> INFO 0ec block number [110]: previous hash matched
2018-03-08 02:57:17.809 IST [ledgerfsck] Verify -> INFO 0ed Block [seq = 110], hash = [095199245fad8ce6cea11f3c0cfdd909af9f5d6651af44578b7aea045293d49d], previous hash = [4ea907eb885439b7aaf07ad397d00f34eb89a3a4b8b15cb227a5cdfba43eece4], VERIFICATION PASSED
2018-03-08 02:57:17.809 IST [ledgerfsck] Verify -> INFO 0ee block number [111]: previous hash matched
2018-03-08 02:57:17.809 IST [ledgerfsck] Verify -> INFO 0ef Block [seq = 111], hash = [3e2111ccf555210d8c7bfdd90985cd157ac1af411ddcf2c2533e53bd03aafdab], previous hash = [095199245fad8ce6cea11f3c0cfdd909af9f5d6651af44578b7aea045293d49d], VERIFICATION PASSED
2018-03-08 02:57:17.809 IST [ledgerfsck] Verify -> INFO 0f0 block number [112]: previous hash matched
2018-03-08 02:57:17.810 IST [ledgerfsck] Verify -> INFO 0f1 Block [seq = 112], hash = [890e8fa34aa604c3d65f5e2fdc27bd4f34bed9c2c7f54abf1fdb5d7b5f809bee], previous hash = [3e2111ccf555210d8c7bfdd90985cd157ac1af411ddcf2c2533e53bd03aafdab], VERIFICATION PASSED
2018-03-08 02:57:17.810 IST [ledgerfsck] Verify -> INFO 0f2 block number [113]: previous hash matched
2018-03-08 02:57:17.810 IST [ledgerfsck] Verify -> INFO 0f3 Block [seq = 113], hash = [020dc88b8311050fa0580d745ffa94a974ad71eb3723b56caaa8b7528c59cd45], previous hash = [890e8fa34aa604c3d65f5e2fdc27bd4f34bed9c2c7f54abf1fdb5d7b5f809bee], VERIFICATION PASSED
2018-03-08 02:57:17.810 IST [ledgerfsck] Verify -> INFO 0f4 block number [114]: previous hash matched
2018-03-08 02:57:17.811 IST [ledgerfsck] Verify -> INFO 0f5 Block [seq = 114], hash = [1b64d820db528ac48e2331150d2db0c65bf2b2a283c851dcd48d8bef764b20b8], previous hash = [020dc88b8311050fa0580d745ffa94a974ad71eb3723b56caaa8b7528c59cd45], VERIFICATION PASSED
2018-03-08 02:57:17.811 IST [ledgerfsck] Verify -> INFO 0f6 block number [115]: previous hash matched
2018-03-08 02:57:17.811 IST [ledgerfsck] Verify -> INFO 0f7 Block [seq = 115], hash = [36d57ac3839e0a520362c6b986cde8171f746a725a0bf13037c6a7f5a54a4ba7], previous hash = [1b64d820db528ac48e2331150d2db0c65bf2b2a283c851dcd48d8bef764b20b8], VERIFICATION PASSED
2018-03-08 02:57:17.811 IST [ledgerfsck] Verify -> INFO 0f8 block number [116]: previous hash matched
2018-03-08 02:57:17.812 IST [ledgerfsck] Verify -> INFO 0f9 Block [seq = 116], hash = [d3420103600e8ced3c54e6c26f947b60b7ffcfb5f6f65897638ff7f962b484ac], previous hash = [36d57ac3839e0a520362c6b986cde8171f746a725a0bf13037c6a7f5a54a4ba7], VERIFICATION PASSED
2018-03-08 02:57:17.812 IST [ledgerfsck] Verify -> INFO 0fa block number [117]: previous hash matched
2018-03-08 02:57:17.812 IST [ledgerfsck] Verify -> INFO 0fb Block [seq = 117], hash = [c6fa496ab34af7481a4e0dacfceb7e5c1dce517880c81d3de3b4b59700e0d5e8], previous hash = [d3420103600e8ced3c54e6c26f947b60b7ffcfb5f6f65897638ff7f962b484ac], VERIFICATION PASSED
2018-03-08 02:57:17.812 IST [ledgerfsck] Verify -> INFO 0fc block number [118]: previous hash matched
2018-03-08 02:57:17.813 IST [ledgerfsck] Verify -> INFO 0fd Block [seq = 118], hash = [0f4eb4ac5460570655881adf59ddf9984532ce4ae3455edd53a64893115910c8], previous hash = [c6fa496ab34af7481a4e0dacfceb7e5c1dce517880c81d3de3b4b59700e0d5e8], VERIFICATION PASSED
2018-03-08 02:57:17.813 IST [ledgerfsck] Verify -> INFO 0fe block number [119]: previous hash matched
2018-03-08 02:57:17.813 IST [ledgerfsck] Verify -> INFO 0ff Block [seq = 119], hash = [33744d41f05eeda20ec96b5e4f252fcc50d72dfe87a6803e96266292e4f6e661], previous hash = [0f4eb4ac5460570655881adf59ddf9984532ce4ae3455edd53a64893115910c8], VERIFICATION PASSED


```