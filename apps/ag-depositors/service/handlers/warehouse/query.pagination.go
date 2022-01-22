package warehouse

import (
	"fmt"
	"net/http"

	"github.com/anil8753/onesheds/apps/warehousemen/service/handlers/utils"
	"github.com/anil8753/onesheds/apps/warehousemen/service/nethttp"
	"github.com/gin-gonic/gin"
)

func (s *Warehouse) QueryPaginationHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		q := utils.GetQueryStringFromContext(ctx)
		if q == "" {
			return
		}

		pageSize := utils.GetPageSizeFromContext(ctx)
		if pageSize == 0 {
			return
		}

		bookmark := utils.GetPaginationBookmarkFromContext(ctx)

		udata := utils.GetUserFromContext(ctx, s.Database)
		if udata == nil {
			return
		}

		contract, err := s.Ledger.GetUserContract(udata.Crypto)
		if err != nil {
			ctx.JSON(
				http.StatusInternalServerError,
				nethttp.NewHttpResponseWithMsg(nethttp.ServerIssue, err.Error()),
			)
			return
		}

		resp, err := contract.EvaluateTransaction("GetWarehousesByRichQueryWithPagination", q, fmt.Sprintf("%d", pageSize), bookmark)
		if err != nil {
			ctx.JSON(
				http.StatusInternalServerError,
				nethttp.NewHttpResponseWithMsg(nethttp.ServerIssue, err.Error()),
			)
			return
		}

		ctx.JSON(
			http.StatusOK,
			nethttp.NewHttpResponseWithMsg(nethttp.Sucess, string(resp)),
		)
	}
}
