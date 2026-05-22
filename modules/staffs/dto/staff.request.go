package dto

import "github.com/google/uuid"

type CreateStaffRequest struct {
	EmployeeCode string `json:"employee_code" validate:"required,min=3"`
	FirstName    string `json:"first_name" validate:"required"`
	LastName     string `json:"last_name" validate:"required"`
	Username     string `json:"username" validate:"required,min=4"`
	Password     string `json:"password" validate:"required,min=6"`
	BranchId     string `json:"branch_id" validate:"omitempty,uuid"`
	DepartmentId string `json:"department_id" validate:"omitempty,uuid"`
	PositionId   string `json:"position_id" validate:"omitempty,uuid"`
	CreatedBy    string `json:"created_by" validate:"omitempty,uuid"`
}

type GetStaffById struct {
	ID uuid.UUID `json:"id"`
}

type GetStaffListRequest struct {
	Username     string `json:"username"`
	EmCode       string `json:"em_code"`
	BranchesId   string `json:"branches_id"`
	FullName     string `json:"full_name"`
	Status       string `json:"status"`
	DepartmentId string `json:"department_id"`
	Page         int    `json:"page"`
	Limit        int    `json:"limit"`
}
