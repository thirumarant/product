package utils

import (
	"github.com/google/uuid"
	"strings"
)

func GenerateUUID() string {
	uuid := uuid.New()
	return strings.ToUpper(uuid.String())
}

func IsValidUUID(id string) bool {
	_, err := uuid.Parse(id)
	if err != nil {
		return false
	}
	return true
}
