package repository

import (
	"context"

	"github.com/BreakDown-CS/erp-setting-cs/modules/staffs/model"
	ports "github.com/BreakDown-CS/erp-setting-cs/modules/staffs/posts"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type repository struct {
	db *pgxpool.Pool
}

var _ ports.StaffRepository = (*repository)(nil)

func NewRepository(db *pgxpool.Pool) ports.StaffRepository {
	return &repository{db: db}
}

func (r *repository) CheckDuplicate(ctx context.Context, tx pgx.Tx, staff model.Staff) (bool, error) {
	var staffDetail bool

	queryCheckStaff := `
		SELECT id FROM erp.staffs WHERE username = $1
	`

	err := tx.QueryRow(ctx, queryCheckStaff, staff.Username).Scan(
		&staffDetail,
	)

	if err != nil {
		return false, err
	}

	return staffDetail, nil
}

func (r *repository) InsertStaff(ctx context.Context, tx pgx.Tx, staff model.Staff) (uuid.UUID, error) {
	var id uuid.UUID

	queryInsertStaff := `
		INSERT INTO erp.staffs 
			( employee_code, first_name, last_name, username, password_hash, branch_id, department_id, position_id, status, created_by )
		VALUES
			( $1, $2, $3, $4, $5, $6, $7, $8, $9, $10 )
		RETURNING id
	`

	err := tx.QueryRow(
		ctx,
		queryInsertStaff,
		staff.EmployeeCode,
		staff.FirstName,
		staff.LastName,
		staff.Username,
		staff.PasswordHash,
		staff.BranchID,
		staff.DepartmentID,
		staff.PositionID,
		staff.Status,
		staff.CreatedBy,
	).Scan(&id)

	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}
