package asset

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-chaincode-go/shim"
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
	return QueryWithParams(ctx, q)
}

func QueryByOwnerId(ctx contractapi.TransactionContextInterface, ownerId string) ([]*AssetData, error) {

	qFmt := `{ "selector": { "ownerId": "%s" } }`
	q := fmt.Sprintf(qFmt, ownerId)

	return QueryWithParams(ctx, q)
}

func QueryWithParams(ctx contractapi.TransactionContextInterface, q string) ([]*AssetData, error) {

	itr, err := ctx.GetStub().GetQueryResult(q)
	if err != nil {
		return nil, err
	}

	defer itr.Close()

	entries, err := constructQueryResponseFromIterator(itr)
	if err != nil {
		return nil, err
	}

	return entries, nil
}

// PaginatedQueryResult structure used for returning paginated query results and metadata
type PaginatedQueryResult struct {
	Records             []*AssetData `json:"records"`
	FetchedRecordsCount int32        `json:"fetchedRecordsCount"`
	Bookmark            string       `json:"bookmark"`
}

func QueryWithPagination(ctx contractapi.TransactionContextInterface, q string, pageSize int, bookmark string) (*PaginatedQueryResult, error) {

	itr, responseMetadata, err := ctx.GetStub().GetQueryResultWithPagination(q, int32(pageSize), bookmark)
	if err != nil {
		return nil, err
	}
	defer itr.Close()

	assets, err := constructQueryResponseFromIterator(itr)
	if err != nil {
		return nil, err
	}

	return &PaginatedQueryResult{
		Records:             assets,
		FetchedRecordsCount: responseMetadata.FetchedRecordsCount,
		Bookmark:            responseMetadata.Bookmark,
	}, nil
}

// constructQueryResponseFromIterator constructs a slice of assets from the resultsIterator
func constructQueryResponseFromIterator(resultsIterator shim.StateQueryIteratorInterface) ([]*AssetData, error) {

	assets := make([]*AssetData, 0)

	for resultsIterator.HasNext() {
		queryResult, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		var asset AssetData
		err = json.Unmarshal(queryResult.Value, &asset)
		if err != nil {
			return nil, err
		}
		assets = append(assets, &asset)
	}

	return assets, nil
}
