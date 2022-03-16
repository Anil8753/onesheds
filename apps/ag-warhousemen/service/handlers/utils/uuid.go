package utils

import (
	"fmt"

	"github.com/hashicorp/go-uuid"
)

func GenerateUUID(prefix string) (string, error) {
	s, err := uuid.GenerateUUID()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s_%s", prefix, s), nil
}
