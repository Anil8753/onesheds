package order

import (
	"net/http"

	"github.com/anil8753/onesheds/handlers/utils"
	"github.com/anil8753/onesheds/nethttp"
	"github.com/gin-gonic/gin"
)

func (s *Handler) GetAllOrders() gin.HandlerFunc {
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

		resp, err := contract.EvaluateTransaction("QueryDepositorAllOrders", udata.UserId)
		if err != nil {
			nethttp.ServerResponse(ctx, http.StatusInternalServerError, nethttp.ServerIssue, err)
			return
		}

		nethttp.ServerResponse(ctx, http.StatusOK, nethttp.Success, string(resp))
	}
}
