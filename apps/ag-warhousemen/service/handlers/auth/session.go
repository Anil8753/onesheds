package auth

import (
	"fmt"
	"net/http"

	"github.com/anil8753/onesheds/apps/warehousemen/service/interfaces"
	"github.com/anil8753/onesheds/apps/warehousemen/service/ledger"
	"github.com/anil8753/onesheds/apps/warehousemen/service/nethttp"
	"github.com/gin-gonic/gin"
)

type SessionData struct {
	UserId string
	Crypto *ledger.UserCrpto
}

var usersCrypto map[string]*ledger.UserCrpto

func GetUserFromSession(ctx *gin.Context, db interfaces.Database) *SessionData {

	userId := ctx.GetString("userId")

	crypto, err := RetriveCrypto(userId)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			nethttp.NewHttpResponseWithMsg(nethttp.ServerIssue, err.Error()),
		)
		return nil
	}

	return &SessionData{UserId: userId, Crypto: crypto}
}

func StoreCrypto(userId string, crypto *ledger.UserCrpto) {

	if usersCrypto == nil {
		usersCrypto = make(map[string]*ledger.UserCrpto)
	}

	usersCrypto[userId] = crypto
}

func RetriveCrypto(userId string) (*ledger.UserCrpto, error) {

	crypto, ok := usersCrypto[userId]
	if !ok {
		return nil, fmt.Errorf("failed to get crypto for user: %s", userId)
	}

	return crypto, nil
}
