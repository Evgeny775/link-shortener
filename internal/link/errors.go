package link

import(
	"errors"
)

var(
	noSuchId = errors.New("No such id")
	hashAlreadyExist = errors.New("This hash already exists")
)