package helpers

import (
	"github.com/go-playground/validator/v10"
	"github.com/riyan-eng/api-praxis-online-class/models"
)

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

var validate = validator.New()

func ValidateClass(class models.Class) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(class)
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

func ValidateUserType(user_type models.UserType) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(user_type)
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
