package usecase

import (
	"context"
	"log"

	"github.com/BreakDown-CS/erp-setting-cs/modules/employee/dto"
	ports "github.com/BreakDown-CS/erp-setting-cs/modules/employee/posts"
)

type EmployeeUsecase struct {
	repo ports.EmployeeRepository
}

func NewUsecase(r ports.EmployeeRepository) ports.EmployeeUsecase {
	return &EmployeeUsecase{
		repo: r,
	}
}

func (u *EmployeeUsecase) InsertEmployee(ctx context.Context, payload dto.InsetEmpolyeeRequest) (dto.EmployeeResponse, error) {
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

	// employeeUUID, err := u.repo.ChackFullNameDuplicate(ctx, tx, payload.FirstNameTh, payload.LastNameTh)
	// if err != nil {
	// 	tx.Rollback(ctx)
	// 	return dto.EmployeeResponse{}, err
	// }

	// if employeeUUID != uuid.Nil {
	// 	tx.Rollback(ctx)
	// 	return dto.EmployeeResponse{
	// 		EmployeeId: employeeUUID,
	// 	}, nil
	// }

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

	log.Println("branchId = ", branchId)
	log.Println("departmentId = ", departmentId)
	log.Println("statusId = ", statusId)
	log.Println("positionId = ", positionId)

	// err = u.repo.SaveEmployee(ctx, tx)
	// if err != nil {
	//     _ = tx.Rollback(ctx)
	//     return dto.EmployeeResponse{}, err
	// }

	if err := tx.Commit(ctx); err != nil {
		return dto.EmployeeResponse{}, err
	}
	return dto.EmployeeResponse{}, nil
}
