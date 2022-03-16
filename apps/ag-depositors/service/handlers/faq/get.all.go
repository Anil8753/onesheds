package faq

import (
	"net/http"

	"github.com/anil8753/onesheds/apps/warehousemen/service/handlers/utils"
	"github.com/anil8753/onesheds/apps/warehousemen/service/nethttp"
	"github.com/gin-gonic/gin"
)

func (s *Handler) GetAllFAQ() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		warehouse_id, found := ctx.Params.Get("warehouse_id")
		if !found || warehouse_id == "" {
			nethttp.ServerResponse(ctx, http.StatusBadRequest, nethttp.InvalidRequestData, "warehouse_id is missing")
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

		resp, err := contract.EvaluateTransaction("GetAllFAQ", udata.UserId)
		if err != nil {
			nethttp.ServerResponse(ctx, http.StatusInternalServerError, nethttp.ServerIssue, err)
			return
		}

		nethttp.ServerResponse(ctx, http.StatusOK, nethttp.Success, string(resp))
	}
}
