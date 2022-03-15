package contract

import (
	"encoding/json"
	"fmt"

	"github.com/anil8753/onesheds/chaincode/core/review"
	"github.com/anil8753/onesheds/chaincode/core/utils"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func GetAllReviews(
	ctx contractapi.TransactionContextInterface,
	warehouseId string,
) (string, error) {
	qFmt := `{ "selector": { "docType": "%s", "warehouseId": "%s" } }`
	q := fmt.Sprintf(qFmt, review.ReviewDocType, warehouseId)
	r, err := review.RichQuery(ctx, q)
	if err != nil {
		return "", err
	}

	return utils.ToJSON(r)
}

func GetReview(
	ctx contractapi.TransactionContextInterface,
	reviewId string,
) (string, error) {

	r, err := review.Get(ctx, reviewId)
	if err != nil {
		return "", err
	}

	return utils.ToJSON(r)
}

type AddReviewData struct {
	WarehouseId string  `json:"warehouseId"`
	UserRating  float32 `json:"userRating"`
	ReviewText  string  `json:"reviewText"`
}

func AddUserRating(
	ctx contractapi.TransactionContextInterface,
	input string,
) (string, error) {

	var in AddReviewData
	if err := json.Unmarshal([]byte(input), &in); err != nil {
		return "", err
	}

	r, err := review.AddUserRating(ctx, in.WarehouseId, in.UserRating, in.ReviewText)
	if err != nil {
		return "", err
	}

	return utils.ToJSON(r)
}

type AddReplyData struct {
	ReviewId  string `json:"reviewId"`
	TargetId  string `json:"targetId"`
	ReplyText string `json:"replyText"`
}

func AddReply(
	ctx contractapi.TransactionContextInterface,
	input string,
) (string, error) {

	var in AddReplyData
	if err := json.Unmarshal([]byte(input), &in); err != nil {
		return "", err
	}

	r, err := review.AddReply(ctx, in.ReviewId, in.TargetId, in.ReplyText)
	if err != nil {
		return "", err
	}

	return utils.ToJSON(r)
}
