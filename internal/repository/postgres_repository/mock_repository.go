package postgresrepo

import (
	"context"

	"github.com/TurkmenistanRailways/payment/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository interface {
	Create(ctx context.Context, user *model.User) error
	GetByID(ctx context.Context, id int64) (*model.User, error)
}

type MockUserRepository struct {
	db *pgxpool.Pool
}

func (r *MockUserRepository) Create(ctx context.Context, user *model.User) error {
	return nil
}

func (r *MockUserRepository) GetByID(ctx context.Context, id int64) (*model.User, error) {
	return &model.User{
		ID:       id,
		Name:     "John Doe",
		Email:    "john.doe@example.com",
		Password: "password",
	}, nil
}
