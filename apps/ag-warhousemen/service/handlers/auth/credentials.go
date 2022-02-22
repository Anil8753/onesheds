package auth

import (
	"html"
	"strings"

	"github.com/anil8753/onesheds/apps/warehousemen/service/interfaces"
	"golang.org/x/crypto/bcrypt"
)

type Credentials struct {
	UserId   string
	User     string
	Password string
}

func (s *Credentials) SaveUser(db interfaces.Database) (*Credentials, error) {

	if err := s.BeforeSave(); err != nil {
		return nil, err
	}

	if err := db.Put(s.User, s); err != nil {
		return nil, err
	}

	return s, nil
}

func (s *Credentials) BeforeSave() error {

	//turn password into hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(s.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	s.Password = string(hashedPassword)

	//remove spaces in username
	s.User = html.EscapeString(strings.TrimSpace(s.User))

	return nil
}
