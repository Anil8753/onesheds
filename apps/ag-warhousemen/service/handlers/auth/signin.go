package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/anil8753/onesheds/apps/warehousemen/service/ledger"
	"github.com/anil8753/onesheds/apps/warehousemen/service/nethttp"
	"github.com/anil8753/onesheds/apps/warehousemen/service/token"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Binding from JSON
type SigninReq struct {
	User     string `json:"user" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (s *Auth) SigninHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var reqData SigninReq
		if err := ctx.ShouldBindJSON(&reqData); err != nil {
			ctx.JSON(
				http.StatusBadRequest,
				nethttp.NewHttpResponseWithMsg(nethttp.InvalidRequestData, err.Error()),
			)
			return
		}

		creds, err := s.getCredentials(reqData.User)
		if err != nil {
			ctx.JSON(
				http.StatusBadRequest,
				nethttp.NewHttpResponseWithMsg(nethttp.UserNotExist, err.Error()),
			)
			return
		}

		// check password
		err = s.loginCheck(creds.Password, reqData.Password)
		if err != nil {
			ctx.JSON(
				http.StatusUnauthorized,
				nethttp.NewHttpResponse(nethttp.WrongCredentials),
			)
			return
		}

		// Gets the public and private key from the cert service
		if err := s.retriveUserCrypto(creds.UserId); err != nil {
			ctx.JSON(
				http.StatusInternalServerError,
				nethttp.NewHttpResponseWithMsg(nethttp.CryptoRetriveFailed, err.Error()),
			)
			return
		}

		// generate jwt token
		tokenPair, err := token.GenerateTokenPair(&token.UserData{User: creds.User, UserId: creds.UserId})
		if err != nil {
			ctx.JSON(
				http.StatusInternalServerError,
				nethttp.NewHttpResponseWithMsg(nethttp.WrongCredentials, err.Error()),
			)
			return
		}

		ctx.JSON(
			http.StatusOK,
			nethttp.NewHttpResponseWithMsg(nethttp.Success, tokenPair),
		)
	}
}

func (s *Auth) loginCheck(hashedPassword string, password string) error {

	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return err
	}

	return nil
}

func (s *Auth) getCredentials(user string) (*Credentials, error) {

	iud, err := s.Dep.GetDB().Get(user)
	if err != nil {
		return nil, err
	}

	byt, err := json.Marshal(iud)
	if err != nil {
		return nil, err
	}

	var creds Credentials
	if err := json.Unmarshal(byt, &creds); err != nil {
		return nil, err
	}

	return &creds, nil
}

func (s *Auth) retriveUserCrypto(userId string) error {

	url := fmt.Sprintf("%s/api/v1/users/%s", os.Getenv("IDENTITY_SERVICE_ENDPOINT"), userId)
	resp, err := http.Get(url)

	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed with http status code: %d", resp.StatusCode)
	}

	var certServiceResp struct{ Data ledger.UserCrpto }
	if err := json.NewDecoder(resp.Body).Decode(&certServiceResp); err != nil {
		return err
	}

	// Store public and private key in memory
	StoreCrypto(userId, &certServiceResp.Data)

	return nil
}
