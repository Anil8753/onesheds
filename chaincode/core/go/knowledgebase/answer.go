package knowledgebase

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/anil8753/onesheds/chaincode/core/asset"
	"github.com/anil8753/onesheds/chaincode/core/utils"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func AddAnswer(
	ctx contractapi.TransactionContextInterface,
	warehouseId string,
	index int,
	answerString string,
) (*Document, error) {

	userId, found, err := ctx.GetClientIdentity().GetAttributeValue("userId")
	if err != nil || !found {
		return nil, errors.New("failed to get user from certificate")
	}

	id, err := ctx.GetStub().CreateCompositeKey(IDPrefix, []string{warehouseId})
	if err != nil {
		return nil, fmt.Errorf("failed to create composit key. %w", err)
	}

	doc, err := get(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get FAQ for the warehouse. %s. Error: %w", id, err)
	}

	if len(doc.KnowledgeBase) <= index {
		return nil, fmt.Errorf("invalid index: %d", index)
	}

	if err := CanAddAnswer(ctx, doc); err != nil {
		return nil, err
	}

	var answer Answer
	if err := json.Unmarshal([]byte(answerString), &answer); err != nil {
		return nil, fmt.Errorf("invalid input data. %w", err)
	}

	old := doc.KnowledgeBase[index]
	old.Answer = answer
	old.Answerer = userId // can be warehousenowner or regulator

	doc.KnowledgeBase[index] = old

	outBytes, err := json.Marshal(doc)
	if err != nil {
		return nil, err
	}

	if err := ctx.GetStub().PutState(id, outBytes); err != nil {
		return nil, fmt.Errorf("failed to put data. %w", err)
	}

	return doc, nil
}

// can be updated by regulator or warehousemen only
func CanAddAnswer(
	ctx contractapi.TransactionContextInterface,
	doc *Document,
) error {
	mspId, err := ctx.GetClientIdentity().GetMSPID()
	if err != nil {
		return fmt.Errorf("failed to get MSP. %w", err)
	}

	// If user is regulator
	if mspId == utils.RegulatorMSP {
		return nil
	}

	wh, err := asset.Query(ctx, doc.WarehouseId)
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
