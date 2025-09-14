package repository

import (
	"jima/entity/model"
	"jima/helper"
	"strings"

	"gorm.io/gorm"
)

type GroupRepository interface {
	CreateGroup(group *model.Group, userSerial string) (response *model.Group, err error)
}

type groupRepository struct {
	pgdb *gorm.DB
}

func NewGroupRepository(pgdb *gorm.DB) GroupRepository {
	return &groupRepository{pgdb}
}

func (r *groupRepository) CreateGroup(group *model.Group, userSerial string) (response *model.Group, err error) {
	err = r.pgdb.Transaction(func(tx *gorm.DB) error {
		// Create group
		if err := tx.Create(&group).Error; err != nil {
			return err
		}

		// Generate yser group serial
		userGroupSerial := helper.GenerateSerialFromString(model.UserGroupSerialPrefix, strings.Split(group.Serial, "-")[1])

		// Assign user to group
		userGroup := model.UserGroup{
			Serial:      userGroupSerial,
			UserSerial:  userSerial,
			GroupSerial: group.Serial,
		}
		if err := tx.Create(&userGroup).Error; err != nil {
			return err
		}

		response = group
		return nil
	})

	return response, err
}
