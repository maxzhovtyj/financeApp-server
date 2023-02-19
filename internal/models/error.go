package models

import "errors"

var (
	ErrUserNotFound      = errors.New("user was not found")
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrInvalidInputBody  = errors.New("invalid input body")
)
