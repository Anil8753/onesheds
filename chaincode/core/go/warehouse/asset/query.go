package asset

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func Query(ctx contractapi.TransactionContextInterface, warehouseId string) (*AssetData, error) {

	sBytes, err := ctx.GetStub().GetState(warehouseId)
	if err != nil {
		return nil, err
	}

	if len(sBytes) == 0 {
		return nil, fmt.Errorf("'%s' warehouse is not found", warehouseId)
	}

	var data AssetData
	if err := json.Unmarshal(sBytes, &data); err != nil {
		return nil, err
	}

	return &data, nil
}

func QueryAll(ctx contractapi.TransactionContextInterface) ([]*AssetData, error) {

	qFmt := `{ "selector": { "docType": "%s" } }`
	q := fmt.Sprintf(qFmt, WarehouseDocType)
	return query(ctx, q)
}

func QueryByOwnerId(ctx contractapi.TransactionContextInterface, ownerId string) ([]*AssetData, error) {

	qFmt := `{ "selector": { "ownerId": "%s" } }`
	q := fmt.Sprintf(qFmt, ownerId)

	return query(ctx, q)
}

func query(ctx contractapi.TransactionContextInterface, q string) ([]*AssetData, error) {

	itr, err := ctx.GetStub().GetQueryResult(q)
	if err != nil {
		return nil, err
	}

	defer itr.Close()

	entries := make([]*AssetData, 0)

	for itr.HasNext() {
		kv, err := itr.Next()
		if err != nil {
			return nil, err
		}

		var entry AssetData
		if err := json.Unmarshal(kv.Value, &entry); err != nil {
			return nil, err
		}

		entries = append(entries, &entry)
	}

	return entries, nil
}
