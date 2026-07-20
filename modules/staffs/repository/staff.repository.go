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

func (r *repository) CheckDuplicate(ctx context.Context, tx pgx.Tx, username string, employeeCode string) (bool, error) {
	var exists bool

	query := `
		SELECT EXISTS (
			SELECT 1 FROM erp.staffs 
			WHERE (username = $1 OR employee_code = $2) AND deleted_at IS NULL
		)
	`
	err := tx.QueryRow(ctx, query, username, employeeCode).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("check duplicate failed: %w", err)
	}

	return exists, nil
}

func (r *repository) InsertStaff(ctx context.Context, tx pgx.Tx, staff model.Staff) (uuid.UUID, error) {
	var id uuid.UUID

	query := `
		INSERT INTO erp.staffs 
			(employee_code, first_name, last_name, username, password_hash, 
			 branch_id, department_id, position_id, status, created_by)
		VALUES
			($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		RETURNING id
	`

	err := tx.QueryRow(ctx, query,
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
		return uuid.Nil, fmt.Errorf("insert staff failed: %w", err)
	}

	return id, nil
}

// queryCondition represents a SQL WHERE condition
type queryCondition struct {
	condition string      // SQL condition template, use $%d for parameter placeholder
	arg       interface{} // argument value, nil for conditions without parameters
	hasParam  bool        // whether this condition has a parameter
}

// buildWhereClause builds a WHERE clause from a list of conditions
func buildWhereClause(conditions []queryCondition) (string, []interface{}) {
	if len(conditions) == 0 {
		return "", nil
	}

	var clauses []string
	var args []interface{}
	paramIdx := 1

	for _, cond := range conditions {
		if cond.hasParam {
			clauses = append(clauses, fmt.Sprintf(cond.condition, paramIdx))
			args = append(args, cond.arg)
			paramIdx++
		} else {
			clauses = append(clauses, cond.condition)
		}
	}

	return "WHERE " + strings.Join(clauses, " AND "), args
}

func (r *repository) GetStaffList(ctx context.Context, req dto.GetStaffListRequest) ([]model.StaffList, int, error) {
	// Default pagination
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.Limit <= 0 || req.Limit > 100 {
		req.Limit = 50
	}

	offset := (req.Page - 1) * req.Limit

	var conditions []queryCondition

	if req.Username != "" {
		conditions = append(conditions, queryCondition{
			condition: "s.username = $%d",
			arg:       req.Username,
		})
	}

	if req.EmployeeCode != "" {
		conditions = append(conditions, queryCondition{
			condition: "s.employee_code = $%d",
			arg:       req.EmployeeCode,
		})
	}

	if req.BranchId != "" {
		conditions = append(conditions, queryCondition{
			condition: "s.branch_id = $%d",
			arg:       req.BranchId,
		})
	}

	if req.FullName != "" {
		conditions = append(conditions, queryCondition{
			condition: "(s.first_name || ' ' || s.last_name) ILIKE $%d",
			arg:       fmt.Sprintf("%%%s%%", req.FullName),
		})
	}

	if req.Status != "" {
		conditions = append(conditions, queryCondition{
			condition: "s.status = $%d",
			arg:       req.Status,
		})
	}

	if req.DepartmentId != "" {
		conditions = append(conditions, queryCondition{
			condition: "s.department_id = $%d",
			arg:       req.DepartmentId,
		})
	}

	// Always filter soft-deleted records
	conditions = append(conditions, queryCondition{
		condition: "s.deleted_at IS NULL",
		arg:       nil, // placeholder only, won't be used
	})

	whereClause, whereArgs := buildWhereClause(conditions)

	// Build select query
	selectQuery := fmt.Sprintf(`
		SELECT
			s.id,
			s.username,
			s.employee_code,
			s.first_name,
			s.last_name,
			bc.code || ' (' || bc.name || ')' AS branch_name,
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
	`, whereClause, len(whereArgs)+1, len(whereArgs)+2)

	queryArgs := append(whereArgs, req.Limit, offset)

	rows, err := r.db.Query(ctx, selectQuery, queryArgs...)
	if err != nil {
		return nil, 0, fmt.Errorf("query staff list failed: %w", err)
	}
	defer rows.Close()

	var staffList []model.StaffList
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
			return nil, 0, fmt.Errorf("scan staff row failed: %w", err)
		}
		staffList = append(staffList, staff)
	}

	if err := rows.Err(); err != nil {
		return nil, 0, fmt.Errorf("rows iteration failed: %w", err)
	}

	// Total count query
	countQuery := fmt.Sprintf(`SELECT COUNT(*) FROM erp.staffs s %s`, whereClause)
	var total int
	err = r.db.QueryRow(ctx, countQuery, whereArgs...).Scan(&total)
	if err != nil {
		return nil, 0, fmt.Errorf("count staff failed: %w", err)
	}

	return staffList, total, nil
}

func (r *repository) GetStaffById(ctx context.Context, id uuid.UUID) (model.Staff, error) {
	var staff model.Staff

	query := `
		SELECT id, employee_code, first_name, last_name, username, status
		FROM erp.staffs
		WHERE id = $1 AND deleted_at IS NULL
	`

	err := r.db.QueryRow(ctx, query, id).Scan(
		&staff.ID,
		&staff.EmployeeCode,
		&staff.FirstName,
		&staff.LastName,
		&staff.Username,
		&staff.Status,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return staff, fmt.Errorf("staff not found: %w", err)
		}
		return staff, fmt.Errorf("get staff by id failed: %w", err)
	}

	return staff, nil
}
