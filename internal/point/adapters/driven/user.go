package driven

import (
	"github.com/google/uuid"
	"mkfolder.dev/wire-playground/internal/point/core"
	user "mkfolder.dev/wire-playground/internal/user/core"
)

type UserService struct {
	svc *user.UserService
}

func NewUserService(svc *user.UserService) *UserService {
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
