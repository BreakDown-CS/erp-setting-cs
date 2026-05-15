package ports

import (
	"context"

	"github.com/BreakDown-CS/erp-setting-cs/modules/products/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type ProductRepository interface {
	CheckDuplicateCatOrBrandOrModel(context.Context, pgx.Tx, string, string) (bool, error)
	InsertCategory(context.Context, pgx.Tx, model.ProductCategories) (uuid.UUID, error)
	InsertBrand(context.Context, pgx.Tx, model.ProductBrands) (uuid.UUID, error)
	InsertModel(context.Context, pgx.Tx, model.ProductModels) (uuid.UUID, error)
	GetListCatOrBrandOrModel(string) ([]model.ListCatOrBrandOrModel, error)
	CheckDuplicateProduct(context.Context, pgx.Tx, model.Products) (bool, error)
	InsertProduct(context.Context, pgx.Tx, model.Products) (uuid.UUID, error)
}
