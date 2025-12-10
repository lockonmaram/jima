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

	CreatedAt time.Time  `db:"created_at" json:"created_at"`
	CreatedBy string     `db:"created_by" json:"created_by"`
	UpdatedAt *time.Time `db:"updated_at" json:"updated_at"`
	UpdatedBy *string    `db:"updated_by" json:"updated_by"`
	DeletedAt *time.Time `db:"deleted_at" json:"deleted_at"`
	DeletedBy *string    `db:"deleted_by" json:"deleted_by"`
}

func (UserGroup) TableName() string {
	return "jima_auth.user_groups"
}
