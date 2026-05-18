package model

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type UsersPassport struct {
	Id           uuid.UUID  `json:"id" db:"id"`
	Username     string     `json:"username" db:"username"`
	Password     string     `json:"password" db:"password"`
	BranchId     *uuid.UUID `json:"branch_id" db:"branch_id"`
	DepartmentId *uuid.UUID `json:"department_id" db:"department_id"`
	PositionId   *uuid.UUID `json:"position_id" db:"position_id"`
	RoleId       *uuid.UUID `json:"role_id" db:"role_id"`
}

type UsersClaims struct {
	Id       uuid.UUID `json:"user_id"`
	Username string    `json:"username"`
	jwt.RegisteredClaims
}

type Permissions struct {
	PermissionId   uuid.UUID `json:"permission_id" db:"permission_id"`
	PermissionName string    `json:"permission_name" db:"permission_name"`
}
