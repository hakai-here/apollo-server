package handler

import "github.com/gofiber/fiber/v2"

func HomePage(c *fiber.Ctx) error {
	return c.Render("home", fiber.Map{
		"title": "Squid Services",
	})
}

func RegisterPage(c *fiber.Ctx) error {
	return c.Render("register", fiber.Map{
		"title": "Welcome to Squid service",
	})
}

func LoginPage(c *fiber.Ctx) error {
	return c.Render("login", fiber.Map{
		"title": "Welcome back to Squid service",
	})
}

func VerifyPage(c *fiber.Ctx) error {
	return c.Render("verify", fiber.Map{
		"title": "Successfully registered. Verify your email",
	})
}

func Page404(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotFound).Render("404", fiber.Map{
		"title": "404 | Page not found",
	})
}

func Error(c *fiber.Ctx) error {
	return c.Status(fiber.StatusInternalServerError).Render("error", fiber.Map{
		"title": "Error Squid is not well",
	})
}
