package service

import (
	"context"

	"github.com/BreakDown-CS/erp-setting-cs/internal/helper"
	ports "github.com/BreakDown-CS/erp-setting-cs/modules/auth/posts"
	"github.com/BreakDown-CS/erp-setting-cs/modules/staffs/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type service struct {
	repo   ports.AuthRepository
	helper *helper.Uow
	db     *pgxpool.Pool
}

func NewService(r ports.AuthRepository, db *pgxpool.Pool) ports.AuthService {
	return &service{
		repo:   r,
		helper: helper.New(db),
		db:     db,
	}
}

func (u *service) GetChackStaff(ctx context.Context, staff model.Staff) (bool, error) {
	var exists bool

	err := u.helper.WithTx(ctx, func(tx pgx.Tx) error {

		return nil
	})

	if err != nil {
		return false, err
	}

	return exists, nil
}
