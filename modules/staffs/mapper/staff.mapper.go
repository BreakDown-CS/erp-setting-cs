package mapper

import (
	"github.com/BreakDown-CS/erp-setting-cs/modules/staffs/dto"
	"github.com/BreakDown-CS/erp-setting-cs/modules/staffs/model"
)

func ToStaffList(u model.StaffList) dto.StaffListResponse {
	return dto.StaffListResponse{
		ID:             u.ID.String(),
		EmployeeCode:   u.EmployeeCode,
		FirstName:      u.FirstName,
		LastName:       u.LastName,
		Username:       u.Username,
		BranchName:     u.BranchName,
		DepartmentName: u.DepartmentName,
		PositionName:   u.PositionName,
		Status:         u.Status,
	}
}

func ToStaffDetail(u model.Staff) dto.StaffDetailResponse {
	return dto.StaffDetailResponse{
		ID:        u.ID.String(),
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Username:  u.Username,
		Status:    u.Status,
	}
}
