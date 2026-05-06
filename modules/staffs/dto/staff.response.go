package dto

import "github.com/google/uuid"

type StaffSaveResponse struct {
	StaffId uuid.UUID `json:"staff_id"`
}

type StaffListResponse struct {
	ID           string `json:"id"`
	EmployeeCode string `json:"employee_code"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Username     string `json:"username"`
}

type StaffDetailResponse struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Status    string `json:"status"`
}
