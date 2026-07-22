package dto

import "github.com/google/uuid"

type InsertEmpolyeeRequest struct {
	EmployeeId      uuid.UUID `json:"employee_id"`
	Username        string    `json:"username"`
	Password        string    `json:"password"`
	PasswordConfirm string    `json:"password_confirm"`
	FirstNameTh     string    `json:"first_name_th"`
	LastNameTh      string    `json:"last_name_th"`
	BranchId        uuid.UUID `json:"branch_id"`
	DepartmentId    uuid.UUID `json:"department_id"`
	PositionId      uuid.UUID `json:"position_id"`
	StatusId        uuid.UUID `json:"status_id"`
	Email           string    `json:"email"`
	Phone           string    `json:"phone"`
	EmCode          string    `json:"em_code"`
	FirstNameEn     string    `json:"first_name_en"`
	LastNameEn      string    `json:"last_name_en"`
	NickName        string    `json:"nickname"`
	Gender          string    `json:"gender"`
	Birthday        string    `json:"birthday"`
	Remark          string    `json:"remark"`
}

type GetListEmpolyeeRequest struct {
	Username     string     `json:"username"`
	EmployeeCode string     `json:"employee_code"`
	FullName     string     `json:"full_name"`
	BranchId     *uuid.UUID `json:"branch_id"`
	DepartmentId *uuid.UUID `json:"department_id"`
	StatusId     *uuid.UUID `json:"status_id"`
	Page         int        `json:"page"`
	Limit        int        `json:"limit"`
}
