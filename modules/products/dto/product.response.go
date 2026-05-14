package dto

import "github.com/google/uuid"

type ProductsSaveResponse struct {
	Id uuid.UUID `json:"id"`
}

type ListCatOrBrandOrModel struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}
