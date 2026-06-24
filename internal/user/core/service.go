package core

import "github.com/google/uuid"

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Create(user *User) error {
	return s.repo.Create(user)
}

func (s *UserService) FindByID(id uuid.UUID) (*User, error) {
	return s.repo.FindByID(id)
}
