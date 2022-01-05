package ledger

import (
	"encoding/json"
	"log"
)

func (s *Ledger) GetWarehouseUser(ucryp *UserCrpto) ([]byte, error) {

	contract, err := s.GetUserContract(ucryp)
	if err != nil {
		return nil, err
	}

	data, err := contract.EvaluateTransaction("GetWarehouseUser", ucryp.UserId)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *Ledger) RegisterWarehouseUser(ucryp *UserCrpto, r *RegisterationData) ([]byte, error) {

	contract, err := s.GetUserContract(ucryp)
	if err != nil {
		return nil, err
	}

	inBytes, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}

	log.Println(string(inBytes))

	data, err := contract.SubmitTransaction("RegisterWarehouseUser", string(inBytes))
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *Ledger) UpdateWarehouseUser(ucryp *UserCrpto, r *RegisterationData) ([]byte, error) {

	contract, err := s.GetUserContract(ucryp)
	if err != nil {
		return nil, err
	}

	inBytes, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}

	log.Println(string(inBytes))

	data, err := contract.SubmitTransaction("UpdateWarehouseUser", string(inBytes))
	if err != nil {
		return nil, err
	}

	return data, nil
}
