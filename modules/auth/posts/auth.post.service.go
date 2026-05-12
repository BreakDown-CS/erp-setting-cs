package posts

import (
	"context"

	"github.com/BreakDown-CS/erp-setting-cs/modules/staffs/model"
)

type AuthService interface {
	GetChackStaff(ctx context.Context, staff model.Staff) (bool, error)
}
