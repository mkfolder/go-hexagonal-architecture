package driving

import (
	"github.com/google/uuid"
	"mkfolder.dev/wire-playground/internal/shared"
	"mkfolder.dev/wire-playground/internal/user/core"
)

type HexagonAdapter struct {
	svc *core.UserService
}

func NewHexagonAdapter(svc *core.UserService) *HexagonAdapter {
	return &HexagonAdapter{svc: svc}
}

func (a *HexagonAdapter) FindByID(id uuid.UUID) (*shared.UserDTO, error) {
	user, err := a.svc.FindByID(id)
	if err != nil {
		return nil, err
	}
	dto := user.DTO()
	return &dto, nil
}
