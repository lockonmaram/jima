package model

import "time"

type Group struct {
	Serial    string     `db:"serial" json:"serial"`
	Name      string     `db:"name" json:"name"`
	CreatedAt time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt *time.Time `db:"updated_at" json:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at" json:"deleted_at"`
}

func (Group) TableName() string {
	return "auth.groups"
}
