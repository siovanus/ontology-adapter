# Ontology External Adapter

## Running

- Install Go dependencies `go install`

- Build executable `go build -o ontology-adapter`

- Run `./ontology-adapter` and enter your password

## Configuration

### config.json

```json
{
  "OracleContractAddress": "b54dd842fadc8b04f0c58b1ea921f49bf54d04f0",
  "OntologyWalletPath": "wallet.dat",
  "OntologyRpc": "http://polaris1.ont.io:20336",
  "Listening": "0.0.0.0:8090"
}
```

`OracleContractAddress`: your oracle contract address in ontology network

`OntologyWalletPath`: path of your ontology wallet file

`OntologyRpc`: ontology node and port you want to connect, for testnet use http://polaris1.ont.io:20336, for mainnet use http://dappnode1.ont.io:20336, you can use your own node as well.

`Listening`: port this ontology external adapter listening to.

