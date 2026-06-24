package core

import (
	"github.com/google/uuid"
	"mkfolder.dev/wire-playground/internal/user/core"
)

type UserFetcher interface {
	FindByID(id uuid.UUID) (*core.User, error)
}

type PointsFetcher interface {
	GetPoints(userID uuid.UUID) (uint64, error)
}
