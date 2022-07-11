package controller

import (
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type DefaultHandler struct {
	v *validator.Validate
	//logger
}

func NewDefaultHandler() *DefaultHandler {
	return &DefaultHandler{
		v: validator.New(),
	}
}

func ErrorResponse(c echo.Context, code int, err error) {
	c.JSON(code, err.Error())
}

func SuccessResponse(c echo.Context, body interface{}) {
	c.JSON(http.StatusOK, body)
}
