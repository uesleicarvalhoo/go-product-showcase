package validator

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator"
)

var v = validator.New() //nolint: gochecknoglobals

func Validate(s any) error {
	err := v.Struct(s)
	if err == nil {
		return nil
	}

	validationErrors, ok := err.(validator.ValidationErrors) //nolint:errorlint
	if !ok {
		return err
	}

	errMessages := make([]string, len(validationErrors))

	for i, err := range validationErrors {
		switch err.Tag() {
		case "required":
			errMessages[i] = fmt.Sprintf("%s is required", err.Field())
		case "min":
			errMessages[i] = fmt.Sprintf("%s min length is %s", err.Field(), err.Param())
		case "max":
			errMessages[i] = fmt.Sprintf("%s max length is %s", err.Field(), err.Param())
		default:
			errMessages[i] = fmt.Sprintf(
				"'%s' has value of '%v' which doesn't satisfy '%s'", err.Field(), err.Value(), err.Tag(),
			)
		}
	}

	return errors.New(strings.Join(errMessages, "\n"))
}
