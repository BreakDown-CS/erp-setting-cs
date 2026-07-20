package usecase

import (
	"context"

	"github.com/BreakDown-CS/erp-setting-cs/internal/helper"
	"github.com/BreakDown-CS/erp-setting-cs/modules/setup/dto"
	ports "github.com/BreakDown-CS/erp-setting-cs/modules/setup/posts"
	"github.com/jackc/pgx/v5/pgxpool"
)

type usecase struct {
	repo   ports.SetupRepository
	helper *helper.Uow
	db     *pgxpool.Pool
}

func NewUsecase(r ports.SetupRepository, db *pgxpool.Pool) ports.SetupUsecase {
	return &usecase{
		repo:   r,
		helper: helper.New(db),
		db:     db,
	}
}

func (u *usecase) GetAllBranches(ctx context.Context) (res []dto.SetUpBool, err error) {
	branches, err := u.repo.FindAllBranches(ctx)
	if err != nil {
		return nil, err
	}

	response := []dto.SetUpBool{}

	for _, item := range branches {
		response = append(response, dto.SetUpBool{
			Label: item.Name,
			Value: item.Id,
		})
	}

	return response, nil
}

func (u *usecase) GetAllDepartment(ctx context.Context) (res []dto.SetUpBool, err error) {
	branches, err := u.repo.FindAllDepartment(ctx)
	if err != nil {
		return nil, err
	}

	response := []dto.SetUpBool{}

	for _, item := range branches {
		response = append(response, dto.SetUpBool{
			Label: item.Name,
			Value: item.Id,
		})
	}

	return response, nil
}

func (u *usecase) GetAllEmployeesStatus(ctx context.Context) (res []dto.SetUpBool, err error) {
	branches, err := u.repo.FindAllEmployeesStatus(ctx)
	if err != nil {
		return nil, err
	}

	response := []dto.SetUpBool{}

	for _, item := range branches {
		response = append(response, dto.SetUpBool{
			Label: item.Name,
			Value: item.Id,
		})
	}

	return response, nil
}

func (u *usecase) GetAllPositions(ctx context.Context) (res []dto.SetUpBool, err error) {
	branches, err := u.repo.FindAllPositions(ctx)
	if err != nil {
		return nil, err
	}

	response := []dto.SetUpBool{}

	for _, item := range branches {
		response = append(response, dto.SetUpBool{
			Label: item.Name,
			Value: item.Id,
		})
	}

	return response, nil
}
