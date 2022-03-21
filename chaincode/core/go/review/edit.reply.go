package review

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func EditReply(
	ctx contractapi.TransactionContextInterface,
	reviewId string,
	targetId string,
	replyText string,
) (*Entry, error) {

	userId, found, err := ctx.GetClientIdentity().GetAttributeValue("userId")
	if err != nil || !found {
		return nil, errors.New("failed to get user from certificate")
	}

	if targetId == "" {
		return nil, errors.New("targetId is mandatory")
	}

	data, err := Get(ctx, reviewId)
	if err != nil {
		return nil, fmt.Errorf("review %s not found. Error: %w", reviewId, err)
	}

	// First check at root level
	if data.UserReview.Id == targetId {

		if data.UserReview.UserId != userId {
			return nil, fmt.Errorf("unauthrorized user: %s", userId)
		}

		data.UserReview.Text = replyText

	} else {
		// check at replies
		r := find(data.UserReview.Replies, targetId)
		if r == nil {
			return nil, fmt.Errorf("targetId %s not found and there are no reply", targetId)
		}

		if r.UserId != userId {
			return nil, fmt.Errorf("unauthrorized user: %s", userId)
		}

		r.Text = replyText
	}

	dataBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	if err := ctx.GetStub().PutState(reviewId, dataBytes); err != nil {
		return nil, fmt.Errorf("failed to put state for the reviewId: %s", reviewId)
	}

	return data, nil
}
