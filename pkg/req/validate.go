package req

import (
	"github.com/go-playground/validator"
)

func Validate[payloadType any](payload *payloadType) error {
	validate := validator.New()
	err := validate.Struct(payload)
	if err != nil {
		return err
	}
	return nil
}
