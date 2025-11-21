package model

import "time"

const (
	GroupSerialPrefix = "GRP"
)

type Group struct {
	Serial string `db:"serial" json:"serial"`
	Name   string `db:"name" json:"name"`

	CreatedAt time.Time  `db:"created_at" json:"created_at"`
	CreatedBy string     `db:"created_by" json:"created_by"`
	UpdatedAt *time.Time `db:"updated_at" json:"updated_at"`
	UpdatedBy *string    `db:"updated_by" json:"updated_by"`
	DeletedAt *time.Time `db:"deleted_at" json:"deleted_at"`
	DeletedBy *string    `db:"deleted_by" json:"deleted_by"`
}

func (Group) TableName() string {
	return "jima_auth.groups"
}
