package handler

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/syamsv/apollo-server/common/controllers"
	"github.com/syamsv/apollo-server/pkg/models"
	"github.com/syamsv/apollo-server/pkg/schema"
)

func AuthLoginHandler(c *fiber.Ctx) error {
	loginUser := new(schema.Login)
	if err := c.BodyParser(loginUser); err != nil {
		fmt.Println("1", err)
		return c.Status(fiber.StatusBadRequest).Redirect("/error")
	}
	validate := validator.New()
	if err := validate.Struct(loginUser); err != nil {
		fmt.Println("2", err)
		return c.Status(fiber.StatusBadRequest).Redirect("/error")
	}
	sessionId, err := controllers.LoginUser(loginUser.Email, loginUser.Password,c.IP())
	if err != nil {
		fmt.Println("3", err)
		return c.Status(fiber.StatusBadRequest).Redirect("/error")
	}
	c.Cookie(&fiber.Cookie{
		Name:  "Authorization",
		Value: sessionId,
	})
	return c.Redirect("/")
}

func AuthRegistersHandler(c *fiber.Ctx) error {
	user := new(models.Users)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).Redirect("/error")
	}
	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		return c.Status(fiber.StatusBadRequest).Redirect("/error")
	}
	if err := controllers.CreateUser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).Redirect("/error")
	}

	return c.Status(fiber.StatusCreated).Redirect("/uas/verify")
}

func AuthActivateHandler(c *fiber.Ctx) error {
	token := c.Params("token")
	if token == "" {
		return c.Status(fiber.StatusBadRequest).Render("activate", fiber.Map{
			"title":    "Error in Activation",
			"activate": false,
		})
	}

	if err := controllers.ActivateUserController(token); err != nil {
		return c.Status(fiber.StatusInternalServerError).Render("activate", fiber.Map{
			"title":    "Error in Activation",
			"activate": false,
		})
	}

	return c.Render("activate", fiber.Map{
		"title":    "Successful Activation",
		"activate": true,
	})
}
