package validation

import "github.com/go-playground/validator/v10"

const CDN_URL = "<URL>"

func CdnUrlValidation(fl validator.FieldLevel) bool {
	url := fl.Field().String()
	return len(url) >= 22 && url[:22] == CDN_URL
}
