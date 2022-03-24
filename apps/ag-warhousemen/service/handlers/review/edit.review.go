package review

import (
	"encoding/json"
	"net/http"

	"github.com/anil8753/onesheds/apps/warehousemen/service/handlers/utils"
	"github.com/anil8753/onesheds/apps/warehousemen/service/nethttp"
	"github.com/gin-gonic/gin"
)

func (s *Handler) EditUserReview() gin.HandlerFunc {
	//
	return func(ctx *gin.Context) {

		type EditReviewPostData struct {
			ReviewId   string  `json:"reviewId" binding:"required"`
			UserRating float32 `json:"userRating" binding:"required"`
			ReviewText string  `json:"reviewText" binding:"required"`
		}

		reqData := EditReviewPostData{}
		if err := ctx.ShouldBindJSON(&reqData); err != nil {
			nethttp.ServerResponse(ctx, http.StatusBadRequest, nethttp.InvalidRequestData, err)
			return
		}

		udata := utils.GetUserFromContext(ctx, s.Database)
		if udata == nil {
			return
		}

		contract, err := s.Ledger.GetUserContract(udata.Crypto)
		if err != nil {
			nethttp.ServerResponse(ctx, http.StatusInternalServerError, nethttp.ServerIssue, err)
			return
		}

		outBytes, err := json.Marshal(reqData)
		if err != nil {
			nethttp.ServerResponse(ctx, http.StatusInternalServerError, nethttp.ServerIssue, err)
			return
		}

		resp, err := contract.SubmitTransaction("EditUserReview", string(outBytes))
		if err != nil {
			nethttp.ServerResponse(ctx, http.StatusInternalServerError, nethttp.ServerIssue, err)
			return
		}

		nethttp.ServerResponse(ctx, http.StatusOK, nethttp.Success, string(resp))
	}
}
