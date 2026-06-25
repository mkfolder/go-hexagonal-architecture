package bootstrap

import (
	"github.com/google/wire"
	"mkfolder.dev/wire-playground/internal/point/adapters/driven"
	"mkfolder.dev/wire-playground/internal/point/adapters/driving"
	"mkfolder.dev/wire-playground/internal/point/core"
)

var PointSet = wire.NewSet(
	driven.NewUserService,
	core.NewPointService,
	driving.NewHTTPAdapter,
	wire.Bind(new(core.UserFetcher), new(*driven.UserService)),
)
