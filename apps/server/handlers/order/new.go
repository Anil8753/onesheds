package order

import (
	"encoding/json"
	"net/http"

	"github.com/anil8753/onesheds/handlers/utils"
	"github.com/anil8753/onesheds/nethttp"
	"github.com/gin-gonic/gin"
)

// Binding from JSON
type NewOrderPostData struct {
	WarehouseId string `json:"warehouseId" binding:"required"`
	// DepositorId string `json:"userId" binding:"required"`
	// FromDate    time.Time `json:"fromDate" binding:"required"`
	Attrs interface{} `json:"attrs" binding:"required"`
}

func (s *Handler) NewOrder() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		reqData := NewOrderPostData{}
		if err := ctx.ShouldBindJSON(&reqData); err != nil {
			nethttp.ServerResponse(ctx, http.StatusBadRequest, nethttp.InvalidRequestData, err)
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

		id, err := utils.GenerateUUID("order")
		if err != nil {
			nethttp.ServerResponse(ctx, http.StatusInternalServerError, nethttp.ServerIssue, err)
			return
		}

		d := &OrderData{}
		d.Id = id
		d.DepositorId = udata.UserId
		d.WarehouseId = reqData.WarehouseId
		d.Attrs = reqData.Attrs

		outBytes, err := json.Marshal(d)
		if err != nil {
			nethttp.ServerResponse(ctx, http.StatusInternalServerError, nethttp.ServerIssue, err)
			return
		}

		resp, err := contract.SubmitTransaction("NewOrder", string(outBytes))
		if err != nil {
			nethttp.ServerResponse(ctx, http.StatusInternalServerError, nethttp.ServerIssue, err)
			return
		}

		nethttp.ServerResponse(ctx, http.StatusOK, nethttp.Success, string(resp))
	}
}
