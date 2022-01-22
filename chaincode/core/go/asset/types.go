package asset

import (
	"encoding/json"
	"errors"
)

const regulatorMSP = "regulator-onesheds-com"

const (
	WarehouseDocType  = "WarehouseRegData"
	StatusUnpublished = "Unpublished"
	StatusPublished   = "Published"
	StatusSuspended   = "Suspended"
)

type AssetData struct {
	DocType string `json:"docType,omitempty"`

	WarehouseId     string   `json:"warehouseId"`
	Status          string   `json:"status,omitempty"`
	OwnerId         string   `json:"ownerId,omitempty"`
	TermsConditions []string `json:"termsConditions,omitempty"`

	Properties map[string]interface{} `json:"properties"`
	Photos     map[string]interface{} `json:"photos"`
	Videos     map[string]interface{} `json:"videos"`
}

func NewAssetData(input string) (*AssetData, error) {

	rBytes := []byte(input)

	var data AssetData
	if err := json.Unmarshal(rBytes, &data); err != nil {
		return nil, err
	}

	if data.WarehouseId == "" {
		return nil, errors.New("WarehouseId is mandatory")
	}

	data.DocType = WarehouseDocType

	if data.Properties == nil {
		data.Properties = make(map[string]interface{})
	}

	if data.Photos == nil {
		data.Photos = make(map[string]interface{})
	}

	if data.Videos == nil {
		data.Videos = make(map[string]interface{})
	}

	// if data.TermsConditions == nil {
	// 	data.TermsConditions = make([]string, 0)
	// }

	return &data, nil
}
