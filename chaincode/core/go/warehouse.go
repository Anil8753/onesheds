package main

import (
	"encoding/json"
	"strconv"

	assets "github.com/anil8753/onesheds/chaincode/core/warehouse/asset"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func (c *Contract) RegisterWarehouse(
	ctx contractapi.TransactionContextInterface,
	input string,
) (string, error) {

	r, err := assets.Register(ctx, input)
	if err != nil {
		return "", err
	}

	b, err := json.Marshal(r)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func (c *Contract) GetAllWarehouse(
	ctx contractapi.TransactionContextInterface,
) (string, error) {

	r, err := assets.QueryAll(ctx)

	if err != nil {
		return "", err
	}

	b, err := json.Marshal(r)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func (c *Contract) GetWarehouse(
	ctx contractapi.TransactionContextInterface,
	uniqueId string,
) (string, error) {

	r, err := assets.Query(ctx, uniqueId)
	if err != nil {
		return "", err
	}

	b, err := json.Marshal(r)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func (c *Contract) GetWarehouseByOwnerId(
	ctx contractapi.TransactionContextInterface,
	ownerId string,
) (string, error) {

	r, err := assets.QueryByOwnerId(ctx, ownerId)
	if err != nil {
		return "", err
	}

	b, err := json.Marshal(r)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func (c *Contract) GetWarehousesByRichQuery(
	ctx contractapi.TransactionContextInterface,
	query string,
) (string, error) {

	r, err := assets.RichQuery(ctx, query)
	if err != nil {
		return "", err
	}

	b, err := json.Marshal(r)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func (c *Contract) GetWarehousesByRichQueryWithPagination(
	ctx contractapi.TransactionContextInterface,
	query string,
	pageSize string,
	bookmark string,
) (string, error) {

	ps, err := strconv.Atoi(pageSize)
	if err != nil {
		return "", err
	}

	r, err := assets.RichQueryWithPagination(ctx, query, ps, bookmark)
	if err != nil {
		return "", err
	}

	b, err := json.Marshal(r)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func (c *Contract) UpdateWarehouse(
	ctx contractapi.TransactionContextInterface,
	input string,
) (string, error) {

	r, err := assets.Update(ctx, input)
	if err != nil {
		return "", err
	}

	b, err := json.Marshal(r)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
