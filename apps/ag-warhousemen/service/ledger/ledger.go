package ledger

import (
	"fmt"
	"os"

	"github.com/hyperledger/fabric-gateway/pkg/client"
	"google.golang.org/grpc"
)

type Ledger struct {
	gwConn *grpc.ClientConn
}

func (s *Ledger) Init() {

	var err error
	s.gwConn, err = newGrpcConnection()
	if err != nil {
		panic(fmt.Errorf("failed to estabilish gateway connection. %w", err))
	}

	//defer s.gwConn.Close()
}

func (s *Ledger) ShutDown() {
	s.gwConn.Close()
}

// GetGateway returns the gateway instance. do not forget to close it after use (gateway.Close())
func (s *Ledger) GetGateway(ucrypt *UserCrpto) (*client.Gateway, error) {
	id, err := newIdentity(ucrypt)
	if err != nil {
		return nil, err
	}

	sign, err := newSign(ucrypt)
	if err != nil {
		return nil, err
	}

	// Create a Gateway connection for a specific client identity.
	gateway, err := client.Connect(id, client.WithSign(sign), client.WithClientConnection(s.gwConn))
	if err != nil {
		return nil, err
	}

	return gateway, nil
}

// GetContract retuns contract instance
func (s *Ledger) GetContract(gateway *client.Gateway, channel string, contractName string) (*client.Contract, error) {

	// Obtain smart contract deployed on the network.
	network := gateway.GetNetwork(channel)
	if network == nil {
		return nil, fmt.Errorf("failed tp get network")
	}

	contract := network.GetContract(contractName)
	if contract == nil {
		return nil, fmt.Errorf("failed tp get contract")
	}

	return contract, nil
}

// GetUserContract retuns contract instance
func (s *Ledger) GetUserContract(ucryp *UserCrpto) (*client.Contract, error) {

	gw, err := s.GetGateway(ucryp)
	if err != nil {
		return nil, err
	}

	contract, err := s.GetContract(gw, os.Getenv("LEDGER_CHANNEL"), os.Getenv("LEDGER_CHAINCODE"))
	if err != nil {
		return nil, err
	}

	return contract, nil
}
