package user

import(
	"errors"
)

var(
	userNotExists = errors.New("This user does not exists")
	userAlreadyExists = errors.New("This user already exists")
)