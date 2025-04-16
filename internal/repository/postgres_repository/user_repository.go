package postgresrepo

import (
	"context"

	"github.com/TurkmenistanRailways/payment/internal/config"
	"github.com/TurkmenistanRailways/payment/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresUserRepository struct {
	db *pgxpool.Pool
}

func NewPostgresUserRepository(db *pgxpool.Pool, conf *config.Config) model.UserRepository {

	if conf.MOCK_REPOSITORY == "true" {
		return &MockUserRepository{db}
	}

	return &PostgresUserRepository{db}
}

func (r *PostgresUserRepository) Create(ctx context.Context, user *model.User) error {
	_, err := r.db.Exec(ctx, "INSERT INTO users (name, email, password) VALUES ($1, $2, $3)", user.Name, user.Email, user.Password)

	return err
}

func (r *PostgresUserRepository) GetByID(ctx context.Context, id int64) (*model.User, error) {
	var user model.User
	err := r.db.QueryRow(ctx, "SELECT * FROM users WHERE id = $1", id).Scan(&user.ID, &user.Name, &user.Email, &user.Password)

	return &user, err
}
