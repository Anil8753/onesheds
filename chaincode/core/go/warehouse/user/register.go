package user

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func NewRegisterationData(input string) (*RegisterationData, error) {

	rBytes := []byte(input)

	var regData RegisterationData
	if err := json.Unmarshal(rBytes, &regData); err != nil {
		return nil, err
	}

	if regData.UniqueId == "" {
		return nil, errors.New("UniqueId is mandatory")
	}

	regData.DocType = "userRegData"

	return &regData, nil
}

func Register(ctx contractapi.TransactionContextInterface, input string) error {

	u, err := NewRegisterationData(input)
	if err != nil {
		return fmt.Errorf("invalid registeration data. %w", err)
	}

	if _, err := Query(ctx, u.UniqueId); err == nil {
		return fmt.Errorf("user id (%s) already exist. %w", u.UniqueId, err)
	}

	b, err := json.Marshal(u)
	if err != nil {
		return err
	}

	if err := ctx.GetStub().PutState(u.UniqueId, b); err != nil {
		return fmt.Errorf("failed to register warehouse user. %w", err)
	}

	return nil
}
