package ports

import (
	"context"

	"github.com/BreakDown-CS/erp-setting-cs/modules/staffs/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type StaffRepository interface {
	CheckDuplicate(ctx context.Context, tx pgx.Tx, staff model.Staff) (model.Staff, error)
	InsertStaff(ctx context.Context, tx pgx.Tx, staff model.Staff) (uuid.UUID, error)
}
