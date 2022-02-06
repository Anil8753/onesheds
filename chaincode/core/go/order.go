package main

import (
	"encoding/json"

	txn "github.com/anil8753/onesheds/chaincode/core/txn"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func (c *Contract) NewOrder(
	ctx contractapi.TransactionContextInterface,
	input string,
) (string, error) {

	r, err := txn.NewOrder(ctx, input)

	if err != nil {
		return "", err
	}

	b, err := json.Marshal(r)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func (c *Contract) AcceptOrder(
	ctx contractapi.TransactionContextInterface,
	orderId string,
) error {

	err := txn.AcceptOrder(ctx, orderId)

	if err != nil {
		return err
	}

	return nil
}

func (c *Contract) RejectOrder(
	ctx contractapi.TransactionContextInterface,
	orderId string,
) error {

	err := txn.RejectOrder(ctx, orderId)

	if err != nil {
		return err
	}

	return nil
}

func (c *Contract) QueryDepositorAllOrders(
	ctx contractapi.TransactionContextInterface,
	depositorId string,
) (string, error) {

	r, err := txn.QueryDepositorAllOrders(ctx, depositorId)

	if err != nil {
		return "", err
	}

	b, err := json.Marshal(r)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func (c *Contract) QueryOrder(
	ctx contractapi.TransactionContextInterface,
	orderId string,
) (string, error) {

	r, err := txn.QueryOrder(ctx, orderId)

	if err != nil {
		return "", err
	}

	b, err := json.Marshal(r)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func (c *Contract) QueryWarehouseAllOrders(
	ctx contractapi.TransactionContextInterface,
	warehouseId string,
) (string, error) {

	r, err := txn.QueryWarehouseAllOrders(ctx, warehouseId)

	if err != nil {
		return "", err
	}

	b, err := json.Marshal(r)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func (c *Contract) QueryWarehouseAllOrdersWithPagination(
	ctx contractapi.TransactionContextInterface,
	warehouseId string,
	pageSize int,
	bookmark string,
) (string, error) {

	r, err := txn.QueryWarehouseAllOrdersWithPagination(ctx, warehouseId, pageSize, bookmark)

	if err != nil {
		return "", err
	}

	b, err := json.Marshal(r)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
