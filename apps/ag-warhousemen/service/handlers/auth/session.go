package auth

import (
	"encoding/json"
	"net/http"

	"github.com/anil8753/onesheds/apps/warehousemen/service/interfaces"
	"github.com/anil8753/onesheds/apps/warehousemen/service/nethttp"
	"github.com/gin-gonic/gin"
)

func GetUserFromSession(ctx *gin.Context, db interfaces.Database) *UserData {

	user := ctx.GetString("user")

	iud, err := db.Get(user)
	if err != nil {
		nethttp.ServerResponse(ctx, http.StatusBadRequest, nethttp.UserNotExist, err)
		return nil
	}

	b, err := json.Marshal(iud)
	if err != nil {
		nethttp.ServerResponse(ctx, http.StatusInternalServerError, nethttp.ServerIssue, err)
		return nil
	}

	var udata UserData
	if err := json.Unmarshal(b, &udata); err != nil {
		nethttp.ServerResponse(ctx, http.StatusInternalServerError, nethttp.ServerIssue, err)
		return nil
	}

	return &udata
}
