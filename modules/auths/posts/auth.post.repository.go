package ports

import (
	"github.com/BreakDown-CS/erp-setting-cs/modules/auths/model"
	"github.com/google/uuid"
)

type AuthRepository interface {
	SignUsersAccessToken(model.UsersPassport) (string, error)
	FindOneUser(string) (model.UsersPassport, error)
	FindPermissionByUserId(uuid.UUID) ([]model.Permissions, error)
}
