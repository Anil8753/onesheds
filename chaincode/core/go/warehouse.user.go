package main

import (
	"encoding/json"

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
) (string, error) {

	r, err := whUser.Query(ctx, uniqueId)
	if err != nil {
		return "", err
	}

	b, err := json.Marshal(r)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
