package txn

import (
	"encoding/json"
	"errors"
)

const (
	OrderStatusNew       = "OrderStatusNew"
	OrderStatusRejected  = "OrderStatusRejected"
	OrderStatusAccepted  = "OrderStatusAccepted"
	OrderStatusCompleted = "OrderStatusCompleted"
)

const OrderDocType = "Order"

type Order struct {
	DocType     string
	Id          string
	WarehouseId string
	DepositorId string
	Status      string
	Attr        interface{}
}

func ParsePayload(input []byte) (*Order, error) {

	var order Order

	if err := json.Unmarshal([]byte(input), &order); err != nil {
		return nil, err
	}

	if order.Id == "" {
		return nil, errors.New("id is mandatory")
	}

	if order.DepositorId == "" {
		return nil, errors.New("DepositorId is mandatory")
	}

	if order.WarehouseId == "" {
		return nil, errors.New("WarehouseId is mandatory")
	}

	if order.Attr == nil {
		return nil, errors.New("Attr is mandatory")
	}

	return &order, nil
}
