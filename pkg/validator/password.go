package validator

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

// PasswordValidator is a custom validator for password
var PasswordValidator validator.Func = func(fl validator.FieldLevel) bool {
	field, ok := fl.Field().Interface().(string)
	if ok {
		// check if password is at least 8 characters long
		if len(field) < 8 {
			return false
		}

		containsUppercase := regexp.MustCompile(`[A-Z]`).MatchString(field)
		containsLowercase := regexp.MustCompile(`[a-z]`).MatchString(field)
		containsSpecialChar := regexp.MustCompile(`[!@#$%^&*()_+{}<>\?~]`).MatchString(field)

		return containsUppercase && containsLowercase && containsSpecialChar
	}

	return false
}
