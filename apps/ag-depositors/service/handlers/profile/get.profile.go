package profile

import (
	"encoding/json"

	"net/http"

	"github.com/anil8753/onesheds/apps/warehousemen/service/handlers/auth"
	"github.com/anil8753/onesheds/apps/warehousemen/service/ledger"
	"github.com/anil8753/onesheds/apps/warehousemen/service/nethttp"
	"github.com/gin-gonic/gin"
)

func (s *Profile) GetProfileHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		user := ctx.GetString("user")

		iud, err := s.Database.Get(user)
		if err != nil {
			nethttp.ServerResponse(ctx, http.StatusBadRequest, nethttp.UserNotExist, err)
			return
		}

		b, err := json.Marshal(iud)
		if err != nil {
			nethttp.ServerResponse(ctx, http.StatusInternalServerError, nethttp.ServerIssue, err)
			return
		}

		var udata auth.UserData
		if err := json.Unmarshal(b, &udata); err != nil {
			nethttp.ServerResponse(ctx, http.StatusInternalServerError, nethttp.ServerIssue, err)
			return
		}

		resp, err := s.executeLedger(udata.Crypto, "GetIdentity")
		if err != nil {
			nethttp.ServerResponse(ctx, http.StatusInternalServerError, nethttp.ServerIssue, err)
			return
		}

		nethttp.ServerResponse(ctx, http.StatusOK, nethttp.Sucess, string(resp))
	}

}

func (s *Profile) executeLedger(crypt *ledger.UserCrpto, fn string) ([]byte, error) {
	le := ledger.Ledger{}
	le.Init()
	gw, err := le.GetGateway(crypt)
	if err != nil {
		return nil, err
	}

	contract, err := le.GetContract(gw, "mychannel", "core")
	if err != nil {
		return nil, err
	}

	data, err := contract.EvaluateTransaction(fn)
	if err != nil {
		return nil, err
	}

	return data, nil

}
