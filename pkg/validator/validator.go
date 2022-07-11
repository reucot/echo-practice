package validator

import (
	"fmt"

	"github.com/go-playground/validator"
)

type CustomValidator struct {
	v validator.Validate
}

func NewValidator() *CustomValidator {
	cv := &CustomValidator{
		v: *validator.New(),
	}

	return cv
}

func (cv *CustomValidator) Validate(i interface{}) error {
	fmt.Println("Validate")
	if err := cv.v.Struct(i); err != nil {
		//validationErrors := err.(validator.ValidationErrors)
		return err
	}

	return nil
}
