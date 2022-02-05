package txn

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func NewOrder(ctx contractapi.TransactionContextInterface, input string) error {

	order, err := ParsePayload([]byte(input))
	if err != nil {
		return fmt.Errorf("failed to create NewPayload. %w", err)
	}

	if !strings.HasPrefix(order.Id, "order") {
		return fmt.Errorf("txn id must start with 'order'")
	}

	order.DocType = OrderDocType
	order.Status = OrderStatusNew

	if _, err := ctx.GetStub().GetState(order.Id); err == nil {
		return fmt.Errorf("OrderId (%s) already exist", order.Id)
	}

	bytesOrder, err := json.Marshal(order)
	if err != nil {
		return err
	}

	if err := ctx.GetStub().PutState(order.Id, bytesOrder); err != nil {
		return fmt.Errorf("ctx.PutState failed . %w", err)
	}

	// todo: warehouse space deduction

	return nil
}
