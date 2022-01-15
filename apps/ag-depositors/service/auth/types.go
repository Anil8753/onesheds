package auth

import (
	"github.com/anil8753/onesheds/apps/warehousemen/service/interfaces"
	"github.com/anil8753/onesheds/apps/warehousemen/service/ledger"
)

type Auth struct {
	Database interfaces.Database
	Ledger   *ledger.Ledger
}

type Attribute struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type UserRegistrationData struct {
	UserId     string      `json:"userId"`
	Attributes []Attribute `json:"attributes"`
}
