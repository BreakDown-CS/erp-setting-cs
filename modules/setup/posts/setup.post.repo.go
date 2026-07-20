package posts

import (
	"context"

	"github.com/BreakDown-CS/erp-setting-cs/modules/setup/model"
)

type SetupRepository interface {
	FindAllBranches(context.Context) ([]model.SetUpOptionUUID, error)
	FindAllDepartment(context.Context) ([]model.SetUpOptionUUID, error)
	FindAllEmployeesStatus(context.Context) ([]model.SetUpOptionUUID, error)
	FindAllPositions(context.Context) ([]model.SetUpOptionUUID, error)
}
