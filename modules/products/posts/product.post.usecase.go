package ports

import (
	"context"

	"github.com/BreakDown-CS/erp-setting-cs/modules/products/dto"
)

type ProductUsecase interface {
	CreateCategory(ctx context.Context, req dto.CreateProductsRequest) (dto.ProductsSaveResponse, error)
	CreateBrand(ctx context.Context, req dto.CreateProductsRequest) (dto.ProductsSaveResponse, error)
	CreateModel(ctx context.Context, req dto.CreateProductsRequest) (dto.ProductsSaveResponse, error)

	ListCategories() ([]dto.ListCatOrBrandOrModel, error)
	ListBrands() ([]dto.ListCatOrBrandOrModel, error)
	ListModels() ([]dto.ListCatOrBrandOrModel, error)
}
