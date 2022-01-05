package user

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func Query(ctx contractapi.TransactionContextInterface, uniqueId string) (*RegisterationData, error) {

	sBytes, err := ctx.GetStub().GetState(uniqueId)
	if err != nil {
		return nil, err
	}

	if len(sBytes) == 0 {
		return nil, fmt.Errorf("'%s' user is not found", uniqueId)
	}

	var r RegisterationData
	if err := json.Unmarshal(sBytes, &r); err != nil {
		return nil, err
	}

	return &r, nil
}
