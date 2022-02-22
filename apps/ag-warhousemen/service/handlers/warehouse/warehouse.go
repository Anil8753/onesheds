package warehouse

import (
	"errors"
	"net/http"

	"github.com/anil8753/onesheds/apps/warehousemen/service/handlers/auth"
	"github.com/anil8753/onesheds/apps/warehousemen/service/ledger"
	"github.com/anil8753/onesheds/apps/warehousemen/service/nethttp"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (s *Asset) GetWarehousesHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		udata := auth.GetUserFromSession(ctx, s.Dep.GetDB())
		if udata == nil {
			return
		}

		resp, err := s.Dep.GetLedger().GetUserWarehouses(udata.Crypto)
		if err != nil {
			ctx.JSON(
				http.StatusInternalServerError,
				nethttp.NewHttpResponseWithMsg(nethttp.ServerIssue, err.Error()),
			)
			return
		}

		ctx.JSON(
			http.StatusOK,
			nethttp.NewHttpResponseWithMsg(nethttp.Success, string(resp)),
		)
	}
}

func (s *Asset) CreateWarehouseHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		udata := auth.GetUserFromSession(ctx, s.Dep.GetDB())
		if udata == nil {
			return
		}

		var data ledger.AssetData
		if err := ctx.BindJSON(&data); err != nil {
			ctx.JSON(
				http.StatusBadRequest,
				nethttp.NewHttpResponseWithMsg(nethttp.InvalidRequestData, err.Error()),
			)
			return
		}

		data.WarehouseId = uuid.New().String()
		data.OwnerId = udata.UserId

		resp, err := s.Dep.GetLedger().RegisterWarehouse(udata.Crypto, &data)
		if err != nil {
			ctx.JSON(
				http.StatusInternalServerError,
				nethttp.NewHttpResponseWithMsg(nethttp.ServerIssue, err.Error()),
			)
			return
		}

		ctx.JSON(
			http.StatusOK,
			nethttp.NewHttpResponseWithMsg(nethttp.Success, string(resp)),
		)
	}
}

func (s *Asset) UpdateWarehouseHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		udata := auth.GetUserFromSession(ctx, s.Dep.GetDB())
		if udata == nil {
			return
		}

		var data ledger.AssetData
		if err := ctx.BindJSON(&data); err != nil {
			ctx.JSON(
				http.StatusBadRequest,
				nethttp.NewHttpResponseWithMsg(nethttp.InvalidRequestData, err.Error()),
			)
			return
		}

		resp, err := s.Dep.GetLedger().UpdateWarehouse(udata.Crypto, &data)
		if err != nil {

			msg := err.Error()
			currentErr := err

			for errors.Unwrap(currentErr) != nil {
				currentErr = errors.Unwrap(currentErr)
				msg = msg + currentErr.Error()
			}

			ctx.JSON(
				http.StatusInternalServerError,
				nethttp.NewHttpResponseWithMsg(nethttp.ServerIssue, msg),
			)
			return
		}

		ctx.JSON(
			http.StatusOK,
			nethttp.NewHttpResponseWithMsg(nethttp.Success, string(resp)),
		)
	}
}
