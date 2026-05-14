package dto

type CreateProductsRequest struct {
	Code    string  `json:"code" validate:"required,min=1,max=100"`
	Name    string  `json:"name" validate:"required,min=1,max=100"`
	BrandId *string `json:"brand_id"`
}
