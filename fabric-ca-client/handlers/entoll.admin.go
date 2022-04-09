package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/anil8753/fabric-ca-client/utils"
	"github.com/gin-gonic/gin"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
)

func EnrollAdminHandler() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		nodeType := ctx.Param("node_type")
		log.Println("nodeType: ", nodeType)

		if err := EnrollAdmin(nodeType); err != nil {
			ctx.JSON(
				http.StatusBadRequest,
				utils.HttpError("StatusBadRequest", err.Error()),
			)
			return
		}

		ctx.JSON(
			http.StatusOK,
			utils.HttpSucess("StatusOK", "success"),
		)
	}
}

func EnrollAdmin(nodeType string) error {

	mspClient, caconfig, err := GetMSPClientWithCAConfig(nodeType)
	if err != nil {
		return err
	}

	// Now try to enroll the admin with its configured ID and password
	err = mspClient.Enroll(caconfig.Registrar.EnrollID, msp.WithSecret(caconfig.Registrar.EnrollSecret))
	if err != nil {
		return fmt.Errorf("failed to enroll the admin. %w", err)
	}

	return nil
}
