package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/syamsv/apollo-server/common/session"
)

func AlreadyCookiePresent(c *fiber.Ctx) error {
	if c.Cookies("Authorization") != "" {
		return c.Redirect("/")
	}
	return c.Next()
}

func AuthMiddleware(c *fiber.Ctx) error {
	data := c.Cookies("Authorization")
	if data == "" {
		return c.Redirect("/error")
	}
	sessionData, err := session.GetSession(data)
	if err != nil {
		return c.Redirect("/error")
	}
	c.Locals("userId", sessionData.Id)
	c.Locals("Firstname", sessionData.FirstName)
	c.Locals("Lastname", sessionData.LastName)
	return c.Next()
}
