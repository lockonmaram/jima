package model

import (
	"time"
)

type UserRole string

const (
	UserSerialPrefix = "USR"

	UserRoleAdmin UserRole = "admin"
	UserRoleUser  UserRole = "user"
)

type User struct {
	Serial   string `db:"serial" json:"serial"`
	Username string `db:"username" json:"username"`
	Email    string `db:"email" json:"email"`
	Name     string `db:"name" json:"name"`
	Password string `db:"password" json:"password"`
	Role     string `db:"role" json:"role"`

	CreatedAt time.Time  `db:"created_at" json:"created_at"`
	CreatedBy string     `db:"created_by" json:"created_by"`
	UpdatedAt *time.Time `db:"updated_at" json:"updated_at"`
	UpdatedBy *string    `db:"updated_by" json:"updated_by"`
	DeletedAt *time.Time `db:"deleted_at" json:"deleted_at"`
	DeletedBy *string    `db:"deleted_by" json:"deleted_by"`
}

func (User) TableName() string {
	return "jima_auth.users"
}
