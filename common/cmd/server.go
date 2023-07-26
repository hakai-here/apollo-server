package cmd

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/encryptcookie"
	"github.com/gofiber/template/html/v2"
	"github.com/syamsv/apollo-server/common/router"
	"github.com/syamsv/apollo-server/config"
)

func StartServer() {
	engine := html.New("./template", ".html")
	app := fiber.New(fiber.Config{
		// Prefork:       true, // Enable Prefork in prod only
		Views:         engine,
		ViewsLayout:   "index",
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Apollo Server",
		AppName:       "App Name",
	})
	app.Static("/css", "./static/css")
	app.Static("/js", "./static/js")
	app.Static("/image", "./static/images")
	app.Use(encryptcookie.New(encryptcookie.Config{
		Key: encryptcookie.GenerateKey(),
	}))

	app.Use(cors.New())
	api := app.Group("/app")
	home := app.Group("/")
	router.MountDashboard(api)
	router.MountGetRoutes(home)

	if err := app.Listen(config.SERVER_PORT); err != nil {
		log.Fatal("[SERVER] Error starting server: ", err)
	}
}
