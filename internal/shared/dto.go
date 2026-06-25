package shared

import (
	"time"

	"github.com/google/uuid"
)

type UserDTO struct {
	ID        uuid.UUID
	Email     string
	Username  string
	CreatedAt time.Time
}
