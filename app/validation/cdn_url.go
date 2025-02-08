package validation

import "github.com/go-playground/validator/v10"

// CdnURL cdn url
const CdnURL = "<URL>"

// CdnURLValidation validate cdn url
func CdnURLValidation(fl validator.FieldLevel) bool {
	url := fl.Field().String()
	return len(url) >= 22 && url[:22] == CdnURL
}
