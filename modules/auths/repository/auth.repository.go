package repository

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/BreakDown-CS/erp-setting-cs/modules/auths/model"
	posts "github.com/BreakDown-CS/erp-setting-cs/modules/auths/posts"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) posts.AuthRepository {
	return &repository{db: db}
}

func (r *repository) SignUsersAccessToken(req model.UsersPassport) (string, error) {
	claims := model.UsersClaims{
		Id:       req.Id,
		Username: req.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "access_token",
			Subject:   "users_access_token",
			ID:        uuid.NewString(),
			Audience:  []string{"users"},
		},
	}

	mySigningKey := os.Getenv("JWT_SECRET_KEY")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(mySigningKey))
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	return ss, nil
}

func (r *repository) FindOneUser(username string) (model.UsersPassport, error) {

	query := `
		SELECT
			s.id id,
			s.username username,
			s.password_hash,
			bc.id branch_id,
			dpm.id department_id,
			pst.id position_id,
			r.id role_id
		FROM
			erp.staffs s
			LEFT JOIN erp.branches bc ON s.branch_id = bc.id 
			LEFT JOIN erp.departments dpm ON s.department_id = dpm.id
			LEFT JOIN erp.positions pst ON s.position_id = pst.id
			LEFT JOIN erp.staff_roles sr ON s.id = sr.staff_id
			LEFT JOIN erp.roles r ON sr.role_id = r.id
		WHERE
			s.username = $1
	`

	var user model.UsersPassport

	err := r.db.QueryRow(
		context.Background(),
		query,
		username,
	).Scan(
		&user.Id,
		&user.Username,
		&user.Password,
		&user.BranchId,
		&user.DepartmentId,
		&user.PositionId,
		&user.RoleId,
	)

	if err != nil {
		return model.UsersPassport{}, err
	}

	return user, nil
}

func (r *repository) FindPermissionByUserId(roleId uuid.UUID) ([]model.Permissions, error) {

	permissions := []model.Permissions{}

	query := `
		SELECT
			pmt.id permission_id,
			pmt.code || ' (' || pmt.name || ')' permission_name
		FROM
			erp.roles r
			LEFT JOIN erp.role_permissions rpmt ON rpmt.role_id = r.id
			LEFT JOIN erp.permissions pmt ON rpmt.permission_id = pmt.id
		WHERE
			r.id = $1
	`

	rows, err := r.db.Query(
		context.Background(),
		query,
		roleId,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {

		var permission model.Permissions

		err := rows.Scan(
			&permission.PermissionId,
			&permission.PermissionName,
		)

		if err != nil {
			return nil, err
		}

		permissions = append(permissions, permission)
	}

	return permissions, nil
}
