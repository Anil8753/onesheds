package knowledgebase

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func GetByWarehouseId(
	ctx contractapi.TransactionContextInterface,
	warehouseId string,
) (*Document, error) {

	key, err := ctx.GetStub().CreateCompositeKey(IDPrefix, []string{warehouseId})
	if err != nil {
		return nil, fmt.Errorf("failed to create composit key. %w", err)
	}

	return get(ctx, key)
}

func get(
	ctx contractapi.TransactionContextInterface,
	key string,
) (*Document, error) {

	outBytes, err := ctx.GetStub().GetState(key)
	if err != nil {
		return nil, fmt.Errorf("failed to get data for key: %s. Error: %w", key, err)
	}

	if len(outBytes) == 0 {
		return nil, fmt.Errorf("key not found: %s", key)
	}

	var doc Document
	if err := json.Unmarshal(outBytes, &doc); err != nil {
		return nil, err
	}

	return &doc, nil
}
