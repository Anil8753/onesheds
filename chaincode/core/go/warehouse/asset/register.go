package asset

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func Register(ctx contractapi.TransactionContextInterface, input string) (*AssetData, error) {

	data, err := NewAssetData(input)
	if err != nil {
		return nil, fmt.Errorf("invalid registeration data. %w", err)
	}

	if data.OwnerId == "" {
		return nil, errors.New("OwnerId is mandatory")
	}

	// Todo: check the profile ownership from the user x509 certificate
	if !strings.HasPrefix(data.OwnerId, "user") {
		return nil, errors.New("OwnerId must have prefix user")
	}

	if _, err := Query(ctx, data.WarehouseId); err == nil {
		return nil, fmt.Errorf("WarehouseId (%s) already exist. %w", data.WarehouseId, err)
	}

	// Set the status unpublished
	data.Status = StatusUnpublished

	databytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	if err := ctx.GetStub().PutState(data.WarehouseId, databytes); err != nil {
		return nil, fmt.Errorf("failed to register warehouse. %w", err)
	}

	return data, nil
}
