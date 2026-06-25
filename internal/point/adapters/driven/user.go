package driven

import (
	"github.com/google/uuid"
	"mkfolder.dev/wire-playground/internal/point/core"
	"mkfolder.dev/wire-playground/internal/shared"
)

type UserService struct {
	svc shared.UserAdapter
}

func NewUserService(svc shared.UserAdapter) *UserService {
	return &UserService{svc: svc}
}

func (s *UserService) FindByID(id uuid.UUID) (*core.User, error) {
	u, err := s.svc.FindByID(id)
	if err != nil {
		return nil, err
	}
	return &core.User{
		ID:       u.ID,
		Username: u.Username,
		Email:    u.Email,
	}, nil
}
