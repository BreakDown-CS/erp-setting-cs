package dto

type CreateStaffRequest struct {
	EmployeeCode string `json:"employee_code"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	BranchId     string `json:"branch_id"`
	DepartmentId string `json:"department_id"`
	PositionId   string `json:"position_id"`
	CreatedBy    string `json:"created_by"`
}
