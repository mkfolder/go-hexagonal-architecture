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
	"mkfolder.dev/wire-playground/internal/shared"
	userDriving "mkfolder.dev/wire-playground/internal/user/adapters/driving"
	userBootstrap "mkfolder.dev/wire-playground/internal/user/bootstrap"
	user "mkfolder.dev/wire-playground/internal/user/core"
)

type Container struct {
	DB *gorm.DB

	UserService        *user.UserService
	UserHTTPAdapter    *userDriving.HTTPAdapter
	UserHexagonAdapter *userDriving.HexagonAdapter

	ProfileService     *profile.ProfileService
	ProfileHTTPAdapter *profileDriving.HTTPAdapter

	PointService     *point.PointService
	PointHTTPAdapter *pointDriving.HTTPAdapter
}

func NewContainer(
	db *gorm.DB,
	userService *user.UserService,
	userHTTPAdapter *userDriving.HTTPAdapter,
	userHexagonAdapter *userDriving.HexagonAdapter,
	profileService *profile.ProfileService,
	profileHTTPAdapter *profileDriving.HTTPAdapter,
	pointService *point.PointService,
	pointHTTPAdapter *pointDriving.HTTPAdapter,
) Container {
	return Container{
		DB:                 db,
		UserService:        userService,
		UserHTTPAdapter:    userHTTPAdapter,
		UserHexagonAdapter: userHexagonAdapter,
		ProfileService:     profileService,
		ProfileHTTPAdapter: profileHTTPAdapter,
		PointService:       pointService,
		PointHTTPAdapter:   pointHTTPAdapter,
	}
}

func InitializeContainer(router fiber.Router) Container {
	wire.Build(
		database.NewGormDB,
		userBootstrap.UserSet,
		pointBootstrap.PointSet,
		profileBootstrap.ProfileSet,
		NewContainer,

		userDriving.NewHexagonAdapter,
		wire.Bind(new(shared.UserAdapter), new(*userDriving.HexagonAdapter)),
	)
	return Container{}
}
