package model

import "github.com/google/uuid"

type SetUpOptionUUID struct {
	Id   uuid.UUID `json:"id" db:"id"`
	Name string    `json:"name" db:"name"`
}
