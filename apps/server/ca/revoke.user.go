package ca

import (
	"fmt"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
)

// Binding from JSON
type RevokeReq struct {
	UserId string `json:"userId" binding:"required"`
	Reason string ` json:"reason" binding:"required"`
}

func RevokeUserHandler(userId string, reason string) (interface{}, error) {

	mspClient, err := GetMSPClient()
	if err != nil {
		return nil, err
	}

	if _, err := GetSigningIdentityWithMSPClient(mspClient, userId); err != nil {
		return nil, err
	}

	resp, err := RevokeUser(mspClient, userId, reason)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func RevokeUser(mspClient *msp.Client, userId string, reason string) (*msp.RevocationResponse, error) {

	request := &msp.RevocationRequest{
		Name:   userId,
		Reason: reason,
	}

	resp, err := mspClient.Revoke(request)
	if err != nil {
		return nil, fmt.Errorf("user %s revoke failed. %w", userId, err)
	}

	return resp, nil
}
