package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-uuid"

	"github.com/anil8753/onesheds/apps/warehousemen/service/ledger"
	"github.com/anil8753/onesheds/apps/warehousemen/service/nethttp"
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

	// create unique userId
	uid, err := uuid.GenerateUUID()
	if err != nil {
		return nil, err
	}

	userId := fmt.Sprintf("user_%s", uid)

	ucert, err := s.createUserCert(ctx, user, userId)
	if err != nil {
		return nil, err
	}

	u := UserData{
		UserId:   userId,
		User:     user,
		Password: reqData.Password,
		Crypto:   ucert,
	}

	// save user on ledger
	if err := s.doRegister(&u, reqData); err != nil {
		// ideally we should delete the crypto if ledger upadte is failed. But keeping it as todo item at present
		return nil, err
	}

	err = u.BeforeSave()
	if err != nil {
		return nil, err
	}
	_, err = u.SaveUser(s.Dep.GetDB())
	if err != nil {
		return nil, err
	}

	resp := &SignupResp{UserId: userId, User: user}
	return resp, nil
}

func (s *Auth) createUserCert(ctx *gin.Context, user string, userId string) (*ledger.UserCrpto, error) {

	// prepare registration data
	urd := UserRegistrationData{}
	urd.UserId = userId
	urd.Attributes = append(urd.Attributes, Attribute{Key: "userId", Value: userId})
	urd.Attributes = append(urd.Attributes, Attribute{Key: "user", Value: user})
	urd.Attributes = append(urd.Attributes, Attribute{Key: "nodetype", Value: NodeType})

	json_data, err := json.Marshal(urd)
	if err != nil {
		return nil, err
	}

	// post call
	url := fmt.Sprintf("%s/v1/createidentity", os.Getenv("IDENTITY_SERVICE_ENDPOINT"))
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(json_data))

	if err != nil {
		return nil, err
	}

	var uc ledger.UserCrpto
	if err := json.NewDecoder(resp.Body).Decode(&uc); err != nil {
		return nil, err
	}

	return &uc, nil
}

// func (s *Auth) createUserCert(ctx *gin.Context, user string, userId string) (*ledger.UserCrpto, error) {

// 	// create output file
// 	tempFile, err := uuid.GenerateUUID()
// 	if err != nil {
// 		return nil, err
// 	}

// 	outfile := path.Join(os.TempDir(), tempFile)
// 	defer os.Remove(outfile)

// 	// prepare registration data
// 	urd := UserRegistrationData{}
// 	urd.Attributes = append(urd.Attributes, Attribute{Key: "userId", Value: userId})
// 	urd.Attributes = append(urd.Attributes, Attribute{Key: "user", Value: user})
// 	urd.Attributes = append(urd.Attributes, Attribute{Key: "nodetype", Value: NodeType})

// 	regData, _ := json.Marshal(urd)
// 	regDataStr := base64.StdEncoding.EncodeToString(regData)

// 	// prepare command line params
// 	fnArg := fmt.Sprintf("-fn=%s", "createUser")
// 	userArg := fmt.Sprintf("-user=%s", userId)
// 	regDataArg := fmt.Sprintf("-regdata=%s", regDataStr)
// 	outfileArg := fmt.Sprintf("-outfile=%s", outfile)

// 	// launch IdentityApp
// 	cmd := exec.Command(IdentityApp, fnArg, userArg, regDataArg, outfileArg)
// 	if cmd == nil {
// 		return nil, errors.New("failed to launch IdentityApp")
// 	}

// 	err = cmd.Run()
// 	if err != nil {
// 		return nil, err
// 	}

// 	out, err := ioutil.ReadFile(outfile)
// 	if err != nil {
// 		return nil, err
// 	}

// 	type IdentityAppResp struct {
// 		Data   interface{} `json:"Data"`
// 		Status bool        `json:"Status"`
// 	}

// 	var iresp IdentityAppResp
// 	if err := json.Unmarshal(out, &iresp); err != nil {
// 		return nil, fmt.Errorf("data: %s \n error: %w", iresp.Data, err)
// 	}

// 	if !iresp.Status {
// 		return nil, fmt.Errorf("IdentityApp is failed to create identity. %s", iresp.Data)
// 	}

// 	b, err := json.Marshal(iresp.Data)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var resp ledger.UserCrpto
// 	if err := json.Unmarshal(b, &resp); err != nil {
// 		return nil, err
// 	}

// 	return &resp, nil
// }

func (s *Auth) doRegister(u *UserData, reqData *SignupReq) error {

	r := &ledger.RegisterationData{}
	r.UserId = u.UserId
	r.Email = reqData.User

	if _, err := s.Dep.GetLedger().RegisterWarehouseUser(u.Crypto, r); err != nil {
		return err
	}

	return nil
}
