package validators

import (
	"github.com/go-playground/validator/v10"
	"regexp"
	"test1/auth_api/repo"
)

var IsEmail validator.Func = func(fl validator.FieldLevel) bool {
	email, ok := fl.Field().Interface().(string)
	if ok {
		match, _ := regexp.MatchString("[^@ \\t\\r\\n]+@[^@ \\t\\r\\n]+\\.[^@ \\t\\r\\n]+", email)
		return match
	}
	return false
}

var EmailUnique validator.Func = func(fl validator.FieldLevel) bool {
	email, ok := fl.Field().Interface().(string)
	if ok {
		return !repo.EmailUserIdentityExists(email)
	}
	return false
}
