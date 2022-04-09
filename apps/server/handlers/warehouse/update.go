package warehouse

import (
	"encoding/json"
	"net/http"

	"github.com/anil8753/onesheds/handlers/utils"
	"github.com/anil8753/onesheds/nethttp"
	"github.com/gin-gonic/gin"
)

func (s *Warehouse) UpdateWarehouseHandler() gin.HandlerFunc {
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
