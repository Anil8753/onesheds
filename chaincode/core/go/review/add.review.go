package review

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/anil8753/onesheds/chaincode/core/asset"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func AddUserRating(
	ctx contractapi.TransactionContextInterface,
	warehouseId string,
	userRating float32,
	reviewText string,
) (*Entry, error) {

	userId, found, err := ctx.GetClientIdentity().GetAttributeValue("userId")
	if err != nil || !found {
		return nil, errors.New("failed to get user from certificate")
	}

	_, err = asset.Query(ctx, warehouseId)
	if err != nil {
		return nil, fmt.Errorf("warehouse %s not found. Error: %w", warehouseId, err)
	}

	txnId := ctx.GetStub().GetTxID()

	data := Entry{
		Id:          txnId,
		WarehouseId: warehouseId,
		UserReview: UserReview{
			Id:     txnId,
			UserId: userId,
			Text:   reviewText,
		},
	}

	dataBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	if err := ctx.GetStub().PutState(txnId, dataBytes); err != nil {
		return nil, fmt.Errorf("failed to put state for the txnId: %s", txnId)
	}

	return &data, nil
}
