package enquiry

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/anil8753/onesheds/chaincode/core/utils"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func UpdateEnquiry(
	ctx contractapi.TransactionContextInterface,
	enquiryId string,
	input []byte,
) (*EnquiryData, error) {

	outBytes, err := ctx.GetStub().GetState(enquiryId)
	if err != nil {
		return nil, err
	}

	if len(outBytes) == 0 {
		return nil, fmt.Errorf("%s enquiryId not found", enquiryId)
	}

	var data EnquiryData
	if err := json.Unmarshal(outBytes, &data); err != nil {
		return nil, err
	}

	if err := CanUpdate(ctx, &data); err != nil {
		return nil, err
	}

	if err := json.Unmarshal(input, &data.Attributes); err != nil {
		return nil, fmt.Errorf("failed to unmarshal input. %w", err)
	}

	dataBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	if err := ctx.GetStub().PutState(enquiryId, dataBytes); err != nil {
		return nil, fmt.Errorf("failed to put state for the key: %s", enquiryId)
	}

	return &data, nil
}

// Enquiry can be updated by regulator or depositor only
func CanUpdate(
	ctx contractapi.TransactionContextInterface,
	data *EnquiryData,
) error {

	mspId, err := ctx.GetClientIdentity().GetMSPID()
	if err != nil {
		return fmt.Errorf("failed to get MSP. %w", err)
	}

	// If user is regulator
	if mspId == utils.RegulatorMSP {
		return nil
	}

	// If user is depositor
	if mspId == utils.DepositorMSP {
		if err := ctx.GetClientIdentity().AssertAttributeValue("userId", data.Depositor); err != nil {
			return fmt.Errorf("unathourized depositor. %w", err)
		}

		return nil
	}

	return errors.New("unauthorized user")
}