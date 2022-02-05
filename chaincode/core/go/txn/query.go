package txn

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/anil8753/onesheds/chaincode/core/utils"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func QueryOrder(
	ctx contractapi.TransactionContextInterface,
	orderId string,
) (*Order, error) {

	sBytes, err := ctx.GetStub().GetState(orderId)
	if err != nil {
		return nil, err
	}

	if len(sBytes) == 0 {
		return nil, fmt.Errorf("'%s' order is not found", orderId)
	}

	var order Order
	if err := json.Unmarshal(sBytes, &order); err != nil {
		return nil, err
	}

	if err := isReadAllowed(ctx, &order); err != nil {
		return nil, err
	}

	return &order, nil
}

func QueryDepositorAllOrders(
	ctx contractapi.TransactionContextInterface,
	depositorId string,
) ([]*Order, error) {

	clientMSP, err := ctx.GetClientIdentity().GetMSPID()
	if err != nil {
		return nil, err
	}

	err = ctx.GetClientIdentity().AssertAttributeValue("userId", depositorId)
	if err != nil && clientMSP != utils.RegulatorMSP {
		return nil, errors.New("only regulator or actual depositor can access")
	}

	qFmt := `{ "selector": { "docType": "%s", "depositorId": "%s" } }`
	q := fmt.Sprintf(qFmt, OrderDocType, depositorId)
	return richQuery(ctx, q)
}

func QueryWarehouseAllOrders(
	ctx contractapi.TransactionContextInterface,
	warehouseId string,
) ([]*Order, error) {

	clientMSP, err := ctx.GetClientIdentity().GetMSPID()
	if err != nil {
		return nil, err
	}

	err = IsUserWarehouseOwner(ctx, warehouseId)
	if err != nil && clientMSP != utils.RegulatorMSP {
		return nil, errors.New("only regulator or warehouseman can access")
	}

	qFmt := `{ "selector": { "docType": "%s", "warehouseId": "%s" } }`
	q := fmt.Sprintf(qFmt, OrderDocType, warehouseId)
	return richQuery(ctx, q)
}

func richQuery(
	ctx contractapi.TransactionContextInterface,
	q string,
) ([]*Order, error) {

	itr, err := ctx.GetStub().GetQueryResult(q)
	if err != nil {
		return nil, err
	}

	defer itr.Close()

	entries, err := constructQueryResponseFromIterator(ctx, itr)
	if err != nil {
		return nil, err
	}

	return entries, nil
}

// PaginatedQueryResult structure used for returning paginated query results and metadata
type PaginatedQueryResult struct {
	Records             []*Order `json:"records"`
	FetchedRecordsCount int32    `json:"fetchedRecordsCount"`
	Bookmark            string   `json:"bookmark"`
}

func QueryWarehouseAllOrdersWithPagination(
	ctx contractapi.TransactionContextInterface,
	warehouseId string,
	pageSize int,
	bookmark string,
) (*PaginatedQueryResult, error) {

	clientMSP, err := ctx.GetClientIdentity().GetMSPID()
	if err != nil {
		return nil, err
	}

	err = IsUserWarehouseOwner(ctx, warehouseId)
	if err != nil && clientMSP != utils.RegulatorMSP {
		return nil, errors.New("only regulator or warehouseman can access")
	}

	qFmt := `{ "selector": { "docType": "%s", "warehouseId": "%s" } }`
	q := fmt.Sprintf(qFmt, OrderDocType, warehouseId)

	itr, responseMetadata, err := ctx.GetStub().GetQueryResultWithPagination(q, int32(pageSize), bookmark)
	if err != nil {
		return nil, err
	}
	defer itr.Close()

	assets, err := constructQueryResponseFromIterator(ctx, itr)
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
func constructQueryResponseFromIterator(
	ctx contractapi.TransactionContextInterface,
	resultsIterator shim.StateQueryIteratorInterface,
) ([]*Order, error) {

	entries := make([]*Order, 0)

	for resultsIterator.HasNext() {
		queryResult, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		var entry Order
		err = json.Unmarshal(queryResult.Value, &entry)
		if err != nil {
			return nil, err
		}

		entries = append(entries, &entry)
	}

	return entries, nil
}
