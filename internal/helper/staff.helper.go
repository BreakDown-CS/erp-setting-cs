package helper

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func ParseUUID(value string) uuid.UUID {
	id, err := uuid.Parse(value)
	if err != nil {
		return uuid.Nil
	}
	return id
}

func HashPassword(password string, cost int) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	return string(hash), err
}
