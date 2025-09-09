package helper

import "errors"

var (
	ErrInvalidRequest = errors.New("invalid request")

	ErrInvalidUsername = errors.New("invalid username")
	ErrInvalidPassword = errors.New("invalid password")
)
