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

	_, err := ur.Pool.Exec(ctx, q, u.FirstName, u.SecondName, u.DateOfBirth, u.IncomePerYear)

	//TODO: Add internal server error
	if err != nil {
		return fmt.Errorf("psql - user - Insert - ur.Pool.Exec: %w", err)
	}

	return nil
}

func (ur UserRepo) Get(ctx context.Context) (*entity.User, error) {
	return nil, nil
}

func (ur UserRepo) GetAll(ctx context.Context, fu *entity.FilterUser) ([]entity.User, error) {
	us := make([]entity.User, 10)

	q := `
	SELECT
		*
	FROM
	users
	`
	rows, err := ur.Pool.Query(ctx, q)

	if err != nil {
		return nil, fmt.Errorf("psql - user - GetAll - ur.Pool.Query: %w", err)
	}

	for rows.Next() {
		u := new(entity.User)

		err = rows.Scan(&u.Id, &u.FirstName, &u.SecondName, &u.IncomePerYear, &u.DateOfBirth)

		if err != nil {
			return nil, fmt.Errorf("psql - user - GetAll - rows.Scan: %w", err)
		}
	}

	return us, nil
}
