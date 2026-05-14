package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/BreakDown-CS/erp-setting-cs/modules/products/model"
	ports "github.com/BreakDown-CS/erp-setting-cs/modules/products/posts"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) ports.ProductRepository {
	return &repository{db: db}
}

func (r *repository) CheckDuplicateCatOrBrandOrModel(ctx context.Context, tx pgx.Tx, code string, typeMenuProduct string) (bool, error) {

	var exists bool
	var queryCheck string

	switch typeMenuProduct {

	case "Category":
		queryCheck = `
			SELECT EXISTS (
				SELECT 1
				FROM erp.product_categories
				WHERE category_code = $1
			)
		`

	case "Brand":
		queryCheck = `
			SELECT EXISTS (
				SELECT 1
				FROM erp.product_brands
				WHERE brand_code = $1
			)
		`

	case "Model":
		queryCheck = `
			SELECT EXISTS (
				SELECT 1
				FROM erp.product_models
				WHERE model_code = $1
			)
		`

	default:
		return false, fmt.Errorf("invalid typeMenuProduct: %s", typeMenuProduct)
	}

	err := tx.QueryRow(ctx, queryCheck, code).Scan(&exists)

	if err != nil {
		return false, err
	}

	return exists, nil
}

func (r *repository) InsertCategory(ctx context.Context, tx pgx.Tx, payload model.ProductCategories) (uuid.UUID, error) {
	var id uuid.UUID

	queryInsertCategory := `
		INSERT INTO erp.product_categories
			( category_code, category_name )
		VALUES
			( $1, $2 )
		RETURNING id
	`

	err := tx.QueryRow(ctx,
		queryInsertCategory,
		payload.CategoryCode,
		payload.CategoryName,
	).Scan(&id)

	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}

func (r *repository) InsertBrand(ctx context.Context, tx pgx.Tx, payload model.ProductBrands) (uuid.UUID, error) {
	var id uuid.UUID

	queryInsertBrand := `
		INSERT INTO erp.product_brands
			( brand_code, brand_name )
		VALUES
			( $1, $2 )
		RETURNING id
	`

	err := tx.QueryRow(ctx,
		queryInsertBrand,
		payload.BrandCode,
		payload.BrandName,
	).Scan(&id)

	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}

func (r *repository) InsertModel(ctx context.Context, tx pgx.Tx, payload model.ProductModels) (uuid.UUID, error) {
	var id uuid.UUID

	queryInsertModel := `
		INSERT INTO erp.product_models
			( models_code, models_name, brand_id )
		VALUES
			( $1, $2, $3 )
		RETURNING id
	`

	err := tx.QueryRow(ctx,
		queryInsertModel,
		payload.ModelCode,
		payload.ModelName,
		payload.BrandID,
	).Scan(&id)

	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}

func (r *repository) GetListCatOrBrandOrModel(typeMenuProduct string) ([]model.ListCatOrBrandOrModel, error) {

	queries := map[string]string{
		"Category": `
			SELECT id, category_name AS name FROM erp.product_categories`,
		"Brand": `
			SELECT id, brand_name AS name FROM erp.product_brands`,
		"Model": `
			SELECT id, models_name AS name FROM erp.product_models
		`,
	}

	query, ok := queries[typeMenuProduct]
	if !ok {
		return nil, fmt.Errorf("invalid typeMenuProduct: %s", typeMenuProduct)
	}

	log.Println("dataList : ", query)

	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var dataList []model.ListCatOrBrandOrModel

	for rows.Next() {
		var item model.ListCatOrBrandOrModel

		err := rows.Scan(
			&item.Id,
			&item.Name,
		)

		if err != nil {
			return nil, err
		}

		dataList = append(dataList, item)
	}

	log.Println("dataList : ", dataList)

	return dataList, nil
}
