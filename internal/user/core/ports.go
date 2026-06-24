package core

import (
	"github.com/google/uuid"
)

type UserRepository interface {
	Create(user *User) error
	FindByID(id uuid.UUID) (*User, error)
}
