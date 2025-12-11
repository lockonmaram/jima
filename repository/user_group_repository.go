package repository

import (
	"jima/entity/model"
	"jima/helper"
	"strings"

	"gorm.io/gorm"
)

type UserGroupRepository interface {
	GetUserGroup(userSerial, groupSerial string) (response *model.UserGroup, err error)
	AddUserToGroup(userSerial, groupSerial string) (response *model.UserGroup, err error)
	RemoveUserFromGroup(userGroupSerial string) (err error)
	GetUserGroups(userSerial string) (response []*model.UserGroup, err error)
	GetUserGroupMembersByGroupSerial(groupSerial string) (response []*model.UserGroup, err error)
}

type userGroupRepository struct {
	pgdb *gorm.DB
}

func NewUserGroupRepository(pgdb *gorm.DB) UserGroupRepository {
	return &userGroupRepository{pgdb}
}

func (r *userGroupRepository) GetUserGroup(userSerial, groupSerial string) (response *model.UserGroup, err error) {
	err = r.pgdb.Where("user_serial = ? AND group_serial = ? AND deleted_at IS NULL", userSerial, groupSerial).First(&response).Error
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (r *userGroupRepository) AddUserToGroup(userSerial, groupSerial string) (response *model.UserGroup, err error) {
	// Generate user group serial
	userGroupSerial := helper.GenerateSerialFromString(model.UserGroupSerialPrefix, strings.Split(groupSerial, "-")[1])

	userGroup := model.UserGroup{
		Serial:      userGroupSerial,
		UserSerial:  userSerial,
		GroupSerial: groupSerial,
		Role:        model.UserGroupRoleMember,
		CreatedBy:   userSerial,
	}
	if err := r.pgdb.Create(&userGroup).Error; err != nil {
		return nil, err
	}

	return &userGroup, nil
}

func (r *userGroupRepository) RemoveUserFromGroup(userGroupSerial string) (err error) {
	return r.pgdb.Where("serial = ?", userGroupSerial).Delete(&model.UserGroup{}).Error
}

func (r *userGroupRepository) GetUserGroups(userSerial string) (response []*model.UserGroup, err error) {
	err = r.pgdb.
		Preload("Group").
		Where("user_serial = ?", userSerial).
		Find(&response).Error
	return
}

func (r *userGroupRepository) GetUserGroupMembersByGroupSerial(groupSerial string) (response []*model.UserGroup, err error) {
	err = r.pgdb.
		Preload("User").
		Where("group_serial = ?", groupSerial).
		Find(&response).Error
	return
}
