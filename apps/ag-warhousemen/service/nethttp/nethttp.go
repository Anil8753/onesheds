package nethttp

type HttpResponse struct {
	Code     int
	CodeDesc string
	Desc     interface{}
}

const (
	Sucess = iota
	WrongCredentials
	UserAlreadyExist
	UserNotExist
	InvalidRequestData
	ServerIssue
)

var mappings map[int]string

func init() {

	mappings = map[int]string{
		Sucess:             "success",
		WrongCredentials:   "wrong credentials",
		UserAlreadyExist:   "user already exist",
		UserNotExist:       "user not exist",
		InvalidRequestData: "invalid request data",
		ServerIssue:        "server issue",
	}
}

func NewHttpResponse(code int) *HttpResponse {
	codedesc := mappings[code]
	return &HttpResponse{Code: code, CodeDesc: codedesc}
}

func NewHttpResponseWithMsg(code int, msg interface{}) *HttpResponse {
	codedesc := mappings[code]
	return &HttpResponse{Code: code, CodeDesc: codedesc, Desc: msg}
}
