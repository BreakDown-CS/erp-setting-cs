package dto

type CreateStaffRequest struct {
	Staffname string `json:"staff_name"`
	Name      string `json:"name"`
	Password  string `json:"password"`
}
