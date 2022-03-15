package contract

import (
	"github.com/anil8753/onesheds/chaincode/core/review"
	"github.com/anil8753/onesheds/chaincode/core/utils"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func AddUserRating(
	ctx contractapi.TransactionContextInterface,
	warehouseId string,
	userRating float32,
	reviewText string,
) (string, error) {

	r, err := review.AddUserRating(ctx, warehouseId, userRating, reviewText)
	if err != nil {
		return "", err
	}

	return utils.ToJSON(r)
}

func AddReply(
	ctx contractapi.TransactionContextInterface,
	reviewId string,
	targetId string,
	text string,
) (string, error) {

	r, err := review.AddReply(ctx, reviewId, targetId, text)
	if err != nil {
		return "", err
	}

	return utils.ToJSON(r)
}
