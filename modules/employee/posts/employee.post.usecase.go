package posts

import (
	"context"

	"github.com/BreakDown-CS/erp-setting-cs/modules/employee/dto"
	"github.com/google/uuid"
)

type EmployeeUsecase interface {
	InsertEmployee(context.Context, dto.InsertEmpolyeeRequest) (dto.EmployeeResponse, error)
	GetEmployeeList(context.Context, dto.GetListEmpolyeeRequest) (dto.GetListEmpolyeeResponse, error)
	GetEmployeeDetail(context.Context, uuid.UUID) (dto.EmployeeDetailResponse, error)
}
