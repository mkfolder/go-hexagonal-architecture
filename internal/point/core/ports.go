package core

import "github.com/google/uuid"

type UserFetcher interface {
	FindByID(id uuid.UUID) (*User, error)
}
