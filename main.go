package main

import (
	"fmt"
	"gorestledger/blockchain"
	"gorestledger/web"
	"gorestledger/web/controllers"
	"os"
)

func main() {
	// Definition of the Fabric SDK properties
	fSetup := blockchain.FabricSetup{
		// Network parameters
		OrdererID: "orderer.go.rest.ledger.com",

		// Channel parameters
		ChannelID:     "gorestledger",
		ChannelConfig: os.Getenv("GOPATH") + "/src/gorestledger/fixtures/artifacts/gorestledger.channel.tx",

		// Chaincode parameters
		ChaincodeID:     "gorestledger",
		ChaincodeGoPath: os.Getenv("GOPATH"),
		ChaincodePath:   "gorestledger/chaincode/",
		OrgAdmin:        "Admin",
		OrgName:         "org1",
		ConfigFile:      "config.yaml",

		// CA parameters
		CaID: "ca.org1.go.rest.ledger.com",
	}

	// Initialization of the Fabric SDK from the previously set properties
	err := fSetup.Initialize()
	if err != nil {
		fmt.Printf("Unable to initialize the Fabric SDK: %v\n", err)
		return
	}
	// Close SDK
	defer fSetup.CloseSDK()

	// Install and instantiate the chaincode
	err = fSetup.InstallAndInstantiateCC()
	if err != nil {
		fmt.Printf("Unable to install and instantiate the chaincode: %v\n", err)
		return
	}

	app := &controllers.Application{
		Fabric: &fSetup,
	}
	web.Serve(app)
}
