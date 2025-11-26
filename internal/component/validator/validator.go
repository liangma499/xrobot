package validator

import (
	"sync"
	"xrobot/internal/component/validator/custom/date"

	"github.com/go-playground/validator/v10"
)

var (
	once     sync.Once
	instance *validator.Validate
)

func Instance() *validator.Validate {
	once.Do(func() {
		instance = validator.New()
		instance.RegisterValidation(date.Tag, date.Validation)
	})

	return instance
}
