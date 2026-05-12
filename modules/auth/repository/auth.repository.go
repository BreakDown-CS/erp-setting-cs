package repository

import (
	"context"

	ports "github.com/BreakDown-CS/erp-setting-cs/modules/auth/posts"
	"github.com/BreakDown-CS/erp-setting-cs/modules/staffs/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) ports.AuthRepository {
	return &repository{db: db}
}

func (r *repository) ChackStaff(ctx context.Context, tx pgx.Tx, staff model.Staff) (bool, error) {
	var staffDetail bool

	queryCheckStaff := `
		SELECT EXISTS (
			SELECT 1 FROM erp.staffs WHERE username = $1
		)
	`
	err := tx.QueryRow(ctx, queryCheckStaff, staff.Username).Scan(
		&staffDetail,
	)

	if err != nil {
		return false, err
	}

	return staffDetail, nil
}
