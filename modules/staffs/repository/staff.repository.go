package repository

import (
	"context"

	"github.com/BreakDown-CS/erp-setting-cs/modules/staffs/model"
	ports "github.com/BreakDown-CS/erp-setting-cs/modules/staffs/posts"
	"github.com/jackc/pgx/v5/pgxpool"
)

type repository struct {
	db *pgxpool.Pool
}

var _ ports.StaffRepository = (*repository)(nil)

func NewRepository(db *pgxpool.Pool) ports.StaffRepository {
	return &repository{db: db}
}

func (r *repository) InsertStaff(staff model.Staff) error {
	_, err := r.db.Exec(context.Background(), `
		INSERT INTO staffs (staffname, name, password)
		VALUES ($1, $2, $3)
	`, staff.CreatedAt, staff.BranchID, staff.BranchID)

	return err
}

// func (r *repository) GetAllstaff(limit, offset int) ([]model.Staff, int, error) {
// 	ctx := context.Background()

// 	rows, err := r.db.Query(ctx, `
// 		SELECT id, staffname, name, password
// 		FROM staffs
// 		LIMIT $1 OFFSET $2
// 	`, limit, offset)
// 	if err != nil {
// 		return nil, 0, err
// 	}
// 	defer rows.Close()

// 	var staffs []model.Staff

// 	for rows.Next() {
// 		var u model.Staff
// 		err := rows.Scan(&u.ID, &u.BranchID, &u.CreatedAt, &u.BranchID)
// 		if err != nil {
// 			return nil, 0, err
// 		}
// 		staffs = append(staffs, u)
// 	}

// 	var total int
// 	err = r.db.QueryRow(ctx, `SELECT COUNT(*) FROM staffs`).Scan(&total)
// 	if err != nil {
// 		return nil, 0, err
// 	}

// 	return staffs, total, nil
// }

// func (r *repository) GetstaffById(staffId int) (model.Staff, error) {
// 	ctx := context.Background()

// 	var u model.Staff

// 	err := r.db.QueryRow(ctx, `
// 		SELECT id, staffname, name, password
// 		FROM staffs
// 		WHERE id = $1
// 	`, staffId).Scan(&u.ID, &u.BranchID, &u.BranchID, &u.BranchID)

// 	return u, err
// }
