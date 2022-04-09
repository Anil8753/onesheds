package utils

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/anil8753/onesheds/apps/warehousemen/service/nethttp"
	"github.com/gin-gonic/gin"
)

func GetQueryStringFromContext(ctx *gin.Context) string {

	q := ctx.Query("q")
	if q == "" {
		nethttp.ServerResponse(ctx, http.StatusBadRequest, nethttp.InvalidRequestData, "query string is missing")
		return ""
	}

	var qi interface{}
	if err := json.Unmarshal([]byte(q), &qi); err != nil {
		nethttp.ServerResponse(ctx, http.StatusBadRequest, nethttp.InvalidRequestData, err)
		return ""
	}

	return q
}

func GetPageSizeFromContext(ctx *gin.Context) int32 {

	pageSizeStr := ctx.Query("pagesize")
	if pageSizeStr == "" {
		err := errors.New("pagesize is missing")
		nethttp.ServerResponse(ctx, http.StatusBadRequest, nethttp.InvalidRequestData, err)
		return 0
	}

	ps, err := strconv.Atoi(pageSizeStr)

	if err != nil {
		err := errors.New("pagesize is not a number")
		nethttp.ServerResponse(ctx, http.StatusBadRequest, nethttp.InvalidRequestData, err)
		return 0
	}

	if ps < 1 {
		err := errors.New("pagesize imust be > 0")
		nethttp.ServerResponse(ctx, http.StatusBadRequest, nethttp.InvalidRequestData, err)
		return 0
	}

	return int32(ps)
}

func GetPaginationBookmarkFromContext(ctx *gin.Context) string {

	bookmark := ctx.Query("bookmark")
	return bookmark
}
