package dto

import "github.com/google/uuid"

type UsersLoginRes struct {
	User        DataUsersRes `json:"user"`
	AccessToken string       `json:"access_token"`
}

type DataUsersRes struct {
	Id           uuid.UUID     `json:"id"`
	Username     string        `json:"username"`
	BranchId     uuid.UUID     `json:"branch_id"`
	DepartmentId uuid.UUID     `json:"department_id"`
	PositionId   uuid.UUID     `json:"position_id"`
	RoleId       uuid.UUID     `json:"role_id"`
	Permissions  []Permissions `json:"permissions"`
}

type Permissions struct {
	PermissionId   uuid.UUID `json:"permission_id"`
	PermissionName string    `json:"permission_name"`
}
