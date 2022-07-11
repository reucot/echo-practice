package psql

import (
	"context"
	"echo-practice/entity"
	"echo-practice/pkg/postgres"
	"fmt"
)

type UserRepo struct {
	*postgres.Postgres
}

func NewUserRepo(pg *postgres.Postgres) *UserRepo {
	return &UserRepo{pg}
}

func (ur UserRepo) Insert(ctx context.Context, u *entity.User) error {

	q := `
	INSERT INTO users
	VALUES (DEFAULT, $1, $2, $3, $4)`

	_, err := ur.Pool.Exec(ctx, q, u.FirstName, u.SecondName, u.DateOfBirth.Time.Format("2006-02-01"), u.IncomePerYear.Icy)

	//TODO: Add internal server error
	if err != nil {
		return fmt.Errorf("psql - user - Insert - ur.Pool.Exec: %w", fmt.Errorf("%s: %w", err.Error(), entity.ErrInternalService))
	}

	return nil
}

func (ur UserRepo) Get(ctx context.Context) (*entity.User, error) {
	return nil, nil
}

func (ur UserRepo) GetAll(ctx context.Context, fu *entity.FilterUser) ([]entity.User, error) {
	var us []entity.User

	q := `
	SELECT
		*
	FROM
	users
	`

	filterStr := fu.Filter()

	if len(filterStr) > 0 {
		q += `WHERE ` + filterStr
	}

	rows, err := ur.Pool.Query(ctx, q)

	if err != nil {
		return nil, fmt.Errorf("psql - user - GetAll - ur.Pool.Query: %w", fmt.Errorf("%s: %w", err.Error(), entity.ErrInternalService))
	}

	for rows.Next() {
		u := new(entity.User)

		err = rows.Scan(&u.Id, &u.FirstName, &u.SecondName, &u.DateOfBirth.Time, &u.IncomePerYear.Icy)

		if err != nil {
			return nil, fmt.Errorf("psql - user - GetAll - rows.Scan: %w", fmt.Errorf("%s: %w", err.Error(), entity.ErrInternalService))
		}

		us = append(us, *u)
	}

	return us, nil
}
