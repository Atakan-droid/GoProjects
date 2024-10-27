package validate

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func Validate(val interface{}) (isValid bool, errors []string) {
	err := validate.Struct(val)
	if err == nil {
		return true, nil
	}

	valErrors := err.(validator.ValidationErrors)
	errors = make([]string, len(valErrors))
	for index, valError := range valErrors {
		errors[index] = fmt.Sprintf("Field %s failed validation for rule %s, actual value is %s",
			valError.Field(),
			valError.Tag(),
			valError.Value())
	}
	return false, errors
}
