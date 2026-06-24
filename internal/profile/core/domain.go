package core

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Profile struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	Points    uint64    `json:"points"`
	CreatedAt time.Time `json:"created_at"`
}

func (p *Profile) Validate() error {
	if p.Email == "" {
		return errors.New("email is required")
	}
	if p.Username == "" {
		return errors.New("username is required")
	}
	return nil
}
