package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
)

func AllIdentitiesHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "GET" {
			http.Error(w, "Method is not supported.", http.StatusNotFound)
			return
		}

		resp, err := GetAllIdentities()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		b, err := json.Marshal(resp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}
}

func GetAllIdentities() ([]*msp.IdentityResponse, error) {

	mspClient, err := GetMSPClient()
	if err != nil {
		return nil, err
	}

	return mspClient.GetAllIdentities()
}
