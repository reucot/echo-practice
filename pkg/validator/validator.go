package validator

import (
	"strconv"

	"github.com/go-playground/validator"
)

type CustomValidator struct {
	v validator.Validate
}

func NewValidator() *CustomValidator {
	return &CustomValidator{
		v: *validator.New(),
	}
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.v.Struct(i); err != nil {
		//validationErrors := err.(validator.ValidationErrors)
		return err
	}

	cv.v.RegisterValidation("income_per_year", IncomePerYear)

	return nil
}

func IncomePerYear(fl validator.FieldLevel) bool {
	icyStr := fl.Field().String()

	if icy, err := strconv.ParseFloat(icyStr, 2); err != nil {
		fl.Field().SetInt(int64(icy * 100))
		return true
	}

	return false
}
