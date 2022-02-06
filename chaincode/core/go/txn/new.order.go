package txn

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func NewOrder(ctx contractapi.TransactionContextInterface, input string) (*Order, error) {

	order, err := ParsePayload([]byte(input))
	if err != nil {
		return nil, fmt.Errorf("failed to create NewPayload. %w", err)
	}

	if !strings.HasPrefix(order.Id, "order") {
		return nil, fmt.Errorf("txn id must start with 'order'")
	}

	order.DocType = OrderDocType
	order.Status = OrderStatusNew

	if _, err := ctx.GetStub().GetState(order.Id); err != nil {
		return nil, fmt.Errorf("OrderId (%s) already exist", order.Id)
	}

	bytesOrder, err := json.Marshal(order)
	if err != nil {
		return nil, err
	}

	if err := ctx.GetStub().PutState(order.Id, bytesOrder); err != nil {
		return nil, fmt.Errorf("ctx.PutState failed . %w", err)
	}

	// todo: warehouse space deduction

	return order, nil
}
