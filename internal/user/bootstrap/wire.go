package bootstrap

import (
	"github.com/google/wire"
	"mkfolder.dev/wire-playground/internal/user/adapters/driven"
	"mkfolder.dev/wire-playground/internal/user/adapters/driving"
	"mkfolder.dev/wire-playground/internal/user/core"
)

var UserSet = wire.NewSet(
	driven.NewPostgresRepository,
	core.NewUserService,
	driving.NewHTTPAdapter,
	wire.Bind(new(core.UserRepository), new(*driven.PostgresRepository)),
)
