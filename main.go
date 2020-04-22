package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Starting Substrate adapter")

	contractAddress := os.Getenv("ORACLE_CONTRACT_ADDRESS")
	walletPath := os.Getenv("ONTOLOGY_WALLET_PATH")
	endpoint := os.Getenv("ONTOLOGY_RPC")

	adapter, err := newOntologyAdapter(walletPath, contractAddress, endpoint)
	if err != nil {
		fmt.Println("Failed starting Substrate adapter:", err)
		return
	}

	RunWebserver(adapter.handle)
}
