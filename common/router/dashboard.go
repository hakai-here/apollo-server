package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/syamsv/apollo-server/common/middleware"
)

func MountDashboard(dash fiber.Router) {
	dash.Use(middleware.AuthMiddleware)
	dash.Get("/", func(c *fiber.Ctx) error {
		return c.Render("dashboard", fiber.Map{})
	})
}
