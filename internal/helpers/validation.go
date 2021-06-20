package helpers

import (
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/copier"
)

var copierOptions = copier.Option{
	IgnoreEmpty: true,
}

func ValidateRequest(validation interface{}, req interface{}) error {
	validate := validator.New()

	// copy data from request onto validation struct
	if err := copier.CopyWithOption(validation, req, copierOptions); err != nil {
		return err
	}

	// run the validations
	if err := validate.Struct(validation); err != nil {
		return err
	}

	return nil
}