package shared

import "github.com/google/uuid"

type UserAdapter interface {
	FindByID(id uuid.UUID) (*UserDTO, error)
}
