package review

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func EditUserReview(
	ctx contractapi.TransactionContextInterface,
	reviewId string,
	userRating float32,
	reviewText string,
) (*Entry, error) {
	userId, found, err := ctx.GetClientIdentity().GetAttributeValue("userId")
	if err != nil || !found {
		return nil, errors.New("failed to get user from certificate")
	}

	outBytes, err := ctx.GetStub().GetState(reviewId)
	if err != nil {
		return nil, err
	}

	if len(outBytes) == 0 {
		return nil, fmt.Errorf("key: %s not found", reviewId)
	}

	var data Entry
	if err := json.Unmarshal(outBytes, &data); err != nil {
		return nil, err
	}

	if data.UserReview.UserId != userId {
		return nil, fmt.Errorf("unauthrorized user: %s", userId)
	}

	data.UserReview.Rating = userRating
	data.UserReview.Text = reviewText

	dataBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	if err := ctx.GetStub().PutState(reviewId, dataBytes); err != nil {
		return nil, fmt.Errorf("failed to put state for the key: %s", reviewId)
	}

	return &data, nil
}
