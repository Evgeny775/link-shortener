package req

import (
	"github.com/go-playground/validator"
)

func Validate[payloadType any](payload *payloadType) error {
	var payloadRequest payloadType
	validate := validator.New()
	err := validate.Struct(&payloadRequest)
	if err != nil {
		return err
	}
	return nil
}
