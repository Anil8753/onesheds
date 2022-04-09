package handlers

import (
	"log"
	"net/http"

	"github.com/anil8753/fabric-ca-client/utils"
	"github.com/gin-gonic/gin"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
)

func AllUsersHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		nodeType := ctx.Param("node_type")
		log.Println("nodeType: ", nodeType)

		resp, err := GetAllIdentities(nodeType)
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

func GetAllIdentities(nodeType string) ([]*msp.IdentityResponse, error) {

	mspClient, err := GetMSPClient(nodeType)
	if err != nil {
		return nil, err
	}

	return mspClient.GetAllIdentities()
}
