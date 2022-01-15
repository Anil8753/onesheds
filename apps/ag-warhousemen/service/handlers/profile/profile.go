package profile

import (
	"net/http"

	"github.com/anil8753/onesheds/apps/warehousemen/service/handlers/auth"
	"github.com/anil8753/onesheds/apps/warehousemen/service/ledger"
	"github.com/anil8753/onesheds/apps/warehousemen/service/nethttp"
	"github.com/gin-gonic/gin"
)

func (s *Profile) GetProfileHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		udata := auth.GetUserFromSession(ctx, s.Dep.GetDB())
		if udata == nil {
			return
		}

		resp, err := s.Dep.GetLedger().GetWarehouseUser(udata.Crypto)
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

func (s *Profile) UpdateProfileHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		udata := auth.GetUserFromSession(ctx, s.Dep.GetDB())
		if udata == nil {
			return
		}

		var urd ledger.RegisterationData
		if err := ctx.BindJSON(&urd); err != nil {
			ctx.JSON(
				http.StatusBadRequest,
				nethttp.NewHttpResponseWithMsg(nethttp.InvalidRequestData, err.Error()),
			)
			return
		}

		resp, err := s.Dep.GetLedger().UpdateWarehouseUser(udata.Crypto, &urd)
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
