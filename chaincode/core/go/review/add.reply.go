package review

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func AddReply(
	ctx contractapi.TransactionContextInterface,
	id string,
	userId string,
	targetId string,
	text string,
) error {

	if id == "" {
		return errors.New("id is mandatory")
	}

	if userId == "" {
		return errors.New("userId is mandatory")
	}

	if targetId == "" {
		return errors.New("targetId is mandatory")
	}

	data, err := Get(ctx, id)
	if err != nil {
		return fmt.Errorf("review %s not found. Error: %w", id, err)
	}

	// First check at root level
	if data.UserReview.Id == targetId {

		data.UserReview.Replies = append(
			data.UserReview.Replies,
			Reply{
				Id:     ctx.GetStub().GetTxID(),
				UserId: userId,
				Text:   text,
			},
		)

		return nil
	}

	// check at replies
	r := find(data.UserReview.Replies, targetId)
	if r == nil {
		return fmt.Errorf("targetId %s not found and there are no reply", targetId)
	}

	r.Replies = append(r.Replies, Reply{
		Id:     ctx.GetStub().GetTxID(),
		UserId: userId,
		Text:   text,
	})

	dataBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if err := ctx.GetStub().PutState(id, dataBytes); err != nil {
		return fmt.Errorf("failed to put state for the id: %s", id)
	}

	return nil
}

func find(replies []Reply, targetId string) *Reply {

	for _, reply := range replies {

		if r := find(reply.Replies, targetId); r != nil {
			return r
		}

		if reply.Id == targetId {
			return &reply
		}
	}

	return nil
}
