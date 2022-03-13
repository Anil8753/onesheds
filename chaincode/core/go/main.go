package main

import (
	"log"

	"github.com/anil8753/onesheds/chaincode/core/contract"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func main() {
	chaincode, err := contractapi.NewChaincode(&contract.Contract{})
	if err != nil {
		log.Panicf("Error creating chaincode: %v", err)
	}

	if err := chaincode.Start(); err != nil {
		log.Panicf("Error starting chaincode: %v", err)
	}
}
