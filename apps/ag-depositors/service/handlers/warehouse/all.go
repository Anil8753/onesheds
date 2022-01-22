package warehouse

import (
	"net/http"

	"github.com/anil8753/onesheds/apps/warehousemen/service/handlers/utils"
	"github.com/anil8753/onesheds/apps/warehousemen/service/nethttp"
	"github.com/gin-gonic/gin"
)

func (s *Warehouse) QuaryAllHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		udata := utils.GetUserFromContext(ctx, s.Database)
		if udata == nil {
			return
		}

		contract, err := s.Ledger.GetUserContract(udata.Crypto)
		if err != nil {
			ctx.JSON(
				http.StatusInternalServerError,
				nethttp.NewHttpResponseWithMsg(nethttp.ServerIssue, err.Error()),
			)
			return
		}

		resp, err := contract.EvaluateTransaction("GetAllWarehouse")
		if err != nil {
			ctx.JSON(
				http.StatusInternalServerError,
				nethttp.NewHttpResponseWithMsg(nethttp.ServerIssue, err.Error()),
			)
			return
		}

		ctx.JSON(
			http.StatusOK,
			nethttp.NewHttpResponseWithMsg(nethttp.Sucess, string(resp)),
		)
	}
}
