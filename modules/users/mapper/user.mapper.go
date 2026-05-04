package mapper

import (
	"github.com/BreakDown-CS/erp-setting-cs/modules/users/dto"
	"github.com/BreakDown-CS/erp-setting-cs/modules/users/model"
)

func ToUserList(u model.User) dto.UserListResponse {
	return dto.UserListResponse{
		ID:       u.ID,
		Username: u.Username,
		Name:     u.Name,
	}
}

func ToUserDetail(u model.User) dto.UserDetailResponse {
	return dto.UserDetailResponse{
		ID:       u.ID,
		Username: u.Username,
		Name:     u.Name,
	}
}
