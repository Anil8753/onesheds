package privatetxn

import (
	"encoding/json"
	"errors"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func PutState(
	ctx contractapi.TransactionContextInterface,
	key string,
	value interface{},
	mspOrgs []string,
	pvtData *PrivateData,
) error {

	if pvtData.Secret == "" {
		return errors.New("secret is missing in PrivateData")
	}

	b, err := json.Marshal(value)
	if err != nil {
		return err
	}

	outBytes, err := encrypt(pvtData.Secret, b)
	if err != nil {
		return err
	}

	if err := ctx.GetStub().PutState(key, outBytes); err != nil {
		return err
	}

	pvtBytes, err := json.Marshal(value)
	if err != nil {
		return err
	}

	for _, mspId := range mspOrgs {
		if err := ctx.GetStub().PutPrivateData(getImplicitPrivateCollection(mspId), key, pvtBytes); err != nil {
			return err
		}
	}

	return nil
}
