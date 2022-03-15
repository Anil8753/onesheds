package faq

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func UpdateQuestion(
	ctx contractapi.TransactionContextInterface,
	warehouseId string,
	index int,
	question string,
) (*Entry, error) {

	id, err := ctx.GetStub().CreateCompositeKey(IDPrefix, []string{warehouseId})
	if err != nil {
		return nil, fmt.Errorf("failed to create composit key. %w", err)
	}

	entry, err := get(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get FAQ for the warehouse. %s. Error: %w", id, err)
	}

	if len(entry.FAQs) >= index {
		return nil, fmt.Errorf("invalid index: %d", index)
	}

	if err := CanAdd(ctx, entry); err != nil {
		return nil, err
	}

	old := entry.FAQs[index]
	entry.FAQs[index] = FAQ{Question: question, Answer: old.Answer}

	outBytes, err := json.Marshal(entry)
	if err != nil {
		return nil, err
	}

	if err := ctx.GetStub().PutState(id, outBytes); err != nil {
		return nil, fmt.Errorf("failed to put data. %w", err)
	}

	return entry, nil
}

func UpdateAnswer(
	ctx contractapi.TransactionContextInterface,
	warehouseId string,
	index int,
	answerString string,
) (*Entry, error) {

	id, err := ctx.GetStub().CreateCompositeKey(IDPrefix, []string{warehouseId})
	if err != nil {
		return nil, fmt.Errorf("failed to create composit key. %w", err)
	}

	entry, err := get(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get FAQ for the warehouse. %s. Error: %w", id, err)
	}

	if len(entry.FAQs) >= index {
		return nil, fmt.Errorf("invalid index: %d", index)
	}

	if err := CanAdd(ctx, entry); err != nil {
		return nil, err
	}

	var answer Answer
	if err := json.Unmarshal([]byte(answerString), &answer); err != nil {
		return nil, fmt.Errorf("invalid input data. %w", err)
	}

	old := entry.FAQs[index]
	entry.FAQs[index] = FAQ{Question: old.Question, Answer: answer}

	outBytes, err := json.Marshal(entry)
	if err != nil {
		return nil, err
	}

	if err := ctx.GetStub().PutState(id, outBytes); err != nil {
		return nil, fmt.Errorf("failed to put data. %w", err)
	}

	return entry, nil
}
