package model

import (
	"time"
)

type UserRole string

const (
	UserRoleAdmin UserRole = "admin"
	UserRoleUser  UserRole = "user"

	UserSerialPrefix = "USR"
)

type User struct {
	Serial    string     `db:"serial" json:"serial"`
	Username  string     `db:"username" json:"username"`
	Email     string     `db:"email" json:"email"`
	Name      string     `db:"name" json:"name"`
	Password  string     `db:"password" json:"password"`
	Role      string     `db:"role" json:"role"`
	CreatedAt time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt *time.Time `db:"updated_at" json:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at" json:"deleted_at"`
}

func (User) TableName() string {
	return "auth.users"
}
