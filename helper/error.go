package helper

import (
	"errors"
)

var (
	ErrDatabase       = errors.New("database error")
	ErrInvalidRequest = errors.New("invalid request")

	ErrInvalidUsernameEmail = errors.New("invalid username / email")
	ErrInvalidPassword      = errors.New("invalid password")

	ErrUserAlreadyExists = errors.New("user already exists")
	ErrUserNotFound      = errors.New("user not found")
)
