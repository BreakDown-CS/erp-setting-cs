package dto

import "github.com/google/uuid"

type SetUpBool struct {
	Label string    `json:"label"`
	Value uuid.UUID `json:"value"`
}
