package Util

import (
	"github.com/google/uuid"
)

// GenerateID generates a new UUID
func GenerateID(num int) string {
	return uuid.New().String()[:num]
}
