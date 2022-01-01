package profile

import (
	"encoding/json"

	"net/http"

	"github.com/anil8753/onesheds/apps/warehousemen/service/auth"
	"github.com/anil8753/onesheds/apps/warehousemen/service/ledger"
	"github.com/anil8753/onesheds/apps/warehousemen/service/nethttp"
	"github.com/anil8753/onesheds/apps/warehousemen/service/token"
	"github.com/gin-gonic/gin"
)

func (s *Profile) GetProfileHandler() gin.HandlerFunc {
	return func(c *gin.Context) {

		u, err := token.ExtractUserData(c)
		if err != nil {
			c.JSON(
				http.StatusUnauthorized,
				nethttp.NewHttpResponseWithMsg(nethttp.UserNotAuthorized, err.Error()),
			)
			return
		}

		iud, err := s.Dep.GetDB().Get(u.User)
		if err != nil {
			c.JSON(
				http.StatusBadRequest,
				nethttp.NewHttpResponseWithMsg(nethttp.UserNotExist, err.Error()),
			)
			return
		}

		b, err := json.Marshal(iud)
		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				nethttp.NewHttpResponseWithMsg(nethttp.ServerIssue, err.Error()),
			)
			return
		}

		var udata auth.UserData
		if err := json.Unmarshal(b, &udata); err != nil {
			c.JSON(
				http.StatusInternalServerError,
				nethttp.NewHttpResponseWithMsg(nethttp.ServerIssue, err.Error()),
			)
			return
		}

		resp, err := s.executeLedger(udata.Crypto)
		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				nethttp.NewHttpResponseWithMsg(nethttp.ServerIssue, err.Error()),
			)
			return
		}

		c.JSON(
			http.StatusOK,
			nethttp.NewHttpResponseWithMsg(nethttp.Sucess, string(resp)),
		)
	}

}

func (s *Profile) executeLedger(crypt *ledger.UserCrpto) ([]byte, error) {
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

	data, err := contract.EvaluateTransaction("GetIdentity")
	if err != nil {
		return nil, err
	}

	return data, nil

}
