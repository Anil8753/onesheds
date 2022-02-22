package auth

import (
	"net/http"

	"github.com/anil8753/onesheds/apps/warehousemen/service/nethttp"
	"github.com/anil8753/onesheds/apps/warehousemen/service/token"
	"github.com/gin-gonic/gin"
)

// Binding from JSON
type RefreshTokenReq struct {
	RereshToken string `json:"refreshToken" binding:"required"`
}

func (s *Auth) RefreshTokenHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var reqData RefreshTokenReq
		if err := ctx.ShouldBindJSON(&reqData); err != nil {
			ctx.JSON(
				http.StatusBadRequest,
				nethttp.NewHttpResponseWithMsg(nethttp.InvalidRequestData, err.Error()),
			)
			return
		}

		tokenPair, err := s.refreshTokens(reqData.RereshToken)
		if err != nil {
			ctx.JSON(
				http.StatusUnauthorized,
				nethttp.NewHttpResponseWithMsg(nethttp.UserNotAuthorized, err.Error()),
			)
			return
		}

		ctx.JSON(
			http.StatusOK,
			nethttp.NewHttpResponseWithMsg(nethttp.Success, tokenPair),
		)
	}
}

func (s *Auth) refreshTokens(tokenStr string) (*token.TokenPair, error) {

	user, err := token.GetUserFromRefreshToken(tokenStr)
	if err != nil {
		return nil, err
	}

	creds, err := s.getCredentials(user)
	if err != nil {
		return nil, err
	}

	tokenPair, err := token.GenerateTokenPair(&token.UserData{User: creds.User, UserId: creds.UserId})
	if err != nil {
		return nil, err
	}

	return tokenPair, nil
}
