package service

import (
	"architecture/internal/model"
)

type UserService struct {
	UserRepository model.UserRepository
}

func NewUserService(repo model.UserRepository) *UserService {
	return &UserService{UserRepository: repo}
}

func (s *UserService) CreateUser(user *model.User) error {
	return s.UserRepository.Create(user)
}

// func (s *UserService) GetUserByID(id int64) (*model.User, error) {
// 	return s.UserInframodel.GetByID(id)
// }

// func (s *UserService) GetAllUsers() ([]*model.User, error) {
// 	return s.UserInframodel.GetAll()
// }

// func (s *UserService) UpdateUser(user *model.User) error {
// 	return s.UserInframodel.Update(user)
// }

// func (s *UserService) DeleteUser(id int64) error {
// 	return s.UserInframodel.Delete(id)
// }
