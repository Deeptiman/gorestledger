export FABRIC_CFG_PATH=$PWD

./bin/cryptogen generate --config=./crypto-config.yaml

./bin/configtxgen -profile GoRestLedger -outputBlock ./artifacts/orderer.genesis.block

./bin/configtxgen -profile GoRestLedger -outputCreateChannelTx ./artifacts/gorestledger.channel.tx -channelID gorestledger

./bin/configtxgen -profile GoRestLedger -outputAnchorPeersUpdate ./artifacts/org1.gorestledger.anchors.tx -channelID gorestledger -asOrg GoRestLedgerOrganization1
