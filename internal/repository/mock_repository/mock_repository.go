package mockrepo

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

type MockUserRepository struct {
	DB model.DB
}

func NewMockUserRepository(db model.DB) UserRepository {
	return &MockUserRepository{DB: db}
}

func (r *MockUserRepository) Create(user *model.User) error {
	return nil
}
