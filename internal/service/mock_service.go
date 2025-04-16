package service

import (
	"context"

	"github.com/TurkmenistanRailways/payment/internal/model"
)

type MockUserService struct {
	UserRepository model.UserRepository
}

func (s *MockUserService) CreateUser(ctx context.Context, user *model.User) error {
	return s.UserRepository.Create(ctx, user)
}

func (s *MockUserService) GetUserByID(ctx context.Context, id int64) (*model.User, error) {
	return s.UserRepository.GetByID(ctx, id)
}
