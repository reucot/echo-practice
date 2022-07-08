package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"echo-practice/entity"
	"echo-practice/usecase"
)

type UserHandler struct {
	UserUC usecase.UserUseCase
}

func NewUser(e *echo.Echo) {

	uh := UserHandler{}

	e.POST("/users", uh.Create)
	e.GET("/users", uh.GetAll)
	e.GET("/user", uh.Get)
}

func (uh *UserHandler) Create(c echo.Context) error {
	var u entity.User
	var err error

	if err = c.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err = c.Validate(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return nil
}

func (uh *UserHandler) GetAll(c echo.Context) error {
	return nil
}

func (uh *UserHandler) Get(c echo.Context) error {
	return nil
}
