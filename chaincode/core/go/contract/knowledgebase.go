package contract

import (
	"encoding/json"

	"github.com/anil8753/onesheds/chaincode/core/knowledgebase"
	"github.com/anil8753/onesheds/chaincode/core/utils"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func (c *Contract) GetAllQuesAndAnswer(
	ctx contractapi.TransactionContextInterface,
	warehouseId string,
) (string, error) {
	r, err := knowledgebase.GetByWarehouseId(ctx, warehouseId)
	if err != nil {
		return "", err
	}

	return utils.ToJSON(r)
}

func (c *Contract) AddWarehouseQuestion(
	ctx contractapi.TransactionContextInterface,
	warehouseId string,
	questionStr string,
) (string, error) {

	r, err := knowledgebase.AddQuestion(ctx, warehouseId, questionStr)
	if err != nil {
		return "", err
	}

	return utils.ToJSON(r)
}

type AddWarehouseAnswerData struct {
	WarehouseId string `json:"warehouseId"`
	Index       int    `json:"index"`
	Input       string `json:"input"`
}

func (c *Contract) AddWarehouseAnswer(
	ctx contractapi.TransactionContextInterface,
	input string,
) (string, error) {

	var in AddWarehouseAnswerData
	if err := json.Unmarshal([]byte(input), &in); err != nil {
		return "", err
	}

	r, err := knowledgebase.AddAnswer(ctx, in.WarehouseId, in.Index, in.Input)
	if err != nil {
		return "", err
	}

	return utils.ToJSON(r)
}
