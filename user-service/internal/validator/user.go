package validator

import (
	"github.com/giovannitgl/video-services/user-service/internal/entities"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ValidateUser(user entities.User) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(user)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
