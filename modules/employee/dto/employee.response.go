package dto

import (
	"github.com/google/uuid"
)

type EmployeeResponse struct {
	UserId     uuid.UUID `json:"user_id"`
	EmployeeId uuid.UUID `json:"employee_id"`
}

type GetListEmpolyeeResponse struct {
	Employees      []EmpolyeList `json:"employees"`
	TotalEmployees int64         `json:"total_employees"`
	TotalPages     int           `json:"total_pages"`
	Page           int           `json:"page"`
	Limit          int           `json:"limit"`
}

type EmpolyeList struct {
	EmployeeId     uuid.UUID `json:"employee_id"`
	Username       string    `json:"username"`
	EmployeeCode   string    `json:"employee_code"`
	FullName       string    `json:"full_name"`
	NickName       string    `json:"nick_name"`
	BranchName     string    `json:"branch_name"`
	DepartmentName string    `json:"department_name"`
	PositionName   string    `json:"position_name"`
	EmployeeStatus string    `json:"employee_status"`
}

type EmployeeDetailResponse struct {
	EmployeeId uuid.UUID `json:"employee_id"`
	Username   string    `json:"username"`
}
