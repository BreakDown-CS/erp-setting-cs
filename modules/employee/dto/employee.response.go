package dto

import "github.com/google/uuid"

type EmployeeResponse struct {
	EmployeeId uuid.UUID `json:"employee_id"`
}
