package handlers

import (
	"net/http"

	"github.com/anil8753/fabric-ca-client/utils"
	"github.com/gin-gonic/gin"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
)

func AllUsersHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		resp, err := GetAllIdentities()
		if err != nil {
			ctx.JSON(
				http.StatusInternalServerError,
				utils.HttpError("StatusInternalServerError", err.Error()),
			)
			return
		}

		ctx.JSON(http.StatusOK, utils.HttpSucess("StatusOK", resp))
	}
}

func GetAllIdentities() ([]*msp.IdentityResponse, error) {

	mspClient, err := GetMSPClient()
	if err != nil {
		return nil, err
	}

	return mspClient.GetAllIdentities()
}
