package contract

import (
	"encoding/json"

	"github.com/anil8753/onesheds/chaincode/core/faq"
	"github.com/anil8753/onesheds/chaincode/core/utils"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func (c *Contract) GetAllFAQ(
	ctx contractapi.TransactionContextInterface,
	warehouseId string,
) (string, error) {
	r, err := faq.GetByWarehouseId(ctx, warehouseId)
	if err != nil {
		return "", err
	}

	return utils.ToJSON(r)
}

type AddFAQData struct {
	WarehouseId string `json:"warehouseId"`
	Question    string `json:"question"`
	Input       string `json:"input"`
}

func (c *Contract) AddFAQ(
	ctx contractapi.TransactionContextInterface,
	input string,
) (string, error) {

	var in AddFAQData
	if err := json.Unmarshal([]byte(input), &in); err != nil {
		return "", err
	}

	r, err := faq.Add(ctx, in.WarehouseId, in.Question, in.Input)
	if err != nil {
		return "", err
	}

	return utils.ToJSON(r)
}

type UpdateQuestionData struct {
	WarehouseId string `json:"warehouseId"`
	Index       int    `json:"index"`
	Question    string `json:"question"`
}

func (c *Contract) UpdateFAQQuestion(
	ctx contractapi.TransactionContextInterface,
	input string,
) (string, error) {

	var in UpdateQuestionData
	if err := json.Unmarshal([]byte(input), &in); err != nil {
		return "", err
	}

	r, err := faq.UpdateQuestion(ctx, in.WarehouseId, in.Index, in.Question)
	if err != nil {
		return "", err
	}

	return utils.ToJSON(r)
}

type UpdateAnswerData struct {
	WarehouseId string `json:"warehouseId"`
	Index       int    `json:"index"`
	Input       string `json:"input"`
}

func (c *Contract) UpdateFAQAnswer(
	ctx contractapi.TransactionContextInterface,
	input string,
) (string, error) {

	var in UpdateAnswerData
	if err := json.Unmarshal([]byte(input), &in); err != nil {
		return "", err
	}

	r, err := faq.UpdateAnswer(ctx, in.WarehouseId, in.Index, in.Input)
	if err != nil {
		return "", err
	}

	return utils.ToJSON(r)
}

type DeleteData struct {
	WarehouseId string `json:"warehouseId"`
	Index       int    `json:"index"`
}

func (c *Contract) DeleteFAQ(
	ctx contractapi.TransactionContextInterface,
	input string,
) (string, error) {

	var in DeleteData
	if err := json.Unmarshal([]byte(input), &in); err != nil {
		return "", err
	}

	r, err := faq.Delete(ctx, in.WarehouseId, in.Index)
	if err != nil {
		return "", err
	}

	return utils.ToJSON(r)
}
