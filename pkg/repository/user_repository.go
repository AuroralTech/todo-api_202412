package repository

import (
	"context"

	"github.com/AuroralTech/todo-api_202412/pkg/entity"
	repository "github.com/AuroralTech/todo-api_202412/pkg/repository/interface"
	"github.com/uptrace/bun"
)

type UserRepository struct {
	db *bun.DB
}

func NewUserRepository(db *bun.DB) repository.UserRepository {
	return &UserRepository{db: db}
}

// UpdateUser ユーザーを更新
func (r *UserRepository) UpdateUser(ctx context.Context, params repository.UpdateUserParams) (int64, error) {
	user := &entity.User{
		ID:          params.ID,
		FirebaseUID: params.FirebaseUID,
		Name:        params.Name,
	}

	u, err := r.db.NewUpdate().
		Model(user).
		WherePK().
		Exec(ctx)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := u.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}
