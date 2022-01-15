package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
)

type UserRegistrationData struct {
	UserId string `json:"userId"`

	Attributes []struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	} `json:"attributes"`
}

type UserIdentity struct {
	MSP        string `json:"msp"`
	UserId     string `json:"userId"`
	Cert       string `json:"cert"`
	PrivateKey string `json:"privateKey"`
}

func CreateIdentityHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "POST" {
			http.Error(w, "Method is not supported.", http.StatusNotFound)
			return
		}

		var urd UserRegistrationData
		if err := json.NewDecoder(r.Body).Decode(&urd); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if urd.UserId == "" {
			http.Error(w, "UserId is must", http.StatusBadRequest)
			return
		}

		log.Println("UserRegistrationData", urd)

		ui, err := RegEnrollUser(&urd)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		b, err := json.Marshal(ui)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}
}

func RegEnrollUser(regData *UserRegistrationData) (*UserIdentity, error) {

	mspClient, err := GetMSPClient()
	if err != nil {
		return nil, err
	}

	attrs := []msp.Attribute{}
	for _, attr := range regData.Attributes {
		attrs = append(attrs, msp.Attribute{Name: attr.Key, Value: attr.Value, ECert: true})
	}

	userId := regData.UserId

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

	identity, err := mspClient.GetSigningIdentity(userId)

	if err != nil {
		return nil, fmt.Errorf("failed to get signing identity for user %s. %w", userId, err)
	}

	priKeyBytes, _ := identity.PrivateKey().Bytes()
	certBytes := identity.EnrollmentCertificate()

	return &UserIdentity{
		MSP:        identity.Identifier().MSPID,
		UserId:     userId,
		Cert:       string(certBytes),
		PrivateKey: string(priKeyBytes),
	}, nil
}
