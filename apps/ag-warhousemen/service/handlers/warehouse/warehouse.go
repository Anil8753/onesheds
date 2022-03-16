package warehouse

import (
	"encoding/json"
	"net/http"

	"github.com/anil8753/onesheds/apps/warehousemen/service/handlers/utils"
	"github.com/anil8753/onesheds/apps/warehousemen/service/nethttp"
	"github.com/gin-gonic/gin"
)

func (s *Asset) GetWarehousesHandler() gin.HandlerFunc {
	//
	return func(ctx *gin.Context) {

		udata := utils.GetUserFromContext(ctx, s.Database)
		if udata == nil {
			return
		}

		contract, err := s.Ledger.GetUserContract(udata.Crypto)
		if err != nil {
			nethttp.ServerResponse(ctx, http.StatusInternalServerError, nethttp.ServerIssue, err)
			return
		}

		resp, err := contract.EvaluateTransaction("GetWarehouseByOwnerId", udata.UserId)
		if err != nil {
			nethttp.ServerResponse(ctx, http.StatusInternalServerError, nethttp.ServerIssue, err)
			return
		}

		nethttp.ServerResponse(ctx, http.StatusOK, nethttp.Success, string(resp))
	}
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

func (s *Asset) CreateWarehouseHandler() gin.HandlerFunc {
	//
	return func(ctx *gin.Context) {

		var data AssetData
		if err := ctx.BindJSON(&data); err != nil {
			nethttp.ServerResponse(ctx, http.StatusBadRequest, nethttp.InvalidRequestData, err)
			return
		}

		udata := utils.GetUserFromContext(ctx, s.Database)
		if udata == nil {
			nethttp.ServerResponse(ctx, http.StatusInternalServerError, nethttp.ServerIssue, "failed to get user from context")
			return
		}

		warehouseId, err := utils.GenerateUUID("warehouse")
		if err != nil {
			nethttp.ServerResponse(ctx, http.StatusInternalServerError, nethttp.ServerIssue, err)
			return
		}

		data.WarehouseId = warehouseId
		data.OwnerId = udata.UserId

		inBytes, err := json.Marshal(data)
		if err != nil {
			nethttp.ServerResponse(ctx, http.StatusInternalServerError, nethttp.ServerIssue, err)
			return
		}

		contract, err := s.Ledger.GetUserContract(udata.Crypto)
		if err != nil {
			nethttp.ServerResponse(ctx, http.StatusInternalServerError, nethttp.ServerIssue, err)
			return
		}

		resp, err := contract.SubmitTransaction("RegisterWarehouse", string(inBytes))
		if err != nil {
			nethttp.ServerResponse(ctx, http.StatusInternalServerError, nethttp.ServerIssue, err)
			return
		}

		nethttp.ServerResponse(ctx, http.StatusOK, nethttp.Success, string(resp))
	}
}

func (s *Asset) UpdateWarehouseHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var data AssetData
		if err := ctx.BindJSON(&data); err != nil {
			nethttp.ServerResponse(ctx, http.StatusBadRequest, nethttp.InvalidRequestData, err)
			return
		}

		inBytes, err := json.Marshal(data)
		if err != nil {
			nethttp.ServerResponse(ctx, http.StatusInternalServerError, nethttp.ServerIssue, err)
			return
		}

		udata := utils.GetUserFromContext(ctx, s.Database)
		if udata == nil {
			return
		}

		contract, err := s.Ledger.GetUserContract(udata.Crypto)
		if err != nil {
			nethttp.ServerResponse(ctx, http.StatusInternalServerError, nethttp.ServerIssue, err)
			return
		}

		resp, err := contract.SubmitTransaction("UpdateWarehouse", string(inBytes))
		if err != nil {
			nethttp.ServerResponse(ctx, http.StatusInternalServerError, nethttp.ServerIssue, err)
			return
		}

		nethttp.ServerResponse(ctx, http.StatusOK, nethttp.Success, string(resp))
	}
}
