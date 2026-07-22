package posts

import (
	"context"

	"github.com/BreakDown-CS/erp-setting-cs/modules/employee/dto"
	"github.com/BreakDown-CS/erp-setting-cs/modules/employee/model"
	"github.com/google/uuid"

	"github.com/jackc/pgx/v5"
)

type EmployeeRepository interface {
	BeginTransaction(context.Context) (pgx.Tx, error)
	ChackFullNameDuplicate(context.Context, pgx.Tx, string, string) (uuid.UUID, error)
	GetOneBranches(context.Context, pgx.Tx, uuid.UUID) (int, error)
	GetOneDepartment(context.Context, pgx.Tx, uuid.UUID) (int, error)
	GetOneEmployeeStatus(context.Context, pgx.Tx, uuid.UUID) (int, error)
	GetOnePosition(context.Context, pgx.Tx, uuid.UUID) (int, error)
	GetNextEmployeeCode(context.Context, pgx.Tx) (string, error)
	SaveEmployee(context.Context, pgx.Tx, model.EmployeeInsertModel) (int, error)
	GetNextUsername(context.Context, pgx.Tx) (string, error)
	SaveUser(context.Context, pgx.Tx, model.UserInsertModel) (int, error)
	GetUUIDForAll(context.Context, pgx.Tx, int, string) (uuid.UUID, error)
	GetEmployeesList(context.Context, dto.GetListEmpolyeeRequest) ([]model.GetListEmpolyeeResponse, int64, error)
	GetUserByUserUUID(context.Context, uuid.UUID) (model.GetUserByUserUUIDResponse, error)
}
