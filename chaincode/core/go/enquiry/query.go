package enquiry

import (
	"encoding/json"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func RichQuery(
	ctx contractapi.TransactionContextInterface,
	q string,
) ([]*EnquiryData, error) {

	itr, err := ctx.GetStub().GetQueryResult(q)
	if err != nil {
		return nil, err
	}

	defer itr.Close()

	entries, err := constructQueryResponseFromIterator(ctx, itr)
	if err != nil {
		return nil, err
	}

	return entries, nil
}

// constructQueryResponseFromIterator constructs a slice of assets from the resultsIterator
func constructQueryResponseFromIterator(
	ctx contractapi.TransactionContextInterface,
	resultsIterator shim.StateQueryIteratorInterface,
) ([]*EnquiryData, error) {

	entries := make([]*EnquiryData, 0)

	for resultsIterator.HasNext() {
		queryResult, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		var entry EnquiryData
		err = json.Unmarshal(queryResult.Value, &entry)
		if err != nil {
			return nil, err
		}

		entries = append(entries, &entry)
	}

	return entries, nil
}