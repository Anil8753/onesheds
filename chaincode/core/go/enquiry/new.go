package enquiry

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func NewEnquiry(
	ctx contractapi.TransactionContextInterface,
	enquiryId string,
	depositorId string,
	warehouseId string,
	input []byte,
) (*EnquiryData, error) {

	if enquiryId == "" {
		return nil, fmt.Errorf("enquiryId cannot be blank")
	}

	if depositorId == "" {
		return nil, fmt.Errorf("depositorId cannot be blank")
	}

	if warehouseId == "" {
		return nil, fmt.Errorf("warehouseId cannot be blank")
	}

	_, err := ctx.GetStub().GetState(enquiryId)
	if err == nil {
		return nil, fmt.Errorf("enquiryId '%s' already exist", enquiryId)
	}

	data := EnquiryData{
		DocType:   EnquiryDocType,
		Warehouse: warehouseId,
		Depositor: depositorId,
	}

	if err := json.Unmarshal(input, &data.Attributes); err != nil {
		return nil, fmt.Errorf("failed to unmarshal input. %w", err)
	}

	dataBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	if err := ctx.GetStub().PutState(enquiryId, dataBytes); err != nil {
		return nil, fmt.Errorf("failed to put state for the enquiryId: %s", enquiryId)
	}

	return &data, nil
}
