package validators

import (
	"fmt"
	"regexp"

	"github.com/go-playground/validator/v10"
)

func (cv *cValidator) HexColor() validator.Func {
	return func(fl validator.FieldLevel) bool {
		value := fmt.Sprint(fl.Field().Interface())
		match, err := regexp.MatchString("^#[0-9a-fA-F]{6}$", value)
		if err != nil {
			return false
		}
		return match
	}
}
