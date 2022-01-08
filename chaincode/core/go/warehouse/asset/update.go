package asset

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func Update(ctx contractapi.TransactionContextInterface, input string) (*AssetData, error) {

	// Todo: check the ownership from the user x509 certificate

	pData, err := NewAssetData(input)
	if err != nil {
		return nil, fmt.Errorf("invalid registeration data. %w", err)
	}

	sData, err := Query(ctx, pData.WarehouseId)
	if err != nil {
		return nil, fmt.Errorf("%s warehouse not found in state db. %w", pData.WarehouseId, err)
	}

	if pData.WarehouseId != sData.WarehouseId {
		return nil, fmt.Errorf("passed %s WarehouseId not same as state WarehouseId %s", pData.WarehouseId, sData.WarehouseId)
	}

	updated, err := updateRegData(pData, sData)
	if err != nil {
		return nil, err
	}

	b, err := json.Marshal(updated)
	if err != nil {
		return nil, err
	}

	if err := ctx.GetStub().PutState(updated.WarehouseId, b); err != nil {
		return nil, fmt.Errorf("failed to register warehouse. %w", err)
	}

	return updated, nil
}

// updateRegData updated the field from 'from' to 'to' AssetData
func updateRegData(from *AssetData, to *AssetData) (*AssetData, error) {

	if from.Status != "" {
		to.Status = from.Status
	}

	if from.OwnerId != "" {
		to.OwnerId = from.OwnerId
	}

	if to.Properties == nil {
		to.Properties = make(map[string]interface{})
	}

	if to.Photos == nil {
		to.Photos = make(map[string]interface{})
	}

	if to.Videos == nil {
		to.Videos = make(map[string]interface{})
	}

	for k, v := range from.Properties {
		to.Properties[k] = v
	}

	for k, v := range from.Photos {
		to.Photos[k] = v
	}

	for k, v := range from.Videos {
		to.Videos[k] = v
	}

	return to, nil
}
