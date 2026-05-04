package service

import (
	"users-api/modules/users/dto"
	"users-api/modules/users/mapper"
	"users-api/modules/users/model"
	"users-api/modules/users/repository"
)

type Usecase interface {
	InsetUser(req dto.CreateUserRequest) error
	GetAllUsers(page, limit int) ([]dto.UserListResponse, int, error)
	GetUserById(id int) (dto.UserDetailResponse, error)
}

type usecase struct {
	repo repository.Repository
}

func NewUsecase(repo repository.Repository) Usecase {
	return &usecase{repo: repo}
}

func (u *usecase) InsetUser(req dto.CreateUserRequest) error {
	user := model.User{
		Username: req.Username,
		Name:     req.Name,
		Password: req.Password,
	}
	return u.repo.InsetUser(user)
}

func (u *usecase) GetAllUsers(page int, limit int) ([]dto.UserListResponse, int, error) {

	offset := (page - 1) * limit

	users, total, err := u.repo.GetAllUser(limit, offset)
	if err != nil {
		return nil, 0, err
	}

	res := make([]dto.UserListResponse, 0)

	for _, v := range users {
		res = append(res, mapper.ToUserList(v))
	}

	return res, total, nil
}

func (u *usecase) GetUserById(id int) (dto.UserDetailResponse, error) {
	user, err := u.repo.GetUserById(id)
	if err != nil {
		return dto.UserDetailResponse{}, err
	}

	return mapper.ToUserDetail(user), nil
}
