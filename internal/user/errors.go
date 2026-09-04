package user

import (
	"errors"
)

var (
	NotExists        = errors.New("this user does not exists")
	AlreadyExists    = errors.New("this user already exists")
	WrongCredentials = errors.New("wrong email or password")
)
