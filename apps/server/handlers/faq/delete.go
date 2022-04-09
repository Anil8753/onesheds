package faq

import (
	"encoding/json"
	"net/http"

	"github.com/anil8753/onesheds/handlers/utils"
	"github.com/anil8753/onesheds/nethttp"
	"github.com/gin-gonic/gin"
)

type DeleteFAQDataPostData struct {
	WarehouseId string `json:"warehouseId" binding:"required"`
	Index       int    `json:"index"`
}

func (s *Handler) DeleteFAQ() gin.HandlerFunc {
	//
	return func(ctx *gin.Context) {

		reqData := DeleteFAQDataPostData{}
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

		outBytes, err := json.Marshal(reqData)
		if err != nil {
			nethttp.ServerResponse(ctx, http.StatusInternalServerError, nethttp.ServerIssue, err)
			return
		}

		resp, err := contract.SubmitTransaction("DeleteFAQ", string(outBytes))
		if err != nil {
			nethttp.ServerResponse(ctx, http.StatusInternalServerError, nethttp.ServerIssue, err)
			return
		}

		nethttp.ServerResponse(ctx, http.StatusOK, nethttp.Success, string(resp))
	}
}
