package core

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID
	Username string
	Email    string
}

type PointService struct {
	userFetcher UserFetcher
}

func NewPointService(userFetcher UserFetcher) *PointService {
	return &PointService{userFetcher: userFetcher}
}

func (s *PointService) GetPoints(uuid.UUID) (uint64, error) {
	return 127, nil
}

func (s *PointService) GetLeaderboard(filterID uuid.UUID) ([]User, error) {
	user, err := s.userFetcher.FindByID(filterID)
	if err != nil {
		return nil, err
	}

	users := []User{
		*user,
	}

	return users, nil
}
