package ledger

func (s *Ledger) GetUserIdentity(ucryp *UserCrpto) ([]byte, error) {

	contract, err := s.GetUserContract(ucryp)
	if err != nil {
		return nil, err
	}

	data, err := contract.EvaluateTransaction("GetIdentity")
	if err != nil {
		return nil, err
	}

	return data, nil
}
