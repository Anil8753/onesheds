package review

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func Get(
	ctx contractapi.TransactionContextInterface,
	key string,
) (*Entry, error) {

	outBytes, err := ctx.GetStub().GetState(key)
	if err != nil {
		return nil, fmt.Errorf("failed to get data for key: %s. Error: %w", key, err)
	}

	var data Entry
	if err := json.Unmarshal(outBytes, &data); err != nil {
		return nil, err
	}

	return &data, nil
}
