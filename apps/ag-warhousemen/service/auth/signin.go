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

type SigninResp struct {
	AccessToken string `json:"access_token"`
}

func (s *Auth) SigninHandler() gin.HandlerFunc {
	return func(c *gin.Context) {

		var reqData SigninReq
		if err := c.ShouldBindJSON(&reqData); err != nil {
			c.JSON(
				http.StatusBadRequest,
				nethttp.NewHttpResponseWithMsg(nethttp.InvalidRequestData, err.Error()),
			)
			return
		}

		u, err := s.getUserData(reqData.User)
		if err != nil {
			c.JSON(
				http.StatusBadRequest,
				nethttp.NewHttpResponseWithMsg(nethttp.UserNotExist, err.Error()),
			)
			return
		}

		// check password
		err = s.loginCheck(u.Password, reqData.Password)
		if err != nil {
			c.JSON(
				http.StatusUnauthorized,
				nethttp.NewHttpResponse(nethttp.WrongCredentials),
			)
			return
		}

		// generate jwt token
		token, err := token.GenerateToken(u.UserUniqueId)
		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				nethttp.NewHttpResponseWithMsg(nethttp.WrongCredentials, err.Error()),
			)
			return
		}

		c.JSON(
			http.StatusOK,
			nethttp.NewHttpResponseWithMsg(nethttp.Sucess, SigninResp{AccessToken: token}),
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

func (s *Auth) getUserData(user string) (*UserData, error) {

	iud, err := s.Dep.GetDB().Get(user)
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
