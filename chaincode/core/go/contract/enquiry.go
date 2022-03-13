package contract

import (
	"fmt"

	"github.com/anil8753/onesheds/chaincode/core/enquiry"
	"github.com/anil8753/onesheds/chaincode/core/utils"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func (c *Contract) GetEnquiresByDepositor(
	ctx contractapi.TransactionContextInterface,
	depositorId string,
) (string, error) {

	qFmt := `{ "selector": { "docType": "%s", "depositor": "%s" } }`
	q := fmt.Sprintf(qFmt, enquiry.EnquiryDocType, depositorId)

	r, err := enquiry.RichQuery(ctx, q)
	if err != nil {
		return "", err
	}

	return utils.ToJSON(r)
}

func (c *Contract) GetEnquiresByWarehouse(
	ctx contractapi.TransactionContextInterface,
	warehouseId string,
) (string, error) {

	qFmt := `{ "selector": { "docType": "%s", "warehouse": "%s" } }`
	q := fmt.Sprintf(qFmt, enquiry.EnquiryDocType, warehouseId)

	r, err := enquiry.RichQuery(ctx, q)
	if err != nil {
		return "", err
	}

	return utils.ToJSON(r)
}

func NewEnquiry(
	ctx contractapi.TransactionContextInterface,
	enquiryId string,
	depositorId string,
	warehouseId string,
	input []byte,
) (string, error) {

	r, err := enquiry.NewEnquiry(ctx, enquiryId, depositorId, warehouseId, input)
	if err != nil {
		return "", err
	}

	return utils.ToJSON(r)
}

func UpdateEnquiry(
	ctx contractapi.TransactionContextInterface,
	enquiryId string,
	input []byte,
) (string, error) {

	r, err := enquiry.UpdateEnquiry(ctx, enquiryId, input)
	if err != nil {
		return "", err
	}

	return utils.ToJSON(r)
}
