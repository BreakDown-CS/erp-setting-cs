package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/BreakDown-CS/erp-setting-cs/modules/employee/dto"
	"github.com/BreakDown-CS/erp-setting-cs/modules/employee/model"
	posts "github.com/BreakDown-CS/erp-setting-cs/modules/employee/posts"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type EmployeeRepository interface {
	BeginTransaction(ctx context.Context) (pgx.Tx, error)

	SaveEmployee(
		ctx context.Context,
		tx pgx.Tx,
	) (int, error)
}

type employeeRepository struct {
	DB *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) posts.EmployeeRepository {
	return &employeeRepository{
		DB: db,
	}
}

func (r *employeeRepository) BeginTransaction(ctx context.Context) (pgx.Tx, error) {
	return r.DB.Begin(ctx)
}

func (r *employeeRepository) ChackFullNameDuplicate(ctx context.Context, tx pgx.Tx, firstName, lastName string) (uuid.UUID, error) {

	query := `
		SELECT
			emp.id
		FROM csd.employees emp
		WHERE emp.first_name_th = $1
		AND emp.last_name_th = $2
	`

	var result uuid.UUID

	err := tx.QueryRow(
		ctx,
		query,
		firstName,
		lastName,
	).Scan(&result)

	if err != nil {
		return uuid.Nil, err
	}

	return result, nil
}

func (r *employeeRepository) GetOneBranches(ctx context.Context, tx pgx.Tx, branchUUID uuid.UUID) (int, error) {

	query := `
		SELECT id
		FROM csd.branches
		WHERE
			is_active = true
			AND uuid = $1
	`

	var id int

	err := tx.QueryRow(
		ctx,
		query,
		branchUUID,
	).Scan(&id)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return 0, fmt.Errorf("branch not found")
		}

		return 0, err
	}

	return id, nil
}

func (r *employeeRepository) GetOneDepartment(ctx context.Context, tx pgx.Tx, branchUUID uuid.UUID) (int, error) {

	query := `
		SELECT id
		FROM csd.departments
		WHERE
			is_active = true
			AND uuid = $1
	`

	var id int

	err := tx.QueryRow(
		ctx,
		query,
		branchUUID,
	).Scan(&id)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return 0, fmt.Errorf("branch not found")
		}

		return 0, err
	}

	return id, nil
}

func (r *employeeRepository) GetOneEmployeeStatus(ctx context.Context, tx pgx.Tx, branchUUID uuid.UUID) (int, error) {

	query := `
		SELECT id
		FROM csd.employee_statuses
		WHERE
			is_active = true
			AND uuid = $1
	`

	var id int

	err := tx.QueryRow(
		ctx,
		query,
		branchUUID,
	).Scan(&id)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return 0, fmt.Errorf("branch not found")
		}

		return 0, err
	}

	return id, nil
}

func (r *employeeRepository) GetOnePosition(ctx context.Context, tx pgx.Tx, branchUUID uuid.UUID) (int, error) {

	query := `
		SELECT id
		FROM csd.positions
		WHERE
			is_active = true
			AND uuid = $1
	`

	var id int

	err := tx.QueryRow(
		ctx,
		query,
		branchUUID,
	).Scan(&id)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return 0, fmt.Errorf("branch not found")
		}

		return 0, err
	}

	return id, nil
}

func (r *employeeRepository) GetNextEmployeeCode(ctx context.Context, tx pgx.Tx) (string, error) {

	query := `
		SELECT
			'EMP' || LPAD((COALESCE(MAX(SUBSTRING(employee_code FROM 4)::INT), 0) + 1)::TEXT, 6, '0') AS employee_code
		FROM
			csd.employees
	`

	var result string

	err := tx.QueryRow(ctx, query).Scan(&result)
	if err != nil {
		return "", fmt.Errorf("get next employee code: %w", err)
	}

	return result, nil
}

func (r *employeeRepository) SaveEmployee(ctx context.Context, tx pgx.Tx, modelSave model.EmployeeInsertModel) (int, error) {

	query := `
		INSERT INTO csd.employees
		(
				employee_code,
				first_name_th,
				last_name_th,
				nickname,
				gender,
				birth_date,
				email,
				phone,
				branch_id,
				department_id,
				position_id,
				status_id,
				hire_date
		)
		VALUES
		(
				$1,
				$2,
				$3,
				$4,
				$5,
				$6,
				$7,
				$8,
				$9,
				$10,
				$11,
				$12,
				CURRENT_DATE
		)
		RETURNING id
	`

	var result int

	err := tx.QueryRow(
		ctx,
		query,
		modelSave.EmployeeCode,
		modelSave.FirstNameTh,
		modelSave.LastNameTh,
		modelSave.NickName,
		modelSave.Gender,
		modelSave.Birthday,
		modelSave.Email,
		modelSave.Phone,
		modelSave.BranchId,
		modelSave.DepartmentId,
		modelSave.PositionId,
		modelSave.StatusId,
	).Scan(&result)

	if err != nil {
		return 0, fmt.Errorf("insert employee failed: %w", err)
	}

	return result, nil
}

func (r *employeeRepository) GetNextUsername(ctx context.Context, tx pgx.Tx) (string, error) {

	query := `
		SELECT
			'USN' || LPAD((COALESCE(MAX(SUBSTRING(username FROM 4)::INT), 0) + 1)::TEXT, 6, '0') AS username
		FROM
			csd.users
		WHERE
			username ~ '^USN[0-9]+$'
	`

	var result string

	err := tx.QueryRow(ctx, query).Scan(&result)
	if err != nil {
		return "", fmt.Errorf("get next employee code: %w", err)
	}

	return result, nil
}

func (r *employeeRepository) SaveUser(ctx context.Context, tx pgx.Tx, modelSave model.UserInsertModel) (int, error) {

	query := `
		INSERT INTO csd.users
		(
				employee_id,
				username,
				password_hash
		)
		VALUES
		(
				$1,
				$2,
				$3
		)
		RETURNING id
	`

	var result int

	err := tx.QueryRow(
		ctx,
		query,
		modelSave.EmployeeId,
		modelSave.Username,
		modelSave.PasswordHash,
	).Scan(&result)

	if err != nil {
		return 0, fmt.Errorf("insert user failed: %w", err)
	}

	return result, nil
}

func (r *employeeRepository) GetUUIDForAll(ctx context.Context, tx pgx.Tx, id int, typeId string) (uuid.UUID, error) {

	query := ""

	switch typeId {
	case "user":
		query = "SELECT us.uuid uuid FROM csd.users us WHERE us.id = $1"
	case "employee":
		query = "SELECT emp.uuid uuid FROM csd.employees emp WHERE emp.id = $1"
	}

	var result uuid.UUID

	err := tx.QueryRow(ctx, query, id).Scan(&result)
	if err != nil {
		return uuid.Nil, fmt.Errorf("get uuid for %s failed: %w", typeId, err)
	}

	return result, nil
}

func (r *employeeRepository) GetEmployeesList(ctx context.Context, payload dto.GetListEmpolyeeRequest) ([]model.GetListEmpolyeeResponse, int64, error) {

	if payload.Page <= 0 {
		payload.Page = 1
	}

	if payload.Limit <= 0 {
		payload.Limit = 50
	}

	querySQL := `
		SELECT
			emp.uuid AS employee_uuid,
			us.username,
			emp.employee_code,
			emp.first_name_th || ' ' || emp.last_name_th AS full_name,
			emp.nickname nick_name,
			bc.branch_name_th || '(' || bc.branch_code || '-' || bc.branch_name_en || ')' AS branch_name,
			dpm.department_name_th || '(' || dpm.department_code || '-' || dpm.department_name_en || ')' AS department_name,
			pst.position_name_th || '(' || pst.position_code || '-' || pst.position_name_en || ')' AS position_name,
			emps.status_name_th AS employee_status
		FROM csd.employees emp
			LEFT JOIN csd.users us ON us.employee_id = emp.id
			LEFT JOIN csd.branches bc	ON emp.branch_id = bc.id
			LEFT JOIN csd.departments dpm	ON emp.department_id = dpm.id
			LEFT JOIN csd.positions pst ON emp.position_id = pst.id
			LEFT JOIN csd.employee_statuses emps ON emp.status_id = emps.id
	`

	whereQuery := ` WHERE 1=1`

	filterArgs := make([]any, 0)
	args := make([]any, 0)

	index := 1

	if payload.BranchId != nil {
		whereQuery += fmt.Sprintf(" AND bc.uuid = $%d", index)
		filterArgs = append(filterArgs, payload.BranchId)
		args = append(args, payload.BranchId)
		index++
	}

	if payload.DepartmentId != nil {
		whereQuery += fmt.Sprintf(" AND dpm.uuid = $%d", index)
		filterArgs = append(filterArgs, payload.DepartmentId)
		args = append(args, payload.DepartmentId)
		index++
	}

	if payload.StatusId != nil {
		whereQuery += fmt.Sprintf(" AND emps.uuid = $%d", index)
		filterArgs = append(filterArgs, payload.StatusId)
		args = append(args, payload.StatusId)
		index++
	}

	orderQuery := `
		ORDER BY emp.id ASC
	`

	limitQuery := fmt.Sprintf(" LIMIT $%d", index)
	args = append(args, payload.Limit)
	index++

	offset := (payload.Page - 1) * payload.Limit
	offsetQuery := fmt.Sprintf(" OFFSET $%d", index)
	args = append(args, offset)

	query := querySQL +
		whereQuery +
		orderQuery +
		limitQuery +
		offsetQuery

	totalQuery := `
		SELECT 
			COUNT(*) 
		FROM csd.employees emp
			LEFT JOIN csd.users us ON us.employee_id = emp.id
			LEFT JOIN csd.branches bc	ON emp.branch_id = bc.id
			LEFT JOIN csd.departments dpm	ON emp.department_id = dpm.id
			LEFT JOIN csd.positions pst ON emp.position_id = pst.id
			LEFT JOIN csd.employee_statuses emps ON emp.status_id = emps.id
		` + whereQuery

	rows, err := r.DB.Query(ctx, query, args...)
	if err != nil {
		return nil, 0, fmt.Errorf("get employees list failed: %w", err)
	}
	defer rows.Close()

	result := make([]model.GetListEmpolyeeResponse, 0)

	for rows.Next() {
		var emp model.GetListEmpolyeeResponse

		err := rows.Scan(
			&emp.EmployeeUUID,
			&emp.Username,
			&emp.EmployeeCode,
			&emp.FullName,
			&emp.NickName,
			&emp.BranchName,
			&emp.DepartmentName,
			&emp.PositionName,
			&emp.EmployeeStatus,
		)
		if err != nil {
			return nil, 0, fmt.Errorf("scan employee failed: %w", err)
		}

		result = append(result, emp)
	}

	if err := rows.Err(); err != nil {
		return nil, 0, fmt.Errorf("rows iteration failed: %w", err)
	}

	var total int64

	err = r.DB.QueryRow(ctx, totalQuery, filterArgs...).Scan(&total)
	if err != nil {
		return nil, 0, fmt.Errorf("get employees count failed: %w", err)
	}

	return result, total, nil
}

func (r *employeeRepository) GetUserByUserUUID(ctx context.Context, userUUID uuid.UUID) (model.GetUserByUserUUIDResponse, error) {

	query := `
		SELECT
			emp.uuid employee_uuid,
			us.username
		FROM
			csd.employees emp
			LEFT JOIN csd.users us ON us.employee_id = emp.id
		WHERE
			emp.uuid = $1
	`

	var result model.GetUserByUserUUIDResponse

	err := r.DB.QueryRow(ctx, query, userUUID).Scan(
		&result.EmployeeUUID,
		&result.Username,
	)
	if err != nil {
		return model.GetUserByUserUUIDResponse{}, fmt.Errorf("get user by user uuid failed: %w", err)
	}

	return result, nil
}
