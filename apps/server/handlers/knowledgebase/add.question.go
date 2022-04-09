package knowledgebase

import (
	"fmt"
	"net/http"

	"github.com/anil8753/onesheds/handlers/utils"
	"github.com/anil8753/onesheds/nethttp"
	"github.com/gin-gonic/gin"
)

func (s *Handler) AddQuestion() gin.HandlerFunc {
	//
	return func(ctx *gin.Context) {

		type PostData struct {
			WarehouseId string `json:"warehouseId" binding:"required"`
			Question    string `json:"question" binding:"required"`
		}

		var postData PostData
		if err := ctx.ShouldBindJSON(&postData); err != nil {
			nethttp.ServerResponse(ctx, http.StatusBadRequest, nethttp.InvalidRequestData, err)
			return
		}

		fmt.Println(postData)

		udata := utils.GetUserFromContext(ctx, s.Database)
		if udata == nil {
			return
		}

		contract, err := s.Ledger.GetUserContract(udata.Crypto)
		if err != nil {
			nethttp.ServerResponse(ctx, http.StatusInternalServerError, nethttp.ServerIssue, err)
			return
		}

		resp, err := contract.SubmitTransaction("AddWarehouseQuestion", postData.WarehouseId, postData.Question)
		if err != nil {
			nethttp.ServerResponse(ctx, http.StatusInternalServerError, nethttp.ServerIssue, err)
			return
		}

		nethttp.ServerResponse(ctx, http.StatusOK, nethttp.Success, string(resp))
	}
}
