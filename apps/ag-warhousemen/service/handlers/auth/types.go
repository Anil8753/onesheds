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
	Key   string
	Value string
}

type UserRegistrationData struct {
	UserId     string
	Attributes []Attribute
}
