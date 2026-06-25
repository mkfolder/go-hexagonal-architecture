package bootstrap

import (
	"github.com/google/wire"
	"mkfolder.dev/wire-playground/internal/profile/adapters/driven"
	"mkfolder.dev/wire-playground/internal/profile/adapters/driving"
	"mkfolder.dev/wire-playground/internal/profile/core"
)

var ProfileSet = wire.NewSet(
	driven.NewUserService,
	driven.NewPointService,
	core.NewProfileService,
	driving.NewHTTPAdapter,
	wire.Bind(new(core.PointsFetcher), new(*driven.PointService)),
	wire.Bind(new(core.UserFetcher), new(*driven.UserService)),
)
