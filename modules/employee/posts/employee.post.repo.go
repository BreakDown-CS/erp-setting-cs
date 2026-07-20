package posts

import (
	"context"

	"github.com/google/uuid"

	"github.com/jackc/pgx/v5"
)

type EmployeeRepository interface {
	BeginTransaction(ctx context.Context) (pgx.Tx, error)
	ChackFullNameDuplicate(ctx context.Context, tx pgx.Tx, firstName string, lastName string) (uuid.UUID, error)
	GetOneBranches(context.Context, pgx.Tx, uuid.UUID) (int, error)
	GetOneDepartment(context.Context, pgx.Tx, uuid.UUID) (int, error)
	GetOneEmployeeStatus(context.Context, pgx.Tx, uuid.UUID) (int, error)
	GetOnePosition(context.Context, pgx.Tx, uuid.UUID) (int, error)
	// SaveEmployee(ctx context.Context, tx pgx.Tx) (int, error)
}
