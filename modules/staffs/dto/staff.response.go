package dto

import "github.com/google/uuid"

type StaffSaveResponse struct {
	StaffId uuid.UUID `json:"staff_id"`
}

type StaffListResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type StaffDetailResponse struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}
