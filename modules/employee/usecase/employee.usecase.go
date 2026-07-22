package usecase

import (
	"context"
	"log"

	"github.com/BreakDown-CS/erp-setting-cs/internal/helper"
	"github.com/BreakDown-CS/erp-setting-cs/modules/employee/dto"
	"github.com/BreakDown-CS/erp-setting-cs/modules/employee/model"
	ports "github.com/BreakDown-CS/erp-setting-cs/modules/employee/posts"
	"github.com/google/uuid"
)

type EmployeeUsecase struct {
	repo ports.EmployeeRepository
}

func NewUsecase(r ports.EmployeeRepository) ports.EmployeeUsecase {
	return &EmployeeUsecase{
		repo: r,
	}
}

func (u *EmployeeUsecase) InsertEmployee(ctx context.Context, payload dto.InsertEmpolyeeRequest) (dto.EmployeeResponse, error) {
	tx, err := u.repo.BeginTransaction(ctx)
	if err != nil {
		return dto.EmployeeResponse{}, err
	}

	defer func() {
		if r := recover(); r != nil {
			_ = tx.Rollback(ctx)
			log.Println("Recovered in EmployeeUsecase, rolling back:", r)
		}
	}()

	branchId, err := u.repo.GetOneBranches(ctx, tx, payload.BranchId)
	if err != nil {
		tx.Rollback(ctx)
		return dto.EmployeeResponse{}, err
	}

	departmentId, err := u.repo.GetOneDepartment(ctx, tx, payload.DepartmentId)
	if err != nil {
		tx.Rollback(ctx)
		return dto.EmployeeResponse{}, err
	}

	statusId, err := u.repo.GetOneEmployeeStatus(ctx, tx, payload.StatusId)
	if err != nil {
		tx.Rollback(ctx)
		return dto.EmployeeResponse{}, err
	}

	positionId, err := u.repo.GetOnePosition(ctx, tx, payload.PositionId)
	if err != nil {
		tx.Rollback(ctx)
		return dto.EmployeeResponse{}, err
	}

	Birthday, err := helper.DateStringToDate(payload.Birthday)
	if err != nil {
		tx.Rollback(ctx)
		return dto.EmployeeResponse{}, err
	}

	nextEmpCode, err := u.repo.GetNextEmployeeCode(ctx, tx)
	if err != nil {
		tx.Rollback(ctx)
		return dto.EmployeeResponse{}, err
	}

	payloadInsertEmployee := model.EmployeeInsertModel{
		EmployeeCode: nextEmpCode,
		FirstNameTh:  payload.FirstNameTh,
		LastNameTh:   payload.LastNameTh,
		FirstNameEn:  payload.FirstNameEn,
		LastNameEn:   payload.LastNameEn,
		NickName:     payload.NickName,
		Gender:       payload.Gender,
		Birthday:     Birthday,
		Email:        payload.Email,
		Phone:        payload.Phone,
		BranchId:     branchId,
		DepartmentId: departmentId,
		PositionId:   positionId,
		StatusId:     statusId,
		Remark:       payload.Remark,
	}

	employeeId, err := u.repo.SaveEmployee(ctx, tx, payloadInsertEmployee)
	if err != nil {
		tx.Rollback(ctx)
		return dto.EmployeeResponse{}, err
	}

	nextUsername, err := u.repo.GetNextUsername(ctx, tx)
	if err != nil {
		tx.Rollback(ctx)
		return dto.EmployeeResponse{}, err
	}

	bcryptCost := helper.GetEnvInt("BCRYPT_COST", 15)

	passwordHash, err := helper.HashPassword(payload.Password, bcryptCost)
	if err != nil {
		tx.Rollback(ctx)
		return dto.EmployeeResponse{}, err
	}

	payloadInsertUser := model.UserInsertModel{
		Username:     nextUsername,
		PasswordHash: passwordHash,
		EmployeeId:   employeeId,
	}

	userId, err := u.repo.SaveUser(ctx, tx, payloadInsertUser)
	if err != nil {
		tx.Rollback(ctx)
		return dto.EmployeeResponse{}, err
	}

	userUUId, err := u.repo.GetUUIDForAll(ctx, tx, userId, "user")
	if err != nil {
		tx.Rollback(ctx)
		return dto.EmployeeResponse{}, err
	}

	employeeUUId, err := u.repo.GetUUIDForAll(ctx, tx, employeeId, "employee")
	if err != nil {
		tx.Rollback(ctx)
		return dto.EmployeeResponse{}, err
	}

	if err := tx.Commit(ctx); err != nil {
		return dto.EmployeeResponse{}, err
	}

	return dto.EmployeeResponse{
		UserId:     userUUId,
		EmployeeId: employeeUUId,
	}, nil
}

func (u *EmployeeUsecase) GetEmployeeList(ctx context.Context, payload dto.GetListEmpolyeeRequest) (dto.GetListEmpolyeeResponse, error) {
	employeesList, total, err := u.repo.GetEmployeesList(ctx, payload)
	if err != nil {
		return dto.GetListEmpolyeeResponse{}, err
	}

	erployeeList := []dto.EmpolyeList{}
	for _, emp := range employeesList {
		erployeeList = append(erployeeList, dto.EmpolyeList{
			EmployeeId:     emp.EmployeeUUID,
			EmployeeCode:   emp.EmployeeCode,
			Username:       emp.Username,
			FullName:       emp.FullName,
			NickName:       emp.NickName,
			BranchName:     emp.BranchName,
			DepartmentName: emp.DepartmentName,
			PositionName:   emp.PositionName,
			EmployeeStatus: emp.EmployeeStatus,
		})
	}

	return dto.GetListEmpolyeeResponse{
		Employees:      erployeeList,
		TotalEmployees: total,
		TotalPages:     int((total + int64(payload.Limit) - 1) / int64(payload.Limit)),
		Page:           payload.Page,
		Limit:          payload.Limit,
	}, nil
}

func (u *EmployeeUsecase) GetEmployeeDetail(ctx context.Context, userUUID uuid.UUID) (dto.EmployeeDetailResponse, error) {
	user, err := u.repo.GetUserByUserUUID(ctx, userUUID)
	if err != nil {
		return dto.EmployeeDetailResponse{}, err
	}

	return dto.EmployeeDetailResponse{
		EmployeeId: user.EmployeeUUID,
		Username:   user.Username,
	}, nil
}
