package ports

import (
	"context"

	"github.com/BreakDown-CS/erp-setting-cs/modules/auths/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type AuthRepository interface {
	SignUsersAccessToken(model.UsersPassport) (string, error)
	FindOneUser(string) (model.UsersPassport, error)
	FindPermissionByUserId(uuid.UUID) ([]model.Permissions, error)
}

type EmployeeRepository interface {
	BeginTransaction(
		ctx context.Context,
	) (pgx.Tx, error)

	ChackFullNameDuplicate(
		ctx context.Context,
		tx pgx.Tx,
		firstName string,
		lastName string,
	) (int, error)

	SaveEmployee(
		ctx context.Context,
		tx pgx.Tx,
	) error
}

