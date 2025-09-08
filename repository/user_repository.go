package repository

import "gorm.io/gorm"

type UserRepository interface {
}

type userRepository struct {
	pgdb *gorm.DB
}

func NewUserRepository(pgdb *gorm.DB) UserRepository {
	return &userRepository{pgdb}
}
