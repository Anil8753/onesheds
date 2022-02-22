package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/anil8753/fabric-ca-client/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
)

type UserRegistrationData struct {
	Email      string `json:"email" binding:"required"`
	NodeType   string `json:"nodeType" binding:"required"`
	Attributes []struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	} `json:"attributes"`
}

type UserIdentity struct {
	MSP        string `json:"msp"`
	UserId     string `json:"userId"`
	Cert       string `json:"cert"`
	PrivateKey string `json:"privateKey"`
}

func RegisterUserHandler() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		var urd UserRegistrationData
		if err := ctx.ShouldBindJSON(&urd); err != nil {
			ctx.JSON(
				http.StatusBadRequest,
				utils.HttpError("StatusBadRequest", err.Error()),
			)
			return
		}

		log.Println("UserRegistrationData", urd)

		ui, err := RegEnrollUser(&urd)
		if err != nil {
			ctx.JSON(
				http.StatusInternalServerError,
				utils.HttpError("StatusInternalServerError", err.Error()),
			)
			return
		}

		ctx.JSON(
			http.StatusOK,
			utils.HttpSucess("StatusOK", ui),
		)
	}
}

func RegEnrollUser(regData *UserRegistrationData) (*UserIdentity, error) {

	mspClient, err := GetMSPClient()
	if err != nil {
		return nil, err
	}

	attrs := []msp.Attribute{}
	for _, attr := range regData.Attributes {
		attrs = append(attrs, msp.Attribute{Name: attr.Key, Value: attr.Value, ECert: true})
	}

	attrs = append(attrs, msp.Attribute{Name: "email", Value: regData.Email, ECert: true})
	attrs = append(attrs, msp.Attribute{Name: "nodeType", Value: regData.NodeType, ECert: true})

	userId := fmt.Sprintf("user_%s", uuid.New().String())

	rr := &msp.RegistrationRequest{
		Name:       userId,
		Attributes: attrs,
	}

	log.Println("Register", rr)

	userpw, err := mspClient.Register(rr)
	if err != nil {
		return nil, fmt.Errorf("failed to register user: %s. %w", userId, err)
	}

	//Enroll with returned password
	err = mspClient.Enroll(userId, msp.WithSecret(userpw))
	if err != nil {
		return nil, fmt.Errorf("failed to enroll user %s. %w", userId, err)
	}

	return GetSigningIdentityWithMSPClient(mspClient, userId)
}
