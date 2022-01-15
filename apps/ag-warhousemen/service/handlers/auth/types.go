package auth

import (
	"github.com/anil8753/onesheds/apps/warehousemen/service/interfaces"
)

type Auth struct {
	Dep interfaces.HandlerDependency
}

type Attribute struct {
	Key   string
	Value string
}

type UserRegistrationData struct {
	UserId     string
	Attributes []Attribute
}
