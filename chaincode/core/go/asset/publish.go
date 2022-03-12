package asset

import (
	"encoding/json"
	"fmt"

	utils "github.com/anil8753/onesheds/chaincode/core/utils"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func Publish(ctx contractapi.TransactionContextInterface, warehouseId string) (*AssetData, error) {

	clientMSP, err := ctx.GetClientIdentity().GetMSPID()
	if err != nil {
		return nil, fmt.Errorf("failed to get client msp. %w", err)
	}

	// check the authority
	if clientMSP != utils.RegulatorMSP {
		return nil, fmt.Errorf("unauthrorized node. Only %s can publish warehouse", utils.RegulatorMSP)
	}

	sData, err := Query(ctx, warehouseId)
	if err != nil {
		return nil, fmt.Errorf("%s warehouse not found in state db. %w", warehouseId, err)
	}

	// Change the status to published
	sData.Status = StatusPublished

	b, err := json.Marshal(sData)
	if err != nil {
		return nil, err
	}

	if err := ctx.GetStub().PutState(warehouseId, b); err != nil {
		return nil, fmt.Errorf("failed to register warehouse. %w", err)
	}

	return sData, nil
}
