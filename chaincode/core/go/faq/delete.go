package faq

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func Delete(
	ctx contractapi.TransactionContextInterface,
	warehouseId string,
	index int,
) (*Entry, error) {

	id, err := ctx.GetStub().CreateCompositeKey(IDPrefix, []string{warehouseId})
	if err != nil {
		return nil, fmt.Errorf("failed to create composit key. %w", err)
	}

	entry, err := get(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get FAQ for the warehouse. %s. Error: %w", id, err)
	}

	if len(entry.FAQs) <= index {
		return nil, fmt.Errorf("invalid index: %d", index)
	}

	if err := CanAdd(ctx, entry); err != nil {
		return nil, err
	}

	entry.FAQs = append(entry.FAQs[:index], entry.FAQs[index+1:]...)

	outBytes, err := json.Marshal(entry)
	if err != nil {
		return nil, err
	}

	if err := ctx.GetStub().PutState(id, outBytes); err != nil {
		return nil, fmt.Errorf("failed to put data. %w", err)
	}

	return entry, nil
}
