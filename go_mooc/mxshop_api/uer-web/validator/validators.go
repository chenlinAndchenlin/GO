package myValidator

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func ValidateMobile(fl validator.FieldLevel) bool {
	mobiles := fl.Field().String()
	ok, _ := regexp.MatchString("^[1][2][0-9]{9}$", mobiles)
	if !ok {
		return false
	} else {
		return true
	}
}
