//go:build wireinject

package bootstrap

import (
	"github.com/gofiber/fiber/v3"
	"github.com/google/wire"
	"mkfolder.dev/wire-playground/internal/database"
	pointDriven "mkfolder.dev/wire-playground/internal/point/adapters/driven"
	pointDriving "mkfolder.dev/wire-playground/internal/point/adapters/driving"
	point "mkfolder.dev/wire-playground/internal/point/core"
	profileDriven "mkfolder.dev/wire-playground/internal/profile/adapters/driven"
	profileDriving "mkfolder.dev/wire-playground/internal/profile/adapters/driving"
	profile "mkfolder.dev/wire-playground/internal/profile/core"
	userDriven "mkfolder.dev/wire-playground/internal/user/adapters/driven"
	userDriving "mkfolder.dev/wire-playground/internal/user/adapters/driving"
	user "mkfolder.dev/wire-playground/internal/user/core"
)

type Container struct {
	DB *database.Postgres

	UserService *user.UserService
	UserAdapter *userDriving.HTTPAdapter

	ProfileService *profile.ProfileService
	ProfileAdapter *profileDriving.HTTPAdapter

	PointService *point.PointService
	PointAdapter *pointDriving.HTTPAdapter
}

func NewContainer(
	db *database.Postgres,
	userService *user.UserService,
	userAdapter *userDriving.HTTPAdapter,
	profileService *profile.ProfileService,
	profileAdapter *profileDriving.HTTPAdapter,
	pointService *point.PointService,
	pointAdapter *pointDriving.HTTPAdapter,
) Container {
	return Container{
		DB:             db,
		UserService:    userService,
		UserAdapter:    userAdapter,
		ProfileService: profileService,
		ProfileAdapter: profileAdapter,
		PointService:   pointService,
		PointAdapter:   pointAdapter,
	}
}

func InitializeContainer(router fiber.Router) Container {
	wire.Build(
		database.NewPostgres,
		user.NewUserService,
		userDriven.NewPostgresRepository,
		userDriving.NewHTTPAdapter,
		point.NewPointService,
		pointDriven.NewUserService,
		pointDriving.NewHTTPAdapter,
		profile.NewProfileService,
		profileDriven.NewUserService,
		profileDriving.NewHTTPAdapter,
		NewContainer,

		wire.Bind(new(user.UserRepository), new(*userDriven.PostgresRepository)),
		wire.Bind(new(point.UserFetcher), new(*pointDriven.UserService)),
		wire.Bind(new(profile.UserFetcher), new(*profileDriven.UserService)),
		wire.Bind(new(profile.PointsFetcher), new(*point.PointService)),
	)
	return Container{}
}
