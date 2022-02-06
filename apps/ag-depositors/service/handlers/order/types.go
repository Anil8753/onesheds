package order

import (
	"github.com/anil8753/onesheds/apps/warehousemen/service/interfaces"
	"github.com/anil8753/onesheds/apps/warehousemen/service/ledger"
)

type Handler struct {
	Database interfaces.Database
	Ledger   *ledger.Ledger
}

type OrderData struct {
	Id          string      `json:"id"`
	WarehouseId string      `json:"warehouseId"`
	DepositorId string      `json:"depositorId"`
	Attrs       interface{} `json:"attrs"`
}
