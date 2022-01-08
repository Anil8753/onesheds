package ledger

import (
	"encoding/json"
)

func (s *Ledger) RegisterWarehouse(ucryp *UserCrpto, r *AssetData) ([]byte, error) {

	contract, err := s.GetUserContract(ucryp)
	if err != nil {
		return nil, err
	}

	inBytes, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}

	data, err := contract.SubmitTransaction("RegisterWarehouse", string(inBytes))
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *Ledger) GetUserWarehouses(ucryp *UserCrpto) ([]byte, error) {

	contract, err := s.GetUserContract(ucryp)
	if err != nil {
		return nil, err
	}

	data, err := contract.EvaluateTransaction("GetWarehouseByOwnerId", ucryp.UserId)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *Ledger) UpdateWarehouse(ucryp *UserCrpto, r *AssetData) ([]byte, error) {

	contract, err := s.GetUserContract(ucryp)
	if err != nil {
		return nil, err
	}

	inBytes, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}

	data, err := contract.SubmitTransaction("UpdateWarehouse", string(inBytes))
	if err != nil {
		return nil, err
	}

	return data, nil
}
