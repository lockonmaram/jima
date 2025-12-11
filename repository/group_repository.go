package repository

import (
	api_entity "jima/entity/api"
	"jima/entity/model"
	"jima/helper"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type GroupRepository interface {
	CreateGroup(group *model.Group, userSerial string) (response *model.Group, err error)
	GetGroupBySerial(groupSerial string) (response *model.Group, err error)
	UpdateGroup(request api_entity.GroupsUpdateGroupRequest) (response *model.Group, err error)
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

		// Generate user group serial
		userGroupSerial := helper.GenerateSerialFromString(model.UserGroupSerialPrefix, strings.Split(group.Serial, "-")[1])

		// Assign user to group
		userGroup := model.UserGroup{
			Serial:      userGroupSerial,
			UserSerial:  userSerial,
			GroupSerial: group.Serial,
			Role:        model.UserGroupRoleManager,
			CreatedBy:   userSerial,
		}
		if err := tx.Create(&userGroup).Error; err != nil {
			return err
		}

		response = group
		return nil
	})

	return response, err
}

func (r *groupRepository) GetGroupBySerial(groupSerial string) (response *model.Group, err error) {
	err = r.pgdb.Where("serial = ? AND deleted_at IS NULL", groupSerial).First(&response).Error
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (r *groupRepository) UpdateGroup(request api_entity.GroupsUpdateGroupRequest) (response *model.Group, err error) {
	updatePayload := map[string]any{}

	if request.Name != "" {
		updatePayload["name"] = request.Name
	}

	if len(updatePayload) == 0 {
		return nil, helper.ErrInvalidRequest
	}

	err = r.pgdb.Model(&model.Group{}).
		Where("serial = ?", request.GroupSerial).
		Clauses(clause.Returning{}).
		Updates(updatePayload).
		Scan(&response).Error
	if err != nil {
		return nil, err
	}

	return
}
