package ca

import (
	"fmt"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
)

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
