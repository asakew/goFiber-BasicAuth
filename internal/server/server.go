package server

import (
	"BasicAuth/internal/pkg/database"
	"BasicAuth/internal/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html/v2"
)

func Run() {
	engine := html.New("./web/templates", ".html") // Load Template

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	database.InitDB()

	app.Static("/web/assets", "./web/assets")

	// Basic Auth
	authConfig := basicauth.Config{
		Users: map[string]string{
			"admin": "password123",
		},
	}

	app.Use(basicauth.New(authConfig))

	// Middleware logging
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cors.New())

	routes.UserHandlers(app)

	err := app.Listen(":3000")
	if err != nil {
		return
	}
}
