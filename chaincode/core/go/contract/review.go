package contract

import (
	"encoding/json"
	"fmt"

	"github.com/anil8753/onesheds/chaincode/core/review"
	"github.com/anil8753/onesheds/chaincode/core/utils"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func (c *Contract) GetAllReviews(
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

func (c *Contract) GetReview(
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

func (c *Contract) AddUserRating(
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

type EditReviewData struct {
	ReviewId   string  `json:"reviewId"`
	UserRating float32 `json:"userRating"`
	ReviewText string  `json:"reviewText"`
}

func (c *Contract) EditUserReview(
	ctx contractapi.TransactionContextInterface,
	input string,
) (string, error) {

	var in EditReviewData
	if err := json.Unmarshal([]byte(input), &in); err != nil {
		return "", err
	}

	r, err := review.EditUserReview(ctx, in.ReviewId, in.UserRating, in.ReviewText)
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

func (c *Contract) AddReply(
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

type EditReplyData struct {
	ReviewId  string `json:"reviewId"`
	TargetId  string `json:"targetId"`
	ReplyText string `json:"replyText"`
}

func (c *Contract) EditReply(
	ctx contractapi.TransactionContextInterface,
	input string,
) (string, error) {

	var in EditReplyData
	if err := json.Unmarshal([]byte(input), &in); err != nil {
		return "", err
	}

	r, err := review.EditReply(ctx, in.ReviewId, in.TargetId, in.ReplyText)
	if err != nil {
		return "", err
	}

	return utils.ToJSON(r)
}
