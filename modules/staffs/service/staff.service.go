package service

import (
	"context"
	"os"
	"strconv"

	"github.com/BreakDown-CS/erp-setting-cs/internal/helper"
	"github.com/BreakDown-CS/erp-setting-cs/modules/staffs/dto"
	"github.com/BreakDown-CS/erp-setting-cs/modules/staffs/mapper"
	"github.com/BreakDown-CS/erp-setting-cs/modules/staffs/model"
	ports "github.com/BreakDown-CS/erp-setting-cs/modules/staffs/posts"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type service struct {
	repo   ports.StaffRepository
	helper *helper.Uow
	db     *pgxpool.Pool
}

var _ ports.StaffService = (*service)(nil)

func NewUsecase(r ports.StaffRepository, db *pgxpool.Pool) ports.StaffService {
	return &service{
		repo:   r,
		helper: helper.New(db),
		db:     db,
	}
}

func (u *service) CreateStaff(ctx context.Context, req dto.CreateStaffRequest) (res dto.StaffSaveResponse, err error) {

	var staffId uuid.UUID

	err = u.helper.WithTx(ctx, func(tx pgx.Tx) error {

		staffDetail, err := u.repo.CheckDuplicate(ctx, tx, model.Staff{
			Username: req.Username,
		})

		if err != nil {
			return err
		}

		if staffDetail {
			staffId = uuid.Nil
			return nil
		}

		costStr := os.Getenv("BCRYPT_COST")
		cost, _ := strconv.Atoi(costStr)
		passwordHash, _ := helper.HashPassword(req.Password, cost)

		branchID := helper.ParseUUID(req.BranchId)
		departmentID := helper.ParseUUID(req.DepartmentId)
		positionID := helper.ParseUUID(req.PositionId)
		CreatedBy := helper.ParseUUID(req.CreatedBy)

		staff := model.Staff{
			EmployeeCode: req.EmployeeCode,
			FirstName:    req.FirstName,
			LastName:     req.LastName,
			Username:     req.Username,
			PasswordHash: passwordHash,
			BranchID:     &branchID,
			DepartmentID: &departmentID,
			PositionID:   &positionID,
			Status:       "active",
			CreatedBy:    &CreatedBy,
		}

		staffId, err = u.repo.InsertStaff(ctx, tx, staff)
		return err
	})

	if err != nil {
		return res, err
	}

	return dto.StaffSaveResponse{
		StaffId: staffId,
	}, nil
}

func (u *service) GetStaffList(payload dto.GetStaffListRequest) ([]dto.StaffListResponse, int, error) {

	staffList, total, err := u.repo.GetStaffList(payload)

	if err != nil {
		return nil, 0, err
	}

	var result []dto.StaffListResponse

	for _, staff := range staffList {
		result = append(result, mapper.ToStaffList(staff))
	}

	return result, total, nil
}

func (u *service) GetStaffById(id uuid.UUID) (dto.StaffDetailResponse, error) {

	staff, err := u.repo.GetStaffById(id)

	if err != nil {
		return dto.StaffDetailResponse{}, err
	}

	return mapper.ToStaffDetail(staff), nil
}
