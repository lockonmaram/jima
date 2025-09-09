package repository

import (
	"errors"
	"jima/entity/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	CheckIsUserExist(username string, email string) (isExist bool, err error)
	CreateUser(user model.User) error
}

type userRepository struct {
	pgdb *gorm.DB
}

func NewUserRepository(pgdb *gorm.DB) UserRepository {
	return &userRepository{pgdb}
}

// if 1 param match, returns true
func (r *userRepository) CheckIsUserExist(username string, email string) (isExist bool, err error) {
	var user model.User
	err = r.pgdb.Where("username = ? OR email = ?", username, email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (r *userRepository) CreateUser(user model.User) error {
	return r.pgdb.Create(&user).Error
}
