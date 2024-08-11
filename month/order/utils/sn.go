package utils

import "github.com/google/uuid"

func GenerateSN() string {
	code := uuid.New().String()
	return code[:8]
}
