package dto

type CreateProductsRequest struct {
	Code    string  `json:"code" validate:"required,min=1,max=100"`
	Name    string  `json:"name" validate:"required,min=1,max=100"`
	BrandId *string `json:"brand_id"`
}

type CreateProducts struct {
	ProductCode   string  `json:"product_code" validate:"required,min=1,max=100"`
	ProductName   string  `json:"product_name" validate:"required,min=1,max=100"`
	CategoryID    string  `json:"category_id" validate:"required,min=1,max=100"`
	BrandID       string  `json:"brand_id" validate:"required,min=1,max=100"`
	ModelID       string  `json:"model_id" validate:"required,min=1,max=100"`
	Descriptions  string  `json:"descriptions" validate:"required,min=1,max=100"`
	Unit          string  `json:"unit" validate:"required,min=1,max=100"`
	StandardPrice float64 `json:"standard_price" validate:"required,min=1,max=100"`
	StaffId       string  `json:"staff_id" validate:"required,min=1,max=100"`
}
