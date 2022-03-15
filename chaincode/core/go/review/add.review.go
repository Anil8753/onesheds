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
	userId string,
	warehouseId string,
	userRating float32,
	reviewText string,
) error {

	if userId == "" {
		return errors.New("userId is mandatory")
	}

	_, err := asset.Query(ctx, warehouseId)
	if err != nil {
		return fmt.Errorf("warehouse %s not found. Error: %w", warehouseId, err)
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
		return err
	}

	if err := ctx.GetStub().PutState(txnId, dataBytes); err != nil {
		return fmt.Errorf("failed to put state for the txnId: %s", txnId)
	}

	return nil
}
