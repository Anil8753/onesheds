package ledger

type UserCrpto struct {
	MSP        string `json:"msp"`
	UserId     string `json:"userId"`
	Cert       string `json:"cert"`
	PrivateKey string `json:"privateKey"`
}
