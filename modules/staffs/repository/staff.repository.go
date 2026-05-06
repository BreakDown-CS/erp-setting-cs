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

func (r *repository) GetStaffList(page, limit int) ([]model.Staff, int, error) {
	var staffList []model.Staff

	if page <= 0 {
		page = 1
	}

	if limit <= 0 {
		limit = 50
	}

	offset := (page - 1) * limit

	query := `
		SELECT id, employee_code, first_name, last_name, username
		FROM erp.staffs
		LIMIT $1 OFFSET $2
	`

	rows, err := r.db.Query(context.Background(), query, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	for rows.Next() {
		var staff model.Staff

		err := rows.Scan(
			&staff.ID,
			&staff.EmployeeCode,
			&staff.FirstName,
			&staff.LastName,
			&staff.Username,
		)
		if err != nil {
			return nil, 0, err
		}

		staffList = append(staffList, staff)
	}

	if err := rows.Err(); err != nil {
		return nil, 0, err
	}

	var total int
	countQuery := `SELECT COUNT(*) FROM erp.staffs`
	err = r.db.QueryRow(context.Background(), countQuery).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	return staffList, total, nil
}
