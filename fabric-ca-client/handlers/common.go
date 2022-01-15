package handlers

import (
	"errors"
	"fmt"
	"os"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
	pmsp "github.com/hyperledger/fabric-sdk-go/pkg/common/providers/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	rmsp "github.com/hyperledger/fabric-sdk-go/pkg/msp"
)

func GetMSPClient() (*msp.Client, error) {

	//Load configuration from connection profile
	cp := os.Getenv("ConnectionProfile")
	cnfg := config.FromFile(cp)
	sdk, err := fabsdk.New(cnfg)
	if err != nil {
		return nil, fmt.Errorf("failed to create new SDK. %w", err)
	}
	defer sdk.Close()

	ctxProvider := sdk.Context()
	mspClient, err := msp.New(
		ctxProvider,
		msp.WithOrg(os.Getenv("ORG")),
		msp.WithCAInstance(os.Getenv("CA")),
	)

	if err != nil {
		return nil, fmt.Errorf("failed create msp client. %w", err)
	}

	return mspClient, nil
}

func GetMSPClientWithCAConfig() (*msp.Client, *pmsp.CAConfig, error) {

	//Load configuration from connection profile
	cp := os.Getenv("ConnectionProfile")
	cnfg := config.FromFile(cp)
	sdk, err := fabsdk.New(cnfg)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create new SDK. %w", err)
	}
	defer sdk.Close()

	ctxProvider := sdk.Context()
	mspClient, err := msp.New(
		ctxProvider,
		msp.WithOrg(os.Getenv("ORG")),
		msp.WithCAInstance(os.Getenv("CA")),
	)

	if err != nil {
		return nil, nil, fmt.Errorf("failed create msp client. %w", err)
	}

	// Try to get some configuration data from the connection profile
	sdkcfg, err := sdk.Config()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to SDK config. %w", err)
	}

	idcfg, err := rmsp.ConfigFromBackend(sdkcfg)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to rmsp config from backend. %w", err)
	}

	caconfig, ok := idcfg.CAConfig(os.Getenv("CA"))
	if !ok {
		return nil, nil, errors.New("could not get the caconfiguration")
	}

	return mspClient, caconfig, nil
}
