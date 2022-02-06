package auth

import (
	"encoding/json"
	"net/http"

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
			nethttp.ServerResponse(ctx, http.StatusBadRequest, nethttp.InvalidRequestData, err)
			return
		}

		u, err := s.getUserData(reqData.User)
		if err != nil {
			nethttp.ServerResponse(ctx, http.StatusBadRequest, nethttp.UserNotExist, err)
			return
		}

		// check password
		err = s.loginCheck(u.Password, reqData.Password)
		if err != nil {
			nethttp.ServerResponse(ctx, http.StatusUnauthorized, nethttp.WrongCredentials, u.Password)
			return
		}

		// generate jwt token
		tokenPair, err := token.GenerateTokenPair(&token.UserData{User: u.User, UserId: u.UserId})
		if err != nil {
			nethttp.ServerResponse(ctx, http.StatusInternalServerError, nethttp.WrongCredentials, err)
			return
		}

		nethttp.ServerResponse(ctx, http.StatusOK, nethttp.Sucess, tokenPair)
	}
}

func (s *Auth) loginCheck(hashedPassword string, password string) error {

	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return err
	}

	return nil
}

func (s *Auth) getUserData(user string) (*UserData, error) {

	iud, err := s.Database.Get(user)
	if err != nil {
		return nil, err
	}

	byt, err := json.Marshal(iud)
	if err != nil {
		return nil, err
	}

	var u UserData
	if err := json.Unmarshal(byt, &u); err != nil {
		return nil, err
	}

	return &u, nil
}
