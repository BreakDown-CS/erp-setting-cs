package repository

import (
	"users-api/modules/users/model"

	"gorm.io/gorm"
)

type Repository interface {
	InsetUser(user model.User) error
	GetAllUser(limit, offset int) ([]model.User, int, error)
	GetUserById(userId int) (model.User, error)
}

type repository struct {
	db *gorm.DB
}

var _ Repository = (*repository)(nil)

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) InsetUser(user model.User) error {
	return r.db.Create(&user).Error
}

func (r *repository) GetAllUser(limit, offset int) ([]model.User, int, error) {

	var users []model.User
	var total int64

	r.db.Model(&model.User{}).Count(&total)

	err := r.db.
		Limit(limit).
		Offset(offset).
		Find(&users).Error

	return users, int(total), err
}

func (r *repository) GetUserById(userId int) (model.User, error) {
	var user model.User

	query := `
		SELECT id, username, name, password
		FROM users
		WHERE id = ?
	`

	err := r.db.Raw(query, userId).Scan(&user).Error
	return user, err
}
