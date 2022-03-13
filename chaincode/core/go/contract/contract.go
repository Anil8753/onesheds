package contract

import (
	"crypto/x509"
	"encoding/json"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type Contract struct {
	contractapi.Contract
}

func (c *Contract) Init(ctx contractapi.TransactionContextInterface) error {
	return nil
}

func (c *Contract) GetIdentity(ctx contractapi.TransactionContextInterface) (string, error) {

	type Identity struct {
		Cert *x509.Certificate
		ID   string
		MSP  string
	}

	identity := ctx.GetClientIdentity()

	cert, err := identity.GetX509Certificate()
	if err != nil {
		return "", err
	}

	id, err := identity.GetID()
	if err != nil {
		return "", err
	}

	msp, err := identity.GetMSPID()
	if err != nil {
		return "", err
	}

	i := Identity{Cert: cert, ID: id, MSP: msp}

	b, err := json.MarshalIndent(i, "", " ")
	if err != nil {
		return "", err
	}

	return string(b), nil
}
