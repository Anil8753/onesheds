package profile

import (
	"net/http"

	"github.com/anil8753/onesheds/apps/warehousemen/service/handlers/auth"
	"github.com/anil8753/onesheds/apps/warehousemen/service/nethttp"
	"github.com/gin-gonic/gin"
)

func (s *Profile) GetIdentityHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		udata := auth.GetUserFromSession(ctx, s.Dep.GetDB())
		if udata == nil {
			return
		}

		resp, err := s.Dep.GetLedger().GetUserIdentity(udata.Crypto)
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
