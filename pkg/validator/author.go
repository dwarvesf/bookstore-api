package validator

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

// AuthorValidator is a custom validator for author
var AuthorValidator validator.Func = func(fl validator.FieldLevel) bool {
	field, ok := fl.Field().Interface().(string)
	if ok {
		pattern := "^[a-zA-Z ]+$"
		regex := regexp.MustCompile(pattern)

		return regex.MatchString(field)
	}

	return false
}
