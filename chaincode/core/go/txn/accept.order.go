package txn

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/anil8753/onesheds/chaincode/core/asset"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func AcceptOrder(ctx contractapi.TransactionContextInterface, orderId string) error {

	outBytes, err := ctx.GetStub().GetState(orderId)
	if err != nil {
		return err
	}

	if len(outBytes) == 0 {
		return fmt.Errorf("%s key not found", orderId)
	}

	var order Order
	if err := json.Unmarshal(outBytes, &order); err != nil {
		return err
	}

	wh, err := asset.Query(ctx, order.WarehouseId)
	if err != nil {
		return fmt.Errorf("asset.Query failed. %w", err)
	}

	// Only warehouse owner can accept/reject the order
	err = ctx.GetClientIdentity().AssertAttributeValue("userId", wh.OwnerId)
	if err != nil {
		return errors.New("unauthrorized client")
	}

	// Only new order can be accepted
	if order.Status != OrderStatusNew {
		return fmt.Errorf("only OrderStatusNew can be accepted. Current order status is: %s", order.Status)
	}

	// update the order data
	order.Status = OrderStatusAccepted

	bytesOrder, err := json.Marshal(order)
	if err != nil {
		return err
	}

	if err := ctx.GetStub().PutState(orderId, bytesOrder); err != nil {
		return err
	}

	// todo: warehouse space deduction

	return nil
}
