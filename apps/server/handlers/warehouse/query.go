package warehouse

import (
	"net/http"

	"github.com/anil8753/onesheds/handlers/utils"
	"github.com/anil8753/onesheds/nethttp"
	"github.com/gin-gonic/gin"
)

func (s *Warehouse) QueryHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		q := utils.GetQueryStringFromContext(ctx)
		if q == "" {
			nethttp.ServerResponse(ctx, http.StatusBadRequest, nethttp.InvalidRequestData, "query string is missing")
			return
		}

		udata := utils.GetUserFromContext(ctx, s.Database)
		if udata == nil {
			nethttp.ServerResponse(ctx, http.StatusBadRequest, nethttp.InvalidRequestData, "user not found")
			return
		}

		contract, err := s.Ledger.GetUserContract(udata.Crypto)
		if err != nil {
			nethttp.ServerResponse(ctx, http.StatusInternalServerError, nethttp.ServerIssue, err)
			return
		}

		resp, err := contract.EvaluateTransaction("GetWarehousesByRichQuery", q)
		if err != nil {
			nethttp.ServerResponse(ctx, http.StatusInternalServerError, nethttp.ServerIssue, err)
			return
		}

		nethttp.ServerResponse(ctx, http.StatusOK, nethttp.Success, string(resp))
	}
}
