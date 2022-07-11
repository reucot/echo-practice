package main

import (
	"echo-practice/controller"
	"echo-practice/pkg/logger"
	"echo-practice/pkg/postgres"
	"echo-practice/pkg/validator"
	"echo-practice/usecase"
	"echo-practice/usecase/psql"
	"fmt"

	"github.com/labstack/echo/v4"
)

func main() {
	l := logger.New("DEBUG")

	pg, err := postgres.New("postgres://postgres:1q2w3e4r@localhost:5432/tz_users", postgres.MaxPoolSize(10))
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}
	defer pg.Close()

	e := echo.New()
	e.Validator = validator.NewValidator()

	controller.NewUser(e, usecase.NewUserUseCase(psql.NewUserRepo(pg)))

	l.Fatal(e.Start(":1323"))
}
