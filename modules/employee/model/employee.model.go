package model

import (
	"time"

	"github.com/google/uuid"
)

type EmployeeInsertModel struct {
	EmployeeCode string    `json:"employee_code" db:"employee_code"`
	FirstNameTh  string    `json:"first_name_th" db:"first_name_th"`
	LastNameTh   string    `json:"last_name_th" db:"last_name_th"`
	FirstNameEn  string    `json:"first_name_en" db:"first_name_en"`
	LastNameEn   string    `json:"last_name_en" db:"last_name_en"`
	NickName     string    `json:"nickname" db:"nickname"`
	Gender       string    `json:"gender" db:"gender"`
	Birthday     time.Time `json:"birthday" db:"birthday"`
	Email        string    `json:"email" db:"email"`
	Phone        string    `json:"phone" db:"phone"`
	BranchId     int       `json:"branch_id" db:"branch_id"`
	DepartmentId int       `json:"department_id" db:"department_id"`
	PositionId   int       `json:"position_id" db:"position_id"`
	StatusId     int       `json:"status_id" db:"status_id"`
	Remark       string    `json:"remark" db:"remark"`
}

type UserInsertModel struct {
	Username     string `json:"username" db:"username"`
	PasswordHash string `json:"password_hash" db:"password_hash"`
	EmployeeId   int    `json:"employee_id" db:"employee_id"`
}

type GetListEmpolyeeResponse struct {
	EmployeeUUID   uuid.UUID `json:"employee_uuid"`
	Username       string    `json:"username"`
	EmployeeCode   string    `json:"employee_code"`
	FullName       string    `json:"full_name"`
	NickName       string    `json:"nick_name"`
	BranchName     string    `json:"branch_name"`
	DepartmentName string    `json:"department_name"`
	PositionName   string    `json:"position_name"`
	EmployeeStatus string    `json:"employee_status"`
}

type GetUserByUserUUIDResponse struct {
	EmployeeUUID uuid.UUID `json:"employee_uuid"`
	Username     string    `json:"username"`
}
