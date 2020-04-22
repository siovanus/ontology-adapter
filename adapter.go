package main

import (
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
	Value     interface{}
	Result    interface{}
	RequestId interface{} `json:"request_id"`
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
	// Set Value to whatever is defined in the "result" key
	// if the default "value" is empty
	if req.Value == nil || req.Value == "" {
		req.Value = req.Result
	}

	args := []interface{}{FulfillOracleRequest, []interface{}{store.OntTxManager.Account().Address[:],
		requestID, payment, callbackAddress, callbackFunction, expiration, data}}
	hash, err := adapter.sdk.NeoVM.InvokeNeoVMContract(DefaultOntGasPrice, DefaultOntGasLimit, adapter.account,
		adapter.account, adapter.address, args)
	if err != nil {
		return nil, err
	}

	return hash, nil
}
