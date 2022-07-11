package controller

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"

	"echo-practice/entity"
	"echo-practice/usecase"
)

type UserHandler struct {
	UserUC usecase.UserUseCase
}

func NewUser(e *echo.Echo, uuc *usecase.UserUseCase) {

	uh := UserHandler{UserUC: *uuc}

	e.POST("/users", uh.Create)
	e.GET("/users", uh.GetAll)
	//e.GET("/user", uh.Get)
}

func (uh *UserHandler) Create(c echo.Context) error {
	var u entity.User
	var err error

	if err = c.Bind(&u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = uh.UserUC.Create(c.Request().Context(), &u)

	if err != nil {
		if errors.Is(err, entity.ErrInternalService) {
			ErrorResponse(c, http.StatusInternalServerError, err)
		}
		ErrorResponse(c, http.StatusBadRequest, err)

		return nil
	}

	SuccessResponse(c, nil)

	return nil
}

func (uh *UserHandler) GetAll(c echo.Context) error {
	var us []entity.User
	var fu entity.FilterUser
	var err error

	if err = c.Bind(&fu); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	us, err = uh.UserUC.GetAll(c.Request().Context(), &fu)

	if err != nil {
		if errors.Is(err, entity.ErrInternalService) {
			ErrorResponse(c, http.StatusInternalServerError, err)
		}
		ErrorResponse(c, http.StatusBadRequest, err)

		return nil
	}

	SuccessResponse(c, us)

	return nil
}

// func (uh *UserHandler) Get(c echo.Context) error {
// 	var u *entity.User
// 	var err error

// 	// if err = c.Bind(&u); err != nil {
// 	// 	return echo.NewHTTPError(http.StatusBadRequest, err.Error())
// 	// }

// 	u, err = uh.UserUC.Get(c.Request().Context())

// 	if err != nil {
// 		ErrorResponse(c, http.StatusInternalServerError, err)
// 	}

// 	SuccessResponse(c, u)

// 	return nil
// }
