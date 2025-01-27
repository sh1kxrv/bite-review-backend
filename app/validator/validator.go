package validator

import (
	"bitereview/validation"
	"sync"

	"github.com/go-playground/validator/v10"
)

var (
	vinstance *validator.Validate
	once      sync.Once
)

func GetValidatorInstance() *validator.Validate {
	once.Do(func() {
		vinstance = validator.New()
		vinstance.RegisterValidation("cdnURL", validation.CdnUrlValidation)
	})
	return vinstance
}
