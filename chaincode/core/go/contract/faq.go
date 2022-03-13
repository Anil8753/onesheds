package contract

import (
	"github.com/anil8753/onesheds/chaincode/core/faq"
	"github.com/anil8753/onesheds/chaincode/core/utils"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func GetByWarehouseId(
	ctx contractapi.TransactionContextInterface,
	warehouseId string,
) (string, error) {
	r, err := faq.GetByWarehouseId(ctx, warehouseId)
	if err != nil {
		return "", err
	}

	return utils.ToJSON(r)
}

func Add(
	ctx contractapi.TransactionContextInterface,
	warehouseId string,
	question string,
	input string,
) (string, error) {

	r, err := faq.Add(ctx, warehouseId, question, input)
	if err != nil {
		return "", err
	}

	return utils.ToJSON(r)
}

func Update(
	ctx contractapi.TransactionContextInterface,
	warehouseId string,
	index int,
	question string,
	input string,
) (string, error) {

	r, err := faq.Update(ctx, warehouseId, index, question, input)
	if err != nil {
		return "", err
	}

	return utils.ToJSON(r)
}
