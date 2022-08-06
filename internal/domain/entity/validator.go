package entity

import "github.com/uesleicarvalhoo/go-product-showcase/pkg/validator"

func Validate(s any) error {
	return validator.Validate(s)
}
