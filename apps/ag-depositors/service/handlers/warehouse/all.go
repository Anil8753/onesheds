package warehouse

import (
	"encoding/json"
	"net/http"

	"github.com/anil8753/onesheds/apps/warehousemen/service/auth"
	"github.com/anil8753/onesheds/apps/warehousemen/service/nethttp"
	"github.com/gin-gonic/gin"
)

func (s *Warehouse) GetWarehouseHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		user := ctx.GetString("user")

		iud, err := s.Database.Get(user)
		if err != nil {
			ctx.JSON(
				http.StatusBadRequest,
				nethttp.NewHttpResponseWithMsg(nethttp.UserNotExist, err.Error()),
			)
			return
		}

		b, err := json.Marshal(iud)
		if err != nil {
			ctx.JSON(
				http.StatusInternalServerError,
				nethttp.NewHttpResponseWithMsg(nethttp.ServerIssue, err.Error()),
			)
			return
		}

		var udata auth.UserData
		if err := json.Unmarshal(b, &udata); err != nil {
			ctx.JSON(
				http.StatusInternalServerError,
				nethttp.NewHttpResponseWithMsg(nethttp.ServerIssue, err.Error()),
			)
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
