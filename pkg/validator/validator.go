package validator

import "github.com/go-playground/validator"

type CustomValidator struct {
	v validator.Validate
}

func NewValidator() *CustomValidator {
	return &CustomValidator{
		v: *validator.New(),
	}
}

func (cv *CustomValidator) Validate(i interface{}) error {

	return nil
}
