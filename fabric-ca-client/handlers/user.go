package handlers

import (
	"fmt"
	"net/http"

	"github.com/anil8753/fabric-ca-client/utils"
	"github.com/gin-gonic/gin"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
)

func UserHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		userId := ctx.Param("id")
		resp, err := GetSigningIdentity(userId)

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

func GetSigningIdentity(userId string) (*UserIdentity, error) {

	mspClient, err := GetMSPClient()
	if err != nil {
		return nil, err
	}

	return GetSigningIdentityWithMSPClient(mspClient, userId)
}

func GetSigningIdentityWithMSPClient(mspClient *msp.Client, userId string) (*UserIdentity, error) {

	identity, err := mspClient.GetSigningIdentity(userId)

	if err != nil {
		return nil, fmt.Errorf("failed to get signing identity for user %s. %w", userId, err)
	}

	priKeyBytes, err := identity.PrivateKey().Bytes()
	if err != nil {
		return nil, fmt.Errorf("failed private key bytes for user %s. %w", userId, err)
	}

	certBytes := identity.EnrollmentCertificate()

	return &UserIdentity{
		MSP:        identity.Identifier().MSPID,
		UserId:     userId,
		Cert:       string(certBytes),
		PrivateKey: string(priKeyBytes),
	}, nil
}
