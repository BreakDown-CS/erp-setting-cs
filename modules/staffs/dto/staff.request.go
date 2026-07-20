package dto

import "github.com/google/uuid"

type CreateStaffRequest struct {
	EmployeeCode string `json:"employee_code" validate:"required,min=3,max=50"`
	FirstName    string `json:"first_name" validate:"required,min=1,max=100"`
	LastName     string `json:"last_name" validate:"required,min=1,max=100"`
	Username     string `json:"username" validate:"required,min=4,max=50"`
	Password     string `json:"password" validate:"required,min=6,max=100"`
	BranchId     string `json:"branch_id" validate:"omitempty,uuid"`
	DepartmentId string `json:"department_id" validate:"omitempty,uuid"`
	PositionId   string `json:"position_id" validate:"omitempty,uuid"`
	CreatedBy    string `json:"created_by" validate:"omitempty,uuid"`
}

type GetStaffById struct {
	ID uuid.UUID `json:"id" validate:"required"`
}

type GetStaffListRequest struct {
	Username     string `json:"username"`
	EmployeeCode string `json:"employee_code"`
	BranchId     string `json:"branch_id"`
	FullName     string `json:"full_name"`
	Status       string `json:"status"`
	DepartmentId string `json:"department_id"`
	Page         int    `json:"page" validate:"omitempty,min=1"`
	Limit        int    `json:"limit" validate:"omitempty,min=1,max=100"`
}
