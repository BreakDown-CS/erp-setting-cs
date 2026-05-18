package ports

import (
	"context"

	"github.com/BreakDown-CS/erp-setting-cs/modules/auths/dto"
)

type AuthUsecase interface {
	Login(ctx context.Context, req *dto.LoginRequest) (dto.UsersLoginRes, error)
}
