package repository

import (
	"context"
)

type UpdateUserParams struct {
	ID          string
	FirebaseUID string
	Name        string
}

type UserRepository interface {
	Update(ctx context.Context, input UpdateUserParams) (int64, error)
}
