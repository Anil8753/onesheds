package nethttp

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
	CryptoRetriveFailed
	ServerIssue
)

var mappings map[int]string

func init() {

	mappings = map[int]string{
		Success:             "success",
		WrongCredentials:    "wrong credentials",
		UserAlreadyExist:    "user already exist",
		UserNotExist:        "user not exist",
		UserNotAuthorized:   "user not authorized",
		InvalidRequestData:  "invalid request data",
		CryptoRetriveFailed: "crypto retrive failed",
		ServerIssue:         "server issue",
	}
}

func NewHttpResponse(code int) *HttpResponse {
	codedesc := mappings[code]
	return &HttpResponse{Code: code, CodeDesc: codedesc}
}

func NewHttpResponseWithMsg(code int, msg interface{}) *HttpResponse {
	codedesc := mappings[code]
	return &HttpResponse{Code: code, CodeDesc: codedesc, Data: msg}
}
