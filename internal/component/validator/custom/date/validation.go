package date

import "github.com/go-playground/validator/v10"

const Tag = "date"

func Validation(fl validator.FieldLevel) bool {
	return true
}
