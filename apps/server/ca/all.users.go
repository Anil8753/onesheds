package ca

import "github.com/hyperledger/fabric-sdk-go/pkg/client/msp"

func GetAllIdentities() ([]*msp.IdentityResponse, error) {

	mspClient, err := GetMSPClient()
	if err != nil {
		return nil, err
	}

	return mspClient.GetAllIdentities()
}
