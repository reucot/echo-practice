package usecase

import (
	"context"

	"echo-practice/entity"
)

type UserUseCase struct {
	//repo
}

func NewUserUseCase() *UserUseCase {
	return &UserUseCase{}
}

func (uc *UserUseCase) Create(ctx context.Context, u *entity.User) error {

	return nil
}

func (uc *UserUseCase) GetAll(ctx context.Context) ([]entity.User, error) {

	return nil, nil
}

func (uc *UserUseCase) Get(ctx context.Context) (*entity.User, error) {

	return nil, nil
}
