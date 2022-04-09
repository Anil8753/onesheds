package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/anil8753/fabric-ca-client/utils"
	"github.com/gin-gonic/gin"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
)

// Binding from JSON
type RevokeReq struct {
	UserId string `json:"userId" binding:"required"`
	Reason string ` json:"reason" binding:"required"`
}

func RevokeUserHandler() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		nodeType := ctx.Param("node_type")
		log.Println("nodeType: ", nodeType)

		var r RevokeReq
		if err := ctx.ShouldBindJSON(&r); err != nil {
			ctx.JSON(
				http.StatusBadRequest,
				utils.HttpError("StatusBadRequest", err.Error()),
			)
			return
		}

		mspClient, err := GetMSPClient(nodeType)
		if err != nil {
			ctx.JSON(
				http.StatusInternalServerError,
				utils.HttpError("StatusInternalServerError", err.Error()),
			)
			return
		}

		if _, err := GetSigningIdentityWithMSPClient(mspClient, r.UserId); err != nil {
			ctx.JSON(
				http.StatusNotFound,
				utils.HttpError("StatusNotFound", err.Error()),
			)
			return
		}

		resp, err := RevokeUser(mspClient, r.UserId, r.Reason)

		if err != nil {
			ctx.JSON(
				http.StatusInternalServerError,
				utils.HttpError("StatusInternalServerError", err.Error()),
			)
			return
		}

		ctx.JSON(
			http.StatusOK,
			utils.HttpSucess("StatusOK", resp),
		)
	}
}

func RevokeUser(mspClient *msp.Client, userId string, reason string) (*msp.RevocationResponse, error) {

	request := &msp.RevocationRequest{
		Name:   userId,
		Reason: reason,
	}

	resp, err := mspClient.Revoke(request)
	if err != nil {
		return nil, fmt.Errorf("user %s revoke failed. %w", userId, err)
	}

	return resp, nil
}
