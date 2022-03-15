package contract

import (
	"encoding/json"

	"github.com/anil8753/onesheds/chaincode/core/review"
	"github.com/anil8753/onesheds/chaincode/core/utils"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

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
