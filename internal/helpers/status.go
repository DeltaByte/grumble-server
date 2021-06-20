package helpers

import (
	"github.com/go-playground/validator/v10"
	"google.golang.org/genproto/googleapis/rpc/code"
	"google.golang.org/genproto/googleapis/rpc/status"
)

func StatusOK() *status.Status {
	return &status.Status{
		Code: int32(code.Code_OK),
	}
}

func StatusValidationFailed(err error) *status.Status {
	// handle actual errors
	if _, ok := err.(*validator.InvalidValidationError); ok {
		return &status.Status{
			Code: int32(code.Code_INTERNAL),
		}
	}

	validationErrors := err.(validator.ValidationErrors)
	return &status.Status{
		Code: int32(code.Code_INVALID_ARGUMENT),
		Message: "Validation failed.",
		Details: validationErrors,
	}
}