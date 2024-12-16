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
	UpdateUser(ctx context.Context, params UpdateUserParams) (int64, error)
}
