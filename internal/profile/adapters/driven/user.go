package driven

import (
	"github.com/google/uuid"
	"mkfolder.dev/wire-playground/internal/user/core"
)

type UserService struct {
	svc *core.UserService
}

func NewUserService(svc *core.UserService) *UserService {
	return &UserService{svc: svc}
}

func (s *UserService) FindByID(id uuid.UUID) (*core.User, error) {
	return s.svc.FindByID(id)
}
