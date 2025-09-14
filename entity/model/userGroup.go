package model

import "time"

type UserGroupRole string

const (
	UserGroupSerialPrefix = "USRGRP"

	UserGroupRoleManager UserGroupRole = "manager"
	UserGroupRoleMember  UserGroupRole = "member"
)

type UserGroup struct {
	Serial      string        `db:"serial" json:"serial"`
	UserSerial  string        `db:"user_serial" json:"user_serial"`
	GroupSerial string        `db:"group_serial" json:"group_serial"`
	Role        UserGroupRole `db:"role" json:"role"`
	CreatedAt   time.Time     `db:"created_at" json:"created_at"`
	UpdatedAt   *time.Time    `db:"updated_at" json:"updated_at"`
	DeletedAt   *time.Time    `db:"deleted_at" json:"deleted_at"`
}

func (UserGroup) TableName() string {
	return "auth.user_groups"
}
