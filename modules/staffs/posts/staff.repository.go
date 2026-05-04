package ports

import "github.com/BreakDown-CS/erp-setting-cs/modules/staffs/model"

type StaffRepository interface {
	InsertStaff(staff model.Staff) error
}
