package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/syamsv/apollo-server/common/handler"
	"github.com/syamsv/apollo-server/common/middleware"
)

func MountGetRoutes(c fiber.Router) {
	c.Get("/", handler.HomePage)
	c.Get("/error", handler.Error)
	auth := c.Group("/uas")
	auth.Use(middleware.AlreadyCookiePresent)
	auth.Get("/register", handler.RegisterPage)
	auth.Post("/register", handler.AuthRegistersHandler)

	auth.Get("/login", handler.LoginPage)
	auth.Post("/login", handler.AuthLoginHandler)

	auth.Get("/verify", handler.VerifyPage)

	auth.Get("/activate/:token", handler.AuthActivateHandler)

	c.Get("/*", handler.Page404)
}
