package ports

import (
	"context"

	"github.com/BreakDown-CS/erp-setting-cs/modules/staffs/dto"
)

type StaffService interface {
	CreateStaff(ctx context.Context, req dto.CreateStaffRequest) (dto.StaffSaveResponse, error)
	GetStaffList(page, limit int) ([]dto.StaffListResponse, int, error)
}
