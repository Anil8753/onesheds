package knowledgebase

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/anil8753/onesheds/chaincode/core/utils"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func AddQuestion(
	ctx contractapi.TransactionContextInterface,
	warehouseId string,
	question string,
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
		doc = &Document{Id: id, WarehouseId: warehouseId, KnowledgeBase: nil}
	}

	if err := CanAddQuestion(ctx, doc); err != nil {
		return nil, err
	}

	doc.KnowledgeBase = append(doc.KnowledgeBase, Entry{Question: question, Questioner: userId})

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
func CanAddQuestion(
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

	if mspId == utils.DepositorMSP {
		return nil
	}

	return errors.New("unauthorized user")
}
