package controller

import "github.com/go-playground/validator"

type DefaultHandler struct {
	v *validator.Validate
	//logger
}

func NewDefaultHandler() *DefaultHandler {
	return &DefaultHandler{
		v: validator.New(),
	}
}
