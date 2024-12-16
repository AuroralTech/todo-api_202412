package user

import (
	"context"

	repository "github.com/AuroralTech/todo-api_202412/pkg/repository/interface"
)

type UserUsecase struct {
	userRepository repository.UserRepository
}

type UpdateUserInput struct {
	ID   string
	Name string
}

func NewUserUsecase(userRepository repository.UserRepository) *UserUsecase {
	return &UserUsecase{userRepository: userRepository}
}

func (u *UserUsecase) Execute(ctx context.Context, input UpdateUserInput) (int64, error) {
	params := repository.UpdateUserParams{
		ID:   input.ID,
		Name: input.Name,
	}
	return u.userRepository.UpdateUser(ctx, params)
}
