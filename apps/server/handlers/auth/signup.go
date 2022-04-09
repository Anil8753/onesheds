package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/anil8753/onesheds/ledger"
	"github.com/anil8753/onesheds/nethttp"
)

// Binding from JSON
type SignupReq struct {
	User     string `json:"user" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SignupResp struct {
	UserId string `json:"userId" binding:"required"`
	User   string `json:"user" binding:"required"`
}

func (s *Auth) SignupHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		reqData := SignupReq{}
		if err := ctx.ShouldBindJSON(&reqData); err != nil {
			nethttp.ServerResponse(ctx, http.StatusBadRequest, nethttp.InvalidRequestData, err)
			return
		}

		if _, err := s.Database.Get(reqData.User); err == nil {
			nethttp.ServerResponse(ctx, http.StatusBadRequest, nethttp.UserAlreadyExist, err)
			return
		}

		resp, err := s.doSignup(ctx, &reqData)
		if err != nil {
			nethttp.ServerResponse(ctx, http.StatusInternalServerError, nethttp.ServerIssue, err)
			return
		}

		ctx.JSON(http.StatusOK, resp)
	}
}

func (s *Auth) doSignup(ctx *gin.Context, reqData *SignupReq) (*SignupResp, error) {

	user := reqData.User

	ucert, err := s.createUserCert(ctx, user)
	if err != nil {
		return nil, err
	}

	u := UserData{
		UserId:   ucert.UserId,
		User:     user,
		Password: reqData.Password,
		Crypto:   ucert,
	}

	// save user on ledger
	if err := s.registerOnLedger(&u); err != nil {
		// ideally we should delete the crypto if ledger upadte is failed. But keeping it as todo item at present
		return nil, err
	}

	if _, err = u.SaveUser(s.Database); err != nil {
		return nil, err
	}

	resp := &SignupResp{UserId: u.UserId, User: user}
	return resp, nil
}

func (s *Auth) createUserCert(ctx *gin.Context, user string) (*ledger.UserCrpto, error) {
	//
	nodeType := os.Getenv("NODE_TYPE")
	// prepare registration data
	urd := UserRegistrationData{User: user, NodeType: nodeType}
	urd.Attributes = append(urd.Attributes, Attribute{Key: "user", Value: user})

	json_data, err := json.Marshal(urd)
	if err != nil {
		return nil, err
	}

	// post call
	url := fmt.Sprintf("%s/api/v1/registeruser/%s", os.Getenv("IDENTITY_SERVICE_ENDPOINT"), nodeType)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(json_data))
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed with http status code: %d", resp.StatusCode)
	}

	var out struct {
		Data   ledger.UserCrpto
		Status string
	}

	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, err
	}

	return &out.Data, nil
}

// func (s *Auth) createUserCert(ctx *gin.Context, user string) (*ledger.UserCrpto, error) {

// 	// prepare registration data
// 	urd := ca.UserRegistrationData{User: user, NodeType: os.Getenv("NODE_TYPE")}
// 	urd.Attributes = append(urd.Attributes, ca.Attribute{Key: "user", Value: user})

// 	ui, err := ca.RegEnrollUser(&urd)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to register user: %s. error: %w", user, err)
// 	}

// 	resp := &ledger.UserCrpto{
// 		MSP:        ui.MSP,
// 		UserId:     ui.UserId,
// 		Cert:       ui.Cert,
// 		PrivateKey: ui.PrivateKey,
// 	}

// 	return resp, nil
// }

func (s *Auth) registerOnLedger(u *UserData) error {

	r := &ledger.RegisterationData{}
	r.UserId = u.UserId
	r.Email = u.User

	if _, err := s.Ledger.CreateUser(u.Crypto, r); err != nil {
		return err
	}

	return nil
}
