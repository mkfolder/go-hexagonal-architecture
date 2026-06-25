package core

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
}

type UserFetcher interface {
	FindByID(id uuid.UUID) (*User, error)
}

type PointsFetcher interface {
	GetPoints(userID uuid.UUID) (uint64, error)
}
