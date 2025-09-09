package repository

import (
	"jima/entity/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByUsernameOrEmail(username string, email string) (user *model.User, err error)
	CreateUser(user model.User) error
}

type userRepository struct {
	pgdb *gorm.DB
}

func NewUserRepository(pgdb *gorm.DB) UserRepository {
	return &userRepository{pgdb}
}

// if 1 param match, returns true
func (r *userRepository) GetUserByUsernameOrEmail(username string, email string) (user *model.User, err error) {
	err = r.pgdb.Where("username = ? OR email = ?", username, email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) CreateUser(user model.User) error {
	return r.pgdb.Create(&user).Error
}
