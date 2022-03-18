package faq

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func GetByWarehouseId(
	ctx contractapi.TransactionContextInterface,
	warehouseId string,
) (*Entry, error) {

	key, err := ctx.GetStub().CreateCompositeKey(IDPrefix, []string{warehouseId})
	if err != nil {
		return nil, fmt.Errorf("failed to create composit key. %w", err)
	}

	return get(ctx, key)
}

func get(
	ctx contractapi.TransactionContextInterface,
	key string,
) (*Entry, error) {

	outBytes, err := ctx.GetStub().GetState(key)
	if err != nil {
		return nil, fmt.Errorf("failed to get data for key: %s. Error: %w", key, err)
	}

	if len(outBytes) == 0 {
		return nil, fmt.Errorf("key not found: %s", key)
	}

	var data Entry
	if err := json.Unmarshal(outBytes, &data); err != nil {
		return nil, err
	}

	return &data, nil
}
