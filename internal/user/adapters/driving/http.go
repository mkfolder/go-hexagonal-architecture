package driving

import (
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"mkfolder.dev/wire-playground/internal/user/core"
)

type HTTPAdapter struct {
	svc *core.UserService
}

func NewHTTPAdapter(router fiber.Router, svc *core.UserService) *HTTPAdapter {
	adapter := &HTTPAdapter{svc: svc}
	adapter.registerRoutes(router)
	return adapter
}

func (h *HTTPAdapter) registerRoutes(router fiber.Router) {
	router.Get("/users/:id", h.GetUser)
	router.Post("/users", h.CreateUser)
}

func (h *HTTPAdapter) GetUser(c fiber.Ctx) error {
	id := c.Params("id")
	parsedID, err := uuid.Parse(id)
	if err != nil {
		return fiber.ErrBadRequest
	}

	user, err := h.svc.FindByID(parsedID)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	return c.JSON(user)
}

func (h *HTTPAdapter) CreateUser(c fiber.Ctx) error {
	var user core.User
	if err := c.Bind().Body(&user); err != nil {
		return fiber.ErrBadRequest
	}

	if err := user.Validate(); err != nil {
		return fiber.ErrBadRequest
	}

	if err := h.svc.Create(&user); err != nil {
		return fiber.ErrInternalServerError
	}

	return c.JSON(user)
}
