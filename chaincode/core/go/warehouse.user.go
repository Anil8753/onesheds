package main

import (
	whUser "github.com/anil8753/onesheds/chaincode/core/warehouse/user"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func (c *Contract) RegisterWarehouseUser(
	ctx contractapi.TransactionContextInterface,
	input string,
) error {

	return whUser.Register(ctx, input)
}

func (c *Contract) UpdateWarehouseUser(
	ctx contractapi.TransactionContextInterface,
	input string,
) error {

	return whUser.Update(ctx, input)
}

func (c *Contract) GetWarehouseUser(
	ctx contractapi.TransactionContextInterface,
	uniqueId string,
) (*whUser.RegisterationData, error) {

	return whUser.Query(ctx, uniqueId)
}
