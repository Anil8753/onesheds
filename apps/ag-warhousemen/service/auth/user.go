package auth

import (
	"html"
	"strings"

	"github.com/anil8753/onesheds/apps/warehousemen/service/interfaces"
	"github.com/anil8753/onesheds/apps/warehousemen/service/ledger"
	"golang.org/x/crypto/bcrypt"
)

type UserData struct {
	UserUniqueId string
	User         string
	Password     string
	Crypto       *ledger.UserCrpto
}

func (u *UserData) SaveUser(db interfaces.Database) (*UserData, error) {

	if err := db.Put(u.User, u); err != nil {
		return &UserData{}, err
	}

	return u, nil
}

func (u *UserData) BeforeSave() error {

	//turn password into hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)

	//remove spaces in username
	u.User = html.EscapeString(strings.TrimSpace(u.User))

	return nil
}