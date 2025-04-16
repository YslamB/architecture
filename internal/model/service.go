package model

import "context"

type UserService interface {
	CreateUser(ctx context.Context, user *User) error
	GetUserByID(ctx context.Context, id int64) (*User, error)
}
