package validators

import (
	"fifth_test/models"
	"github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
	Value string `json:"value"`
}

func ValidateStruct(post models.Post) []*ErrorResponse {
	var errors []*ErrorResponse

	validate := validator.New()

	err := validate.Struct(post)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.Field = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
