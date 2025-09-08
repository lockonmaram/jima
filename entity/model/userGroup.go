package model

import "time"

type UserGroup struct {
	Serial      string     `db:"serial" json:"serial"`
	UserSerial  string     `db:"user_serial" json:"user_serial"`
	GroupSerial string     `db:"group_serial" json:"group_serial"`
	CreatedAt   time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt   *time.Time `db:"updated_at" json:"updated_at"`
	DeletedAt   *time.Time `db:"deleted_at" json:"deleted_at"`
}
