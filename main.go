package main

import (
	"echo-practice/controller"
	"echo-practice/pkg/validator"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.Validator = validator.NewValidator()

	controller.NewUser(e)

	e.Logger.Fatal(e.Start(":1323"))
}
