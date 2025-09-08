package repository

import "gorm.io/gorm"

type UserGroupRepository interface {
}

type userGroupRepository struct {
	pgdb *gorm.DB
}

func NewUserGroupRepository(pgdb *gorm.DB) UserGroupRepository {
	return &userGroupRepository{pgdb}
}
