package user

import (
	"encoding/json"
	"errors"
)

const UserDocType = "UserRegData"

type RegisterationData struct {
	DocType  string `json:"docType,omitempty"`
	UniqueId string `json:"uniqueId"`
	Status   string `json:"status,omitempty"`

	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	Phone     string `json:"phone,omitempty"`
	Email     string `json:"email,omitempty"`
	Address   string `json:"address,omitempty"`
	PINCode   string `json:"pincode,omitempty"`
	City      string `json:"city,omitempty"`
	District  string `json:"district,omitempty"`
	State     string `json:"state,omitempty"`

	PANCard string `json:"pancard,omitempty"`
	Aadhar  string `json:"aadharcard,omitempty"`
}

func NewRegisterationData(input string) (*RegisterationData, error) {

	rBytes := []byte(input)

	var regData RegisterationData
	if err := json.Unmarshal(rBytes, &regData); err != nil {
		return nil, err
	}

	if regData.UniqueId == "" {
		return nil, errors.New("UniqueId is mandatory")
	}

	regData.DocType = UserDocType

	return &regData, nil
}
