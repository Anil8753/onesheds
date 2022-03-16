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
			nethttp.ServerResponse(ctx, http.StatusBadRequest, nethttp.InvalidRequestData, "query string is missing")
			return
		}

		pageSize := utils.GetPageSizeFromContext(ctx)
		if pageSize == 0 {
			nethttp.ServerResponse(ctx, http.StatusBadRequest, nethttp.InvalidRequestData, "page size is missing")
			return
		}

		bookmark := utils.GetPaginationBookmarkFromContext(ctx)

		udata := utils.GetUserFromContext(ctx, s.Database)
		if udata == nil {
			nethttp.ServerResponse(ctx, http.StatusBadRequest, nethttp.InvalidRequestData, "user not found")
			return
		}

		contract, err := s.Ledger.GetUserContract(udata.Crypto)
		if err != nil {
			nethttp.ServerResponse(ctx, http.StatusInternalServerError, nethttp.ServerIssue, err)
			return
		}

		resp, err := contract.EvaluateTransaction("GetWarehousesByRichQueryWithPagination", q, fmt.Sprintf("%d", pageSize), bookmark)
		if err != nil {
			nethttp.ServerResponse(ctx, http.StatusInternalServerError, nethttp.ServerIssue, err)
			return
		}

		nethttp.ServerResponse(ctx, http.StatusOK, nethttp.Success, string(resp))
	}
}
