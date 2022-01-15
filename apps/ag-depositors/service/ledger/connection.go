package ledger

import (
	"crypto/x509"
	"encoding/pem"
	"errors"
	"os"

	"github.com/hyperledger/fabric-gateway/pkg/identity"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// newGrpcConnection creates a gRPC connection to the Gateway server.
func newGrpcConnection() (*grpc.ClientConn, error) {

	transportCredentials, err := credentials.NewClientTLSFromFile(os.Getenv("TLS_CERT_PATH"), os.Getenv("PEER_URL"))
	if err != nil {
		return nil, err
	}

	connection, err := grpc.Dial(os.Getenv("PEER_ENDPOINT"), grpc.WithTransportCredentials(transportCredentials))
	if err != nil {
		return nil, err
	}

	return connection, nil
}

// newIdentity creates a client identity for this Gateway connection using an X.509 certificate.
func newIdentity(ucrypt *UserCrpto) (*identity.X509Identity, error) {

	x509Cert, err := identity.CertificateFromPEM([]byte(ucrypt.Cert))
	if err != nil {
		return nil, err
	}

	id, err := identity.NewX509Identity(ucrypt.MSP, x509Cert)
	if err != nil {
		return nil, err
	}

	return id, nil
}

// newSign creates a function that generates a digital signature from a message digest using a private key.
func newSign(ucrypt *UserCrpto) (identity.Sign, error) {

	block, _ := pem.Decode([]byte(ucrypt.PrivateKey))
	if block == nil {
		return nil, errors.New("failed to decode private key")
	}

	privateKey, err := x509.ParseECPrivateKey(block.Bytes)

	if err != nil {
		return nil, err
	}

	// privateKey, err := identity.PrivateKeyFromPEM(privateKeyPEM)
	// if err != nil {
	// 	return nil, err
	// }

	sign, err := identity.NewPrivateKeySign(privateKey)
	if err != nil {
		return nil, err
	}

	return sign, nil
}
