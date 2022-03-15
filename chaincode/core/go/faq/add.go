package faq

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/anil8753/onesheds/chaincode/core/asset"
	"github.com/anil8753/onesheds/chaincode/core/utils"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func Add(
	ctx contractapi.TransactionContextInterface,
	warehouseId string,
	question string,
	input string,
) (*Entry, error) {

	id, err := ctx.GetStub().CreateCompositeKey(IDPrefix, []string{warehouseId})
	if err != nil {
		return nil, fmt.Errorf("failed to create composit key. %w", err)
	}

	entry, err := Get(ctx, id)
	if err != nil {
		entry = &Entry{Id: id, WarehouseId: warehouseId}
	}

	if err := CanAdd(ctx, entry); err != nil {
		return nil, err
	}

	var answer Answer
	if err := json.Unmarshal([]byte(input), &answer); err != nil {
		return nil, fmt.Errorf("invalid input data. %w", err)
	}

	entry.FAQs = append(entry.FAQs, FAQ{Question: question, Answer: answer})

	outBytes, err := json.Marshal(entry)
	if err != nil {
		return nil, err
	}

	if err := ctx.GetStub().PutState(id, outBytes); err != nil {
		return nil, fmt.Errorf("failed to put data. %w", err)
	}

	return entry, nil
}

// can be updated by regulator or warehousemen only
func CanAdd(
	ctx contractapi.TransactionContextInterface,
	data *Entry,
) error {

	mspId, err := ctx.GetClientIdentity().GetMSPID()
	if err != nil {
		return fmt.Errorf("failed to get MSP. %w", err)
	}

	// If user is regulator
	if mspId == utils.RegulatorMSP {
		return nil
	}

	wh, err := asset.Query(ctx, data.WarehouseId)
	if err != nil {
		return fmt.Errorf("asset.Query failed. %w", err)
	}

	// If user is warehouse owner
	if mspId == utils.WarehouseMSP {
		if err := ctx.GetClientIdentity().AssertAttributeValue("userId", wh.OwnerId); err != nil {
			return fmt.Errorf("unathourized warehousemen. %w", err)
		}

		return nil
	}

	return errors.New("unauthorized user")
}
