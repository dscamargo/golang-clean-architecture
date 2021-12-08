package adapters

import (
	"golang.org/x/crypto/bcrypt"
)

type Hasher struct{}

func NewHasher() Hasher {
	return Hasher{}
}

func (h *Hasher) Hash(value string) string {
	valueInBytes := []byte(value)
	responseBytes, _ := bcrypt.GenerateFromPassword(valueInBytes, 10)
	return string(responseBytes)
}
