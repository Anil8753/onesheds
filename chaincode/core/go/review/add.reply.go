package review

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func AddReply(
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

	newReply := Reply{
		Id:     ctx.GetStub().GetTxID(),
		UserId: userId,
		Text:   replyText,
	}

	// First check at root level
	if data.UserReview.Id == targetId {
		data.UserReview.Replies = append(data.UserReview.Replies, newReply)

	} else {
		// check at replies
		r := find(data.UserReview.Replies, targetId)
		if r == nil {
			return nil, fmt.Errorf("targetId %s not found and there are no reply", targetId)
		}

		r.Replies = append(r.Replies, newReply)
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

func find(replies []Reply, targetId string) *Reply {

	for i := range replies {

		if r := find(replies[i].Replies, targetId); r != nil {
			return r
		}

		if replies[i].Id == targetId {
			return &replies[i]
		}
	}

	return nil
}
