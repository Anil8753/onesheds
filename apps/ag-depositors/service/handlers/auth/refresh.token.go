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
			nethttp.ServerResponse(ctx, http.StatusBadRequest, nethttp.InvalidRequestData, err)
			return
		}

		tokenPair, err := s.refreshTokens(reqData.RereshToken)
		if err != nil {
			nethttp.ServerResponse(ctx, http.StatusUnauthorized, nethttp.UserNotAuthorized, err)
			return
		}

		nethttp.ServerResponse(ctx, http.StatusOK, nethttp.Sucess, tokenPair)
	}
}

func (s *Auth) refreshTokens(tokenStr string) (*token.TokenPair, error) {

	user, err := token.GetUserFromRefreshToken(tokenStr)
	if err != nil {
		return nil, err
	}

	u, err := s.getUserData(user)
	if err != nil {
		return nil, err
	}

	tokenPair, err := token.GenerateTokenPair(&token.UserData{User: u.User, UserId: u.UserId})
	if err != nil {
		return nil, err
	}

	return tokenPair, nil
}
