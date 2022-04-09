package nethttp

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HttpResponse struct {
	Code     int         `json:"code"`
	CodeDesc string      `json:"codeDesc"`
	Data     interface{} `json:"data"`
}

const (
	Success = iota
	WrongCredentials
	UserAlreadyExist
	UserNotExist
	UserNotAuthorized
	InvalidRequestData
	ServerIssue
)

var mappings map[int]string

func init() {

	mappings = map[int]string{
		Success:                        "success",
		WrongCredentials:               "wrong credentials",
		UserAlreadyExist:               "user already exist",
		UserNotExist:                   "user not exist",
		UserNotAuthorized:              "user not authorized",
		InvalidRequestData:             "invalid request data",
		http.StatusInternalServerError: "server issue",
	}
}

func ServerResponse(ctx *gin.Context, httpCode int, code int, data interface{}) {

	codedesc := mappings[code]
	r := &HttpResponse{Code: code, CodeDesc: codedesc, Data: data}

	ctx.JSON(
		httpCode,
		r,
	)
}
