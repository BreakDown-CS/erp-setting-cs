package posts

import (
	"context"

	"github.com/BreakDown-CS/erp-setting-cs/modules/employee/dto"
)

type EmployeeUsecase interface {
	InsertEmployee(context.Context, dto.InsetEmpolyeeRequest) (dto.EmployeeResponse, error)
}
