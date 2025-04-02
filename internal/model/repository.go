package model

type UserRepository interface {
	Create(user *User) error
	// GetByID(id int64) (*User, error)
	// GetAll() ([]*User, error)

	// Delete(id int64) error
}
