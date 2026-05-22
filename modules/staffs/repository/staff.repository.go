package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/BreakDown-CS/erp-setting-cs/modules/staffs/dto"
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

func (r *repository) GetStaffList(req dto.GetStaffListRequest) ([]model.StaffList, int, error) {
	var staffList []model.StaffList

	if req.Page <= 0 {
		req.Page = 1
	}

	if req.Limit <= 0 {
		req.Limit = 50
	}

	offset := (req.Page - 1) * req.Limit

	conditions := []string{}
	args := []interface{}{}
	argIndex := 1

	if req.Username != "" {
		conditions = append(conditions,
			fmt.Sprintf("s.username = $%d", argIndex))
		args = append(args, req.Username)
		argIndex++
	}

	if req.EmCode != "" {
		conditions = append(conditions,
			fmt.Sprintf("s.employee_code = $%d", argIndex))
		args = append(args, req.EmCode)
		argIndex++
	}

	if req.BranchesId != "" {
		conditions = append(conditions,
			fmt.Sprintf("s.branch_id = $%d", argIndex))
		args = append(args, req.BranchesId)
		argIndex++
	}

	if req.FullName != "" {
		conditions = append(conditions,
			fmt.Sprintf("(s.first_name || ' ' || s.last_name) ILIKE $%d", argIndex))
		args = append(args, fmt.Sprintf("%%%s%%", req.FullName))
		argIndex++
	}

	if req.Status != "" {
		conditions = append(conditions,
			fmt.Sprintf("s.status = $%d", argIndex))
		args = append(args, req.Status)
		argIndex++
	}

	if req.DepartmentId != "" {
		conditions = append(conditions,
			fmt.Sprintf("s.department_id = $%d", argIndex))
		args = append(args, req.DepartmentId)
		argIndex++
	}

	whereQuery := ""
	if len(conditions) > 0 {
		whereQuery = "WHERE " + strings.Join(conditions, " AND ")
	}

	limitArg := argIndex
	offsetArg := argIndex + 1

	args = append(args, req.Limit, offset)

	query := fmt.Sprintf(`
		SELECT
			s.id,
			s.username,
			s.employee_code,
			s.first_name,
			s.last_name,
			bc.code || ' (' || bc.name || ')' AS branches_name,
			dpm.name AS department_name,
			pst.name AS position_name,
			s.status
		FROM erp.staffs s
			LEFT JOIN erp.branches bc ON s.branch_id = bc.id
			LEFT JOIN erp.departments dpm ON s.department_id = dpm.id
			LEFT JOIN erp.positions pst ON s.position_id = pst.id
		%s
		ORDER BY s.id ASC
		LIMIT $%d OFFSET $%d
	`, whereQuery, limitArg, offsetArg)

	rows, err := r.db.Query(context.Background(), query, args...)
	if err != nil {
		fmt.Println("query err:", err)
		return nil, 0, err
	}
	defer rows.Close()

	for rows.Next() {
		var staff model.StaffList

		err := rows.Scan(
			&staff.ID,
			&staff.Username,
			&staff.EmployeeCode,
			&staff.FirstName,
			&staff.LastName,
			&staff.BranchName,
			&staff.DepartmentName,
			&staff.PositionName,
			&staff.Status,
		)
		if err != nil {
			return nil, 0, err
		}

		staffList = append(staffList, staff)
	}

	if err := rows.Err(); err != nil {
		return nil, 0, err
	}

	// total count
	countQuery := fmt.Sprintf(`SELECT COUNT(*) FROM erp.staffs s %s`, whereQuery)

	var total int

	err = r.db.QueryRow(
		context.Background(),
		countQuery,
		args[:len(args)-2]..., // ตัด limit offset ออก
	).Scan(&total)

	if err != nil {
		return nil, 0, err
	}

	return staffList, total, nil
}

func (r *repository) GetStaffById(id uuid.UUID) (model.Staff, error) {
	var staff model.Staff

	query := `
		SELECT id, employee_code, first_name, last_name, username, status
		FROM erp.staffs
		WHERE id = $1
	`

	err := r.db.QueryRow(context.Background(), query, id).Scan(
		&staff.ID,
		&staff.EmployeeCode,
		&staff.FirstName,
		&staff.LastName,
		&staff.Username,
		&staff.Status,
	)

	return staff, err
}
