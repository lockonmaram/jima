package repository

import "gorm.io/gorm"

type GroupRepository interface {
}

type groupRepository struct {
	pgdb *gorm.DB
}

func NewGroupRepository(pgdb *gorm.DB) GroupRepository {
	return &groupRepository{pgdb}
}
