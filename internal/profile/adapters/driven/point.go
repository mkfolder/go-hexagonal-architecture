package driven

import (
	"github.com/google/uuid"
	"mkfolder.dev/wire-playground/internal/point/core"
)

type PointService struct {
	svc *core.PointService
}

func NewPointService(svc *core.PointService) *PointService {
	return &PointService{svc: svc}
}

func (p *PointService) GetPoints(userID uuid.UUID) (uint64, error) {
	return p.svc.GetPoints(userID)
}
