package repository

import (
	"context"
	"errors"
	"fmt"

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
