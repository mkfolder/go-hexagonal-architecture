package core

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"mkfolder.dev/wire-playground/internal/shared"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
}

func (u User) Validate() error {
	if u.Email == "" {
		return errors.New("email is required")
	}
	if u.Username == "" {
		return errors.New("username is required")
	}
	return nil
}

func (u User) DTO() shared.UserDTO {
	return shared.UserDTO{
		ID:        u.ID,
		Email:     u.Email,
		Username:  u.Username,
		CreatedAt: u.CreatedAt,
	}
}
