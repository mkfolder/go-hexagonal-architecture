package core

import "github.com/google/uuid"

type ProfileService struct {
	userFetcher   UserFetcher
	pointsFetcher PointsFetcher
}

func NewProfileService(userFetcher UserFetcher, pointsFetcher PointsFetcher) *ProfileService {
	return &ProfileService{userFetcher: userFetcher, pointsFetcher: pointsFetcher}
}

func (s *ProfileService) FindByID(id uuid.UUID) (*Profile, error) {
	u, err := s.userFetcher.FindByID(id)
	if err != nil {
		return nil, err
	}

	points, err := s.pointsFetcher.GetPoints(u.ID)
	if err != nil {
		return nil, err
	}

	return &Profile{
		ID:        u.ID,
		Email:     u.Email,
		Username:  u.Username,
		Points:    points,
		CreatedAt: u.CreatedAt,
	}, nil
}
