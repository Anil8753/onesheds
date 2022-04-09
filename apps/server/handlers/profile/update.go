package profile

import (
	"encoding/json"
	"net/http"

	"github.com/anil8753/onesheds/handlers/utils"
	"github.com/anil8753/onesheds/ledger"
	"github.com/anil8753/onesheds/nethttp"
	"github.com/gin-gonic/gin"
)

func (s *Profile) UpdateProfileHandler() gin.HandlerFunc {
	//
	return func(ctx *gin.Context) {

		var urd ledger.RegisterationData
		if err := ctx.BindJSON(&urd); err != nil {
			nethttp.ServerResponse(ctx, http.StatusBadRequest, nethttp.InvalidRequestData, err)
			return
		}

		inBytes, err := json.Marshal(urd)
		if err != nil {
			nethttp.ServerResponse(ctx, http.StatusInternalServerError, nethttp.ServerIssue, err)
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

		resp, err := contract.SubmitTransaction("UpdateWarehouseUser", string(inBytes))
		if err != nil {
			nethttp.ServerResponse(ctx, http.StatusInternalServerError, nethttp.ServerIssue, err)
			return
		}

		nethttp.ServerResponse(ctx, http.StatusOK, nethttp.Success, string(resp))
	}
}
