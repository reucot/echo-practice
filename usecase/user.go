package usecase

import (
	"context"
	"fmt"

	"echo-practice/entity"
	"echo-practice/usecase/psql"
)

type UserUseCase struct {
	UserRepo psql.UserRepo
}

func NewUserUseCase(ur *psql.UserRepo) *UserUseCase {
	return &UserUseCase{
		UserRepo: *ur,
	}
}

func (uc *UserUseCase) Create(ctx context.Context, u *entity.User) error {

	nu := *u

	err := uc.UserRepo.Insert(ctx, &nu)

	if err != nil {
		return fmt.Errorf("usecase - user - Create - uc.UserRepo.Insert: %w", err)
	}

	return nil
}

func (uc *UserUseCase) GetAll(ctx context.Context, fu *entity.FilterUser) ([]entity.User, error) {

	us, err := uc.UserRepo.GetAll(ctx, fu)

	if err != nil {
		return nil, fmt.Errorf("usecase - user - GetAll - uc.UserRepo.GetAll: %w", err)
	}

	return us, nil
}

func (uc *UserUseCase) Get(ctx context.Context) (*entity.User, error) {

	return nil, nil
}
