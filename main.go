package main

import (
	"fmt"
	"github.com/smartcontractkit/ontology-adapter/config"
)

func main() {
	fmt.Println("Starting Ontology adapter")
	err := config.DefConfig.Init(config.DEFAULT_CONFIG_FILE_NAME)
	if err != nil {
		fmt.Println("DefConfig.Init error:", err)
		return
	}

	adapter, err := newOntologyAdapter(config.DefConfig.OntologyWalletPath,
		config.DefConfig.OracleContractAddress, config.DefConfig.OntologyRpc)
	if err != nil {
		fmt.Println("Failed starting Ontology adapter:", err)
		return
	}

	RunWebserver(adapter.handle, config.DefConfig.Listening)
}
