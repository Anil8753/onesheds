package txn

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/anil8753/onesheds/chaincode/core/asset"
	"github.com/anil8753/onesheds/chaincode/core/privatetxn"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func AcceptOrder(ctx contractapi.TransactionContextInterface) error {
	return ChangeOrderStatusByWarehouseOwner(ctx, OrderStatusAccepted)
}

func ChangeOrderStatusByWarehouseOwner(ctx contractapi.TransactionContextInterface, newStatus string) error {

	transient, err := ctx.GetStub().GetTransient()
	if err != nil {
		return fmt.Errorf("GetTransient failed. %w", err)
	}

	dataBytes, ok := transient["data"]
	if !ok {
		return errors.New("GetTransient does not have 'data'")
	}

	pl, err := NewChangeOrderStatusPayload(dataBytes)
	if err != nil {
		return fmt.Errorf("failed to create NewRejectOrderPayload. %w", err)
	}

	outBytes, pvtData, err := privatetxn.GetState(ctx, pl.OrderId)
	if err != nil {
		return err
	}

	var order Payload
	if err := json.Unmarshal(outBytes, &order); err != nil {
		return err
	}

	// Only warehouse owner can reject the order
	wh, err := asset.Query(ctx, order.WarehouseId)
	if err != nil {
		return fmt.Errorf("asset.Query failed. %w", err)
	}

	// check the warehouse ownership
	err = ctx.GetClientIdentity().AssertAttributeValue("userId", wh.OwnerId)
	if err != nil {
		return errors.New("unauthrorized client")
	}

	if newStatus == OrderStatusRejected {
		// Only new order can be rejected
		if order.Status != OrderStatusNew {
			return fmt.Errorf("only OrderStatusNew can be rejected. Current order status is: %s", order.Status)
		}

		order.Status = OrderStatusRejected

	} else if newStatus == OrderStatusAccepted {
		// Only new order can be accepted
		if order.Status != OrderStatusNew {
			return fmt.Errorf("only OrderStatusNew can be accepted. Current order status is: %s", order.Status)
		}

		order.Status = OrderStatusAccepted

	} else {
		return fmt.Errorf("invalid  order status is: %s", order.Status)
	}

	if err := privatetxn.PutState(ctx, pl.OrderId, order, pvtData.MSPOrgs, pvtData); err != nil {
		return err
	}

	return nil
}

type ChangeOrderStatusPayload struct {
	OrderId string
}

func NewChangeOrderStatusPayload(in []byte) (*ChangeOrderStatusPayload, error) {
	pl := ChangeOrderStatusPayload{}
	if err := json.Unmarshal([]byte(in), &pl); err != nil {
		return nil, err
	}

	if pl.OrderId == "" {
		return nil, errors.New("orderId is mandatory")
	}

	return &pl, nil
}
