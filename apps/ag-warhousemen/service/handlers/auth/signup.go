package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/anil8753/onesheds/apps/warehousemen/service/ledger"
	"github.com/anil8753/onesheds/apps/warehousemen/service/nethttp"
)

// Binding from JSON
type SignupReq struct {
	User       string      `json:"user" binding:"required"`
	Password   string      `json:"password" binding:"required"`
	Attributes []Attribute `json:"attributes"`
}

type SignupResp struct {
	UserId string `json:"userId" binding:"required"`
	User   string `json:"user" binding:"required"`
}

func (s *Auth) SignupHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		reqData := SignupReq{}
		if err := ctx.ShouldBindJSON(&reqData); err != nil {
			ctx.JSON(
				http.StatusBadRequest,
				nethttp.NewHttpResponseWithMsg(nethttp.InvalidRequestData, err.Error()),
			)
			return
		}

		if _, err := s.Dep.GetDB().Get(reqData.User); err == nil {
			ctx.JSON(
				http.StatusConflict,
				nethttp.NewHttpResponse(nethttp.UserAlreadyExist),
			)
			return
		}

		resp, err := s.doSignup(ctx, &reqData)
		if err != nil {
			ctx.JSON(
				http.StatusInternalServerError,
				nethttp.NewHttpResponseWithMsg(nethttp.ServerIssue, err.Error()),
			)
			return
		}

		ctx.JSON(http.StatusOK, resp)
	}
}

func (s *Auth) doSignup(ctx *gin.Context, reqData *SignupReq) (*SignupResp, error) {

	user := reqData.User

	ucert, err := s.createUserCert(ctx, reqData)
	if err != nil {
		return nil, err
	}

	// register user on ledger
	r := &ledger.RegisterationData{}
	r.UserId = ucert.UserId
	r.Email = user

	if _, err := s.Dep.GetLedger().RegisterWarehouseUser(ucert, r); err != nil {
		return nil, err
	}

	// save credentials
	creds := Credentials{
		UserId:   ucert.UserId,
		User:     user,
		Password: reqData.Password,
	}

	if _, err = creds.SaveUser(s.Dep.GetDB()); err != nil {
		return nil, err
	}

	return &SignupResp{UserId: ucert.UserId, User: user}, nil
}

func (s *Auth) createUserCert(ctx *gin.Context, reqData *SignupReq) (*ledger.UserCrpto, error) {

	// prepare registration data
	urd := UserRegistrationData{}
	urd.User = reqData.User
	urd.NodeType = NodeType

	for _, kv := range reqData.Attributes {
		urd.Attributes = append(urd.Attributes, Attribute{Key: kv.Key, Value: kv.Value})
	}

	json_data, err := json.Marshal(urd)
	if err != nil {
		return nil, err
	}

	// post call
	url := fmt.Sprintf("%s/api/v1/registeruser", os.Getenv("IDENTITY_SERVICE_ENDPOINT"))
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(json_data))

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("failed with http status code: %d. error: %w", resp.StatusCode, err)
		}
		resp.Body.Close()

		return nil, fmt.Errorf("failed with http status code: %d. response: %s", resp.StatusCode, string(body))
	}

	var certServiceResp struct{ Data ledger.UserCrpto }
	if err := json.NewDecoder(resp.Body).Decode(&certServiceResp); err != nil {
		return nil, err
	}

	return &certServiceResp.Data, nil
}
