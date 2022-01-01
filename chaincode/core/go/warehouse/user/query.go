package user

import (
	"encoding/json"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func Query(ctx contractapi.TransactionContextInterface, uniqueId string) (*RegisterationData, error) {

	sBytes, err := ctx.GetStub().GetState(uniqueId)
	if err != nil {
		return nil, err
	}

	var r RegisterationData
	if err := json.Unmarshal(sBytes, &r); err != nil {
		return nil, err
	}

	return &r, nil
}
