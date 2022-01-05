package user

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func Update(ctx contractapi.TransactionContextInterface, input string) error {

	// Todo: check the profile ownership from the user x509 certificate

	uRegData, err := NewRegisterationData(input)
	if err != nil {
		return fmt.Errorf("invalid registeration data. %w", err)
	}

	sRegData, err := Query(ctx, uRegData.UniqueId)
	if err != nil {
		return fmt.Errorf("%s user not found in state db. %w", uRegData.UniqueId, err)
	}

	if uRegData.UniqueId != sRegData.UniqueId {
		return fmt.Errorf("passed %s UniqueId not same as state UniqueId %s", uRegData.UniqueId, sRegData.UniqueId)
	}

	updated, err := updateRegData(uRegData, sRegData)
	if err != nil {
		return err
	}

	b, err := json.Marshal(updated)
	if err != nil {
		return err
	}

	if err := ctx.GetStub().PutState(updated.UniqueId, b); err != nil {
		return fmt.Errorf("failed to register warehouse user. %w", err)
	}

	return nil
}

// updateRegData updated the field from 'from' to 'to' RegisterationData
func updateRegData(from *RegisterationData, to *RegisterationData) (*RegisterationData, error) {

	fBytes, err := json.Marshal(from)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(fBytes, to); err != nil {
		return nil, err
	}

	return to, nil
}
