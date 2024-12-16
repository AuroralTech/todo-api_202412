package usecase

import (
	"context"

	repository "github.com/AuroralTech/todo-api_202412/pkg/repository/interface"
)

type userUsecase struct {
	userRepository repository.UserRepository
}

type UserUsecaseInput struct {
	ID   string
	Name string
}

type UserUsecase interface {
	UpdateUser(ctx context.Context, input UserUsecaseInput) (int64, error)
}

func NewUserUsecase(userRepository repository.UserRepository) UserUsecase {
	return &userUsecase{userRepository: userRepository}
}

func (u *userUsecase) UpdateUser(ctx context.Context, input UserUsecaseInput) (int64, error) {
	params := repository.UpdateUserParams{
		ID:   input.ID,
		Name: input.Name,
	}
	updatedCounts, err := u.userRepository.Update(ctx, params)
	if err != nil {
		return 0, err
	}
	return updatedCounts, nil
}
