package usecase

import (
	"context"

	"github.com/BreakDown-CS/erp-setting-cs/internal/helper"
	"github.com/BreakDown-CS/erp-setting-cs/modules/products/dto"
	"github.com/BreakDown-CS/erp-setting-cs/modules/products/model"
	ports "github.com/BreakDown-CS/erp-setting-cs/modules/products/posts"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type usecase struct {
	repo   ports.ProductRepository
	helper *helper.Uow
	db     *pgxpool.Pool
}

func NewUsecase(r ports.ProductRepository, db *pgxpool.Pool) ports.ProductUsecase {
	return &usecase{
		repo:   r,
		helper: helper.New(db),
		db:     db,
	}
}

func (u *usecase) CreateCategory(ctx context.Context, req dto.CreateProductsRequest) (dto.ProductsSaveResponse, error) {

	var categoryId uuid.UUID

	err := u.helper.WithTx(ctx, func(tx pgx.Tx) error {

		isCategory, err := u.repo.CheckDuplicateCatOrBrandOrModel(ctx, tx, req.Code, "Category")

		if err != nil {
			return err
		}

		if isCategory {
			categoryId = uuid.Nil
			return nil
		}

		category := model.ProductCategories{
			CategoryCode: req.Code,
			CategoryName: req.Name,
		}

		categoryId, err = u.repo.InsertCategory(ctx, tx, category)
		return err
	})

	if err != nil {
		return dto.ProductsSaveResponse{}, err
	}

	return dto.ProductsSaveResponse{
		Id: categoryId,
	}, nil
}

func (u *usecase) CreateBrand(ctx context.Context, req dto.CreateProductsRequest) (dto.ProductsSaveResponse, error) {

	var brandId uuid.UUID

	err := u.helper.WithTx(ctx, func(tx pgx.Tx) error {

		isBrand, err := u.repo.CheckDuplicateCatOrBrandOrModel(ctx, tx, req.Code, "Brand")

		if err != nil {
			return err
		}

		if isBrand {
			brandId = uuid.Nil
			return nil
		}

		brand := model.ProductBrands{
			BrandCode: req.Code,
			BrandName: req.Name,
		}

		brandId, err = u.repo.InsertBrand(ctx, tx, brand)
		return err
	})

	if err != nil {
		return dto.ProductsSaveResponse{}, err
	}

	return dto.ProductsSaveResponse{
		Id: brandId,
	}, nil
}

func (u *usecase) CreateModel(ctx context.Context, req dto.CreateProductsRequest) (dto.ProductsSaveResponse, error) {

	var modelId uuid.UUID

	err := u.helper.WithTx(ctx, func(tx pgx.Tx) error {

		isModel, err := u.repo.CheckDuplicateCatOrBrandOrModel(ctx, tx, req.Code, "Model")

		if err != nil {
			return err
		}

		if isModel {
			modelId = uuid.Nil
			return nil
		}

		branchID := helper.ParseUUID(*req.BrandId)

		model := model.ProductModels{
			ModelCode: req.Code,
			ModelName: req.Name,
			BrandID:   branchID,
		}

		modelId, err = u.repo.InsertModel(ctx, tx, model)
		return nil
	})

	if err != nil {
		return dto.ProductsSaveResponse{}, err
	}

	return dto.ProductsSaveResponse{
		Id: modelId,
	}, nil
}

func (u *usecase) ListCategories() ([]dto.ListCatOrBrandOrModel, error) {

	listCategories, err := u.repo.GetListCatOrBrandOrModel("Category")
	if err != nil {
		return nil, err
	}

	result := make([]dto.ListCatOrBrandOrModel, 0, len(listCategories))

	for _, item := range listCategories {
		result = append(result, dto.ListCatOrBrandOrModel{
			Id:   item.Id,
			Name: item.Name,
		})
	}

	return result, nil
}

func (u *usecase) ListBrands() ([]dto.ListCatOrBrandOrModel, error) {

	listBrands, err := u.repo.GetListCatOrBrandOrModel("Brand")
	if err != nil {
		return nil, err
	}

	result := make([]dto.ListCatOrBrandOrModel, 0, len(listBrands))

	for _, item := range listBrands {
		result = append(result, dto.ListCatOrBrandOrModel{
			Id:   item.Id,
			Name: item.Name,
		})
	}

	return result, nil
}

func (u *usecase) ListModels() ([]dto.ListCatOrBrandOrModel, error) {

	listModels, err := u.repo.GetListCatOrBrandOrModel("Model")
	if err != nil {
		return nil, err
	}

	result := make([]dto.ListCatOrBrandOrModel, 0, len(listModels))

	for _, item := range listModels {
		result = append(result, dto.ListCatOrBrandOrModel{
			Id:   item.Id,
			Name: item.Name,
		})
	}

	return result, nil
}
