package models

import "errors"

var (
	ErrUserNotFound = errors.New("user doesn't exists")
	ErrUserExists   = errors.New("user with the same email already exists")
)
