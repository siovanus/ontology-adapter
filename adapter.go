package main

import (
	"encoding/hex"
	sdk "github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology/common"
	"github.com/ontio/ontology/common/password"
)

const (
	DefaultOntGasPrice = 500
	DefaultOntGasLimit = 20000

	FulfillOracleRequest = "fulfillOracleRequest"
)

type Request struct {
	Address          string      `json:"address"`
	RequestID        string      `json:"requestID"`
	Payment          string      `json:"payment"`
	CallbackAddress  string      `json:"callbackAddress"`
	CallbackFunction string      `json:"callbackFunction"`
	Expiration       string      `json:"expiration"`
	Result           interface{} `json:"result"`
}

type ontologyAdapter struct {
	address common.Address
	account *sdk.Account
	sdk     *sdk.OntologySdk
}

func newOntologyAdapter(path, address, endpoint string) (*ontologyAdapter, error) {
	addr, err := common.AddressFromHexString(address)
	if err != nil {
		return nil, err
	}

	wallet, err := sdk.OpenWallet(path)
	if err != nil {
		return nil, err
	}
	pwd, err := password.GetPassword()
	if err != nil {
		return nil, err
	}
	account, err := wallet.GetDefaultAccount(pwd)
	if err != nil {
		return nil, err
	}

	sdk := sdk.NewOntologySdk()
	sdk.NewRpcClient().SetAddress(endpoint)

	return &ontologyAdapter{
		address: addr,
		account: account,
		sdk:     sdk,
	}, nil
}

func (adapter ontologyAdapter) handle(req Request) (interface{}, error) {
	requestID, err := hex.DecodeString(req.RequestID)
	if err != nil {
		return nil, err
	}
	payment, err := hex.DecodeString(req.Payment)
	if err != nil {
		return nil, err
	}
	callbackAddress, err := hex.DecodeString(req.CallbackAddress)
	if err != nil {
		return nil, err
	}
	callbackFunction, err := hex.DecodeString(req.CallbackFunction)
	if err != nil {
		return nil, err
	}
	expiration, err := hex.DecodeString(req.Expiration)
	if err != nil {
		return nil, err
	}

	args := []interface{}{FulfillOracleRequest, []interface{}{adapter.account.Address[:],
		requestID, payment, callbackAddress, callbackFunction, expiration, req.Result}}
	hash, err := adapter.sdk.NeoVM.InvokeNeoVMContract(DefaultOntGasPrice, DefaultOntGasLimit, adapter.account,
		adapter.account, adapter.address, args)
	if err != nil {
		return nil, err
	}

	return hash, nil
}
