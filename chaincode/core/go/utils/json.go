package utils

import "encoding/json"

func ToJSON(data interface{}) (string, error) {

	b, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
