package postgresrepo

import (
	"architecture/internal/model"
)

type UserRepository interface {
	Create(user *model.User) error
	// GetByID(id int64) (*model.User, error)
	// GetAll() ([]*model.User, error)
	// Update(user *model.User) error
	// Delete(id int64) error
}

type PostgresUserRepository struct {
	DB model.DB
}

func NewPostgresUserRepository(db model.DB) UserRepository {
	return &PostgresUserRepository{DB: db}
}

func (r *PostgresUserRepository) Create(user *model.User) error {
	_, err := r.DB.Exec("INSERT INTO users (name, email, password) VALUES ($1, $2, $3)", user.Name, user.Email, user.Password)

	return err
}
