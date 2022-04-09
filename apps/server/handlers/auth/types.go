package auth

import (
	"github.com/anil8753/onesheds/interfaces"
	"github.com/anil8753/onesheds/ledger"
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
	User       string `json:"user"`
	NodeType   string `json:"nodeType"`
	Attributes []Attribute
}
