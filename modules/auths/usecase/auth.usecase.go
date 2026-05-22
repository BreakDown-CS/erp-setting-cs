package usecase

import (
	"context"

	"github.com/BreakDown-CS/erp-setting-cs/internal/helper"
	"github.com/BreakDown-CS/erp-setting-cs/modules/auths/dto"
	ports "github.com/BreakDown-CS/erp-setting-cs/modules/auths/posts"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

type usecase struct {
	repo   ports.AuthRepository
	helper *helper.Uow
	db     *pgxpool.Pool
}

func NewUsecase(r ports.AuthRepository, db *pgxpool.Pool) ports.AuthUsecase {
	return &usecase{
		repo:   r,
		helper: helper.New(db),
		db:     db,
	}
}

func (u *usecase) Login(ctx context.Context, req *dto.LoginRequest) (res dto.UsersLoginRes, err error) {
	user, err := u.repo.FindOneUser(req.Username)
	if err != nil {
		return dto.UsersLoginRes{}, err
	}

	if user.Id == uuid.Nil {
		return dto.UsersLoginRes{
			User: dto.DataUsersRes{
				Id: uuid.Nil,
			},
		}, nil
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password), // hash จาก DB
		[]byte(req.Password),  // password ที่ user login
	)

	if err != nil {
		return dto.UsersLoginRes{}, fiber.ErrUnauthorized
	}

	token, err := u.repo.SignUsersAccessToken(user)
	if err != nil {
		return dto.UsersLoginRes{}, err
	}

	permissions, err := u.repo.FindPermissionByUserId(*user.RoleId)
	if err != nil {
		return dto.UsersLoginRes{}, err
	}

	var permissionRes []dto.Permissions

	for _, permission := range permissions {
		permissionRes = append(permissionRes, dto.Permissions{
			PermissionId:   permission.PermissionId,
			PermissionName: permission.PermissionName,
		})
	}

	userRes := dto.DataUsersRes{
		Id:           user.Id,
		Username:     user.Username,
		BranchId:     *user.BranchId,
		DepartmentId: *user.DepartmentId,
		PositionId:   *user.PositionId,
		RoleId:       *user.RoleId,
		Permissions:  permissionRes,
	}

	data := dto.UsersLoginRes{
		User:        userRes,
		AccessToken: token,
	}

	return data, nil
}
