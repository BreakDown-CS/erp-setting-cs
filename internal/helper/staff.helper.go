package helper

import (
	"fmt"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// StaffStatus constants to avoid hardcoded strings
const (
	StaffStatusActive    = "active"
	StaffStatusInactive  = "inactive"
	StaffStatusSuspended = "suspended"
)

// ParseUUID parses a UUID string, returning an error if invalid.
// If value is empty, returns uuid.Nil with no error (optional field).
func ParseUUID(value string) (uuid.UUID, error) {
	if value == "" {
		return uuid.Nil, nil
	}
	id, err := uuid.Parse(value)
	if err != nil {
		return uuid.Nil, fmt.Errorf("invalid UUID '%s': %w", value, err)
	}
	return id, nil
}

// ParseUUIDPtr parses a UUID string and returns a pointer.
// Returns nil if value is empty (optional field).
func ParseUUIDPtr(value string) (*uuid.UUID, error) {
	if value == "" {
		return nil, nil
	}
	id, err := uuid.Parse(value)
	if err != nil {
		return nil, fmt.Errorf("invalid UUID '%s': %w", value, err)
	}
	return &id, nil
}

// HashPassword generates a bcrypt hash with the given cost.
func HashPassword(password string, cost int) (string, error) {
	if cost < bcrypt.MinCost || cost > bcrypt.MaxCost {
		cost = bcrypt.DefaultCost
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return "", fmt.Errorf("bcrypt hash failed: %w", err)
	}
	return string(hash), nil
}
