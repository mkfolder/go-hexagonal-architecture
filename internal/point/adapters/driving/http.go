package driving

import (
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"mkfolder.dev/wire-playground/internal/point/core"
)

type HTTPAdapter struct {
	svc *core.PointService
}

func NewHTTPAdapter(router fiber.Router, svc *core.PointService) *HTTPAdapter {
	adapter := &HTTPAdapter{svc: svc}
	adapter.registerRoutes(router)

	return adapter
}

func (a *HTTPAdapter) registerRoutes(router fiber.Router) {
	router.Get("/leaderboard", a.GetLeaderboard)
}

func (a *HTTPAdapter) GetLeaderboard(c fiber.Ctx) error {
	filterID := c.Query("filter_id")
	parsedFilterID, err := uuid.Parse(filterID)
	if err != nil {
		return err
	}

	leaderboard, err := a.svc.GetLeaderboard(parsedFilterID)
	if err != nil {
		return err
	}

	return c.JSON(leaderboard)
}
