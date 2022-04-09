package ca

import (
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
)

type UserRegistrationData struct {
	User       string      `json:"user"`
	NodeType   string      `json:"nodeType"`
	Attributes []Attribute `json:"attributes"`
}

type Attribute struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type UserIdentity struct {
	MSP        string `json:"msp"`
	UserId     string `json:"userId"`
	Cert       string `json:"cert"`
	PrivateKey string `json:"privateKey"`
}

func RegEnrollUser(regData *UserRegistrationData) (*UserIdentity, error) {

	mspClient, err := GetMSPClient()
	if err != nil {
		return nil, err
	}

	userId := fmt.Sprintf("user_%s", uuid.New().String())

	attrs := []msp.Attribute{}
	for _, attr := range regData.Attributes {
		attrs = append(attrs, msp.Attribute{Name: attr.Key, Value: attr.Value, ECert: true})
	}

	attrs = append(attrs, msp.Attribute{Name: "userId", Value: userId, ECert: true})
	attrs = append(attrs, msp.Attribute{Name: "user", Value: regData.User, ECert: true})
	attrs = append(attrs, msp.Attribute{Name: "nodeType", Value: regData.NodeType, ECert: true})

	rr := &msp.RegistrationRequest{
		Name:       userId,
		Attributes: attrs,
	}

	log.Println("Register", rr)

	userpw, err := mspClient.Register(rr)
	if err != nil {
		return nil, fmt.Errorf("failed to register user: %s. %w", userId, err)
	}

	//Enroll with returned password
	err = mspClient.Enroll(userId, msp.WithSecret(userpw))
	if err != nil {
		return nil, fmt.Errorf("failed to enroll user %s. %w", userId, err)
	}

	return GetSigningIdentityWithMSPClient(mspClient, userId)
}
