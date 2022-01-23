package txn

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/anil8753/onesheds/chaincode/core/privatetxn"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func NewOrder(ctx contractapi.TransactionContextInterface) error {

	transient, err := ctx.GetStub().GetTransient()
	if err != nil {
		return fmt.Errorf("GetTransient failed. %w", err)
	}

	dataBytes, ok := transient["data"]
	if !ok {
		return errors.New("GetTransient does not have 'data'")
	}

	pl, err := NewPayload(dataBytes)
	if err != nil {
		return fmt.Errorf("failed to create NewPayload. %w", err)
	}

	if _, _, err := privatetxn.GetState(ctx, pl.OrderId); err == nil {
		return fmt.Errorf("OrderId (%s) already exist", pl.OrderId)
	}

	if err := privatetxn.PutState(ctx, pl.OrderId, pl.Value, pl.pvtData.MSPOrgs, &pl.pvtData); err != nil {
		return fmt.Errorf("privatetxn.PutState failed . %w", err)
	}

	return nil
}

type PayloadWrapper struct {
	OrderId string
	Value   Payload
	pvtData privatetxn.PrivateData
}

type Payload struct {
	WarehouseId string
	DepositorId string
	Status      string
	Attr        interface{}
}

func NewPayload(input []byte) (*PayloadWrapper, error) {

	var pl PayloadWrapper

	if err := json.Unmarshal([]byte(input), &pl); err != nil {
		return nil, err
	}

	if pl.OrderId == "" {
		return nil, errors.New("PayloadWrapper.OrderId is mandatory")
	}

	if pl.Value.DepositorId == "" {
		return nil, errors.New("PayloadWrapper.Value.DepositorId is mandatory")
	}

	if pl.Value.WarehouseId == "" {
		return nil, errors.New("PayloadWrapper.Value.WarehouseId is mandatory")
	}

	if pl.Value.Attr == "" {
		return nil, errors.New("PayloadWrapper.Value.Attr is mandatory")
	}

	if pl.pvtData.Secret == "" {
		return nil, errors.New("PayloadWrapper.PrivateData.Secret is mandatory")
	}

	if pl.pvtData.MSPOrgs == nil || len(pl.pvtData.MSPOrgs) == 0 {
		return nil, errors.New("PayloadWrapper.MSPOrgs is mandatory")
	}

	pl.Value.Status = OrderStatusNew

	return &pl, nil
}
