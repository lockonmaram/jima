package helper

import (
	"errors"
)

var (
	ErrDatabase       = errors.New("database error")
	ErrInvalidRequest = errors.New("invalid request")

	ErrInvalidUsername = errors.New("invalid username")
	ErrInvalidPassword = errors.New("invalid password")

	ErrUserAlreadyExists = errors.New("user already exists")
)
