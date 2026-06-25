//go:build wireinject

package bootstrap

import (
	"github.com/gofiber/fiber/v3"
	"github.com/google/wire"
	"gorm.io/gorm"

	"mkfolder.dev/wire-playground/internal/database"
	pointDriving "mkfolder.dev/wire-playground/internal/point/adapters/driving"
	pointBootstrap "mkfolder.dev/wire-playground/internal/point/bootstrap"
	point "mkfolder.dev/wire-playground/internal/point/core"
	profileDriving "mkfolder.dev/wire-playground/internal/profile/adapters/driving"
	profileBootstrap "mkfolder.dev/wire-playground/internal/profile/bootstrap"
	profile "mkfolder.dev/wire-playground/internal/profile/core"
	userDriving "mkfolder.dev/wire-playground/internal/user/adapters/driving"
	userBootstrap "mkfolder.dev/wire-playground/internal/user/bootstrap"
	user "mkfolder.dev/wire-playground/internal/user/core"
)

type Container struct {
	DB *gorm.DB

	UserService *user.UserService
	UserAdapter *userDriving.HTTPAdapter

	ProfileService *profile.ProfileService
	ProfileAdapter *profileDriving.HTTPAdapter

	PointService *point.PointService
	PointAdapter *pointDriving.HTTPAdapter
}

func NewContainer(
	db *gorm.DB,
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
		database.NewGormDB,
		userBootstrap.UserSet,
		pointBootstrap.PointSet,
		profileBootstrap.ProfileSet,
		NewContainer,
	)
	return Container{}
}
