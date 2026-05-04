package ports

import "github.com/BreakDown-CS/erp-setting-cs/modules/staffs/dto"

type StaffService interface {
	InsertStaff(req dto.CreateStaffRequest) error
}
