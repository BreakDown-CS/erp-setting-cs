package repository

import (
	"context"

	"github.com/BreakDown-CS/erp-setting-cs/modules/setup/model"
	posts "github.com/BreakDown-CS/erp-setting-cs/modules/setup/posts"
	"github.com/jackc/pgx/v5/pgxpool"
)

type repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) posts.SetupRepository {
	return &repository{db: db}
}

func (r *repository) FindAllBranches(ctx context.Context) ([]model.SetUpOptionUUID, error) {
	query := `
		SELECT
			bc.UUID AS id,
			bc.branch_code || '-' || bc.branch_name_th || ' (' || bc.branch_name_en || ')' AS name
		FROM
			csd.branches bc
		WHERE
			bc.is_active
	`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []model.SetUpOptionUUID
	for rows.Next() {
		var b model.SetUpOptionUUID
		if err := rows.Scan(&b.Id, &b.Name); err != nil {
			return nil, err
		}
		result = append(result, b)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func (r *repository) FindAllDepartment(ctx context.Context) ([]model.SetUpOptionUUID, error) {
	query := `
		SELECT
			dpm.UUID AS id,
			dpm.department_code || '-' || dpm.department_name_th || ' (' || dpm.department_name_en || ')' AS name
		FROM
			csd.departments dpm
		WHERE
			dpm.is_active
	`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []model.SetUpOptionUUID
	for rows.Next() {
		var b model.SetUpOptionUUID
		if err := rows.Scan(&b.Id, &b.Name); err != nil {
			return nil, err
		}
		result = append(result, b)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func (r *repository) FindAllEmployeesStatus(ctx context.Context) ([]model.SetUpOptionUUID, error) {
	query := `
		SELECT
			ems.UUID AS id,
			ems.status_name_th || ' (' || ems.status_name_en || ')' AS name
		FROM
			csd.employee_statuses ems
		WHERE
			ems.is_active
	`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []model.SetUpOptionUUID
	for rows.Next() {
		var b model.SetUpOptionUUID
		if err := rows.Scan(&b.Id, &b.Name); err != nil {
			return nil, err
		}
		result = append(result, b)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func (r *repository) FindAllPositions(ctx context.Context) ([]model.SetUpOptionUUID, error) {
	query := `
		SELECT
			pos.UUID AS id,
			pos.position_name_th || ' (' || pos.position_name_en || ')' AS name
		FROM
			csd.positions pos
		WHERE
			pos.is_active
	`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []model.SetUpOptionUUID
	for rows.Next() {
		var b model.SetUpOptionUUID
		if err := rows.Scan(&b.Id, &b.Name); err != nil {
			return nil, err
		}
		result = append(result, b)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
