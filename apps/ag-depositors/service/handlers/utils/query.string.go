package utils

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/anil8753/onesheds/apps/warehousemen/service/nethttp"
	"github.com/gin-gonic/gin"
)

func GetQueryStringFromContext(ctx *gin.Context) string {

	q := ctx.Query("q")
	if q == "" {
		ctx.JSON(
			http.StatusBadRequest,
			nethttp.NewHttpResponseWithMsg(nethttp.InvalidRequestData, "query string is missing"),
		)
		return ""
	}

	var qi interface{}
	if err := json.Unmarshal([]byte(q), &qi); err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			nethttp.NewHttpResponseWithMsg(nethttp.InvalidRequestData, "query string is in bad format"),
		)
		return ""
	}

	return q
}

func GetPageSizeFromContext(ctx *gin.Context) int32 {

	pageSizeStr := ctx.Query("pagesize")
	if pageSizeStr == "" {
		ctx.JSON(
			http.StatusBadRequest,
			nethttp.NewHttpResponseWithMsg(nethttp.InvalidRequestData, "pagesize is missing"),
		)
		return 0
	}

	ps, err := strconv.Atoi(pageSizeStr)

	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			nethttp.NewHttpResponseWithMsg(nethttp.InvalidRequestData, "pagesize is not a number"),
		)
		return 0
	}

	if ps < 1 {
		ctx.JSON(
			http.StatusBadRequest,
			nethttp.NewHttpResponseWithMsg(nethttp.InvalidRequestData, "pagesize imust be > 0"),
		)
		return 0
	}

	return int32(ps)
}

func GetPaginationBookmarkFromContext(ctx *gin.Context) string {

	bookmark := ctx.Query("bookmark")
	return bookmark
}
