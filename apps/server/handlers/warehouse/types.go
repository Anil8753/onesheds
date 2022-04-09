package warehouse

import (
	"github.com/anil8753/onesheds/interfaces"
	"github.com/anil8753/onesheds/ledger"
)

type Warehouse struct {
	Database interfaces.Database
	Ledger   *ledger.Ledger
}

type AssetData struct {
	DocType     string `json:"docType,omitempty"`
	WarehouseId string `json:"warehouseId"`

	Status          string                 `json:"status,omitempty"`
	OwnerId         string                 `json:"ownerId,omitempty"`
	TermsConditions []string               `json:"termsConditions,omitempty"`
	Properties      map[string]interface{} `json:"properties,omitempty"`
	Photos          map[string]interface{} `json:"photos,omitempty"`
	Videos          map[string]interface{} `json:"videos,omitempty"`
}
