package repository

import (
	"jima/entity/model"
	"jima/helper"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepository interface {
	GetUserBySerial(serial string) (user *model.User, err error)
	GetUserByPasswordToken(passwordToken string) (user *model.User, err error)
	GetUserByUsernameOrEmail(username string, email string) (user *model.User, err error)
	CreateUser(user model.User) error
	UpdateUserBySerial(serial string, updatePayload map[string]any) (user *model.User, err error)
	UpdateUserPasswordBySerialOrToken(identifier, password string) (err error)
	SetPasswordToken(serial, passwordToken string) (err error)
}

type userRepository struct {
	pgdb *gorm.DB
}

func NewUserRepository(pgdb *gorm.DB) UserRepository {
	return &userRepository{pgdb}
}

func (r *userRepository) GetUserBySerial(serial string) (user *model.User, err error) {
	err = r.pgdb.Where("serial = ?", serial).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) GetUserByPasswordToken(passwordToken string) (user *model.User, err error) {
	err = r.pgdb.Where("password_token = ?", passwordToken).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

// if 1 param match, returns true
func (r *userRepository) GetUserByUsernameOrEmail(username string, email string) (user *model.User, err error) {
	err = r.pgdb.Where("username = ? OR email = ?", username, email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) CreateUser(user model.User) error {
	return r.pgdb.Create(&user).Error
}

func (r *userRepository) UpdateUserBySerial(serial string, updatePayload map[string]any) (user *model.User, err error) {
	user = &model.User{}
	returningClause := clause.Returning{Columns: helper.FormatUpdatePayloadToClauseColumns(updatePayload)}
	err = r.pgdb.Model(user).Clauses(returningClause).Where("serial = ?", serial).Updates(updatePayload).Error
	return user, err
}

func (r *userRepository) UpdateUserPasswordBySerialOrToken(identifier, password string) (err error) {
	if identifier == "" {
		return helper.ErrInvalidRequest
	}
	return r.pgdb.Model(model.User{}).Where("serial = ? OR password_token = ?", identifier, identifier).Update("password", password).Error
}

func (r *userRepository) SetPasswordToken(serial, passwordToken string) (err error) {
	return r.pgdb.Model(model.User{}).Where("serial = ?", serial).Update("password_token", passwordToken).Error
}
