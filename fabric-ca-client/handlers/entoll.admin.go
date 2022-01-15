package handlers

import (
	"fmt"
	"net/http"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
)

func EnrollAdminHandler() func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "POST" {
			http.Error(w, "Method is not supported.", http.StatusNotFound)
			return
		}

		if err := EnrollAdmin(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	}
}

func EnrollAdmin() error {

	mspClient, caconfig, err := GetMSPClientWithCAConfig()
	if err != nil {
		return err
	}

	// Now try to enroll the admin with its configured ID and password
	err = mspClient.Enroll(caconfig.Registrar.EnrollID, msp.WithSecret(caconfig.Registrar.EnrollSecret))
	if err != nil {
		return fmt.Errorf("failed to enroll the admin. %w", err)
	}

	return nil
}
