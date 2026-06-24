package driving

import (
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"mkfolder.dev/wire-playground/internal/profile/core"
)

type HTTPAdapter struct {
	svc *core.ProfileService
}

func NewHTTPAdapter(router fiber.Router, svc *core.ProfileService) *HTTPAdapter {
	adapter := &HTTPAdapter{svc: svc}
	adapter.registerRoutes(router)
	return adapter
}

func (h *HTTPAdapter) registerRoutes(router fiber.Router) {
	router.Get("/profiles/:id", h.GetProfile)
}

func (h *HTTPAdapter) GetProfile(c fiber.Ctx) error {
	id := c.Params("id")
	parsedID, err := uuid.Parse(id)
	if err != nil {
		return fiber.ErrBadRequest
	}

	profile, err := h.svc.FindByID(parsedID)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	return c.JSON(profile)
}
