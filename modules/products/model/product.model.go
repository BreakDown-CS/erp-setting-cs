package model

import (
	"time"

	"github.com/google/uuid"
)

type ProductCategories struct {
	ID           uuid.UUID
	CategoryCode string
	CategoryName string
	Status       string
	CreatedAt    time.Time
}

type ProductBrands struct {
	ID        uuid.UUID
	BrandCode string
	BrandName string
	Status    string
	CreatedAt time.Time
}

type ProductModels struct {
	ID        uuid.UUID
	ModelCode string
	ModelName string
	BrandID   uuid.UUID
	Status    string
	CreatedAt time.Time
}

type Products struct {
	ID            uuid.UUID
	ProductCode   string
	ProductName   string
	CategoryID    uuid.UUID
	BrandID       uuid.UUID
	ModelID       uuid.UUID
	Descriptions  string
	Unit          string
	StandardPrice float64
	Status        string
	CreatedBy     uuid.UUID
	CreatedAt     time.Time
	UpdatedBy     *uuid.UUID
	UpdatedAt     time.Time
}

type ListCatOrBrandOrModel struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}
