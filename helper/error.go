package helper

import (
	"errors"
)

var (
	ErrForbiddenUserAction = errors.New("user action is not allowed")
	ErrUnauthorizedToken   = errors.New("unauthorized token")
	ErrTokenRequired       = errors.New("token is required")

	ErrDatabase       = errors.New("database error")
	ErrInvalidRequest = errors.New("invalid request")

	ErrInvalidUsernameEmail = errors.New("invalid username / email")
	ErrInvalidPassword      = errors.New("invalid password")

	ErrUserAlreadyExists = errors.New("user already exists")
	ErrUserNotFound      = errors.New("user not found")
)
