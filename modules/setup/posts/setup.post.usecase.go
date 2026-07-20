package posts

import (
	"context"

	"github.com/BreakDown-CS/erp-setting-cs/modules/setup/dto"
)

type SetupUsecase interface {
	GetAllBranches(ctx context.Context) ([]dto.SetUpBool, error)
	GetAllDepartment(ctx context.Context) ([]dto.SetUpBool, error)
	GetAllEmployeesStatus(ctx context.Context) ([]dto.SetUpBool, error)
	GetAllPositions(ctx context.Context) ([]dto.SetUpBool, error)
}
