package posts

import (
	"context"

	"github.com/BreakDown-CS/erp-setting-cs/modules/staffs/model"
	"github.com/jackc/pgx/v5"
)

type AuthRepository interface {
	ChackStaff(ctx context.Context, tx pgx.Tx, staff model.Staff) (bool, error)
}
