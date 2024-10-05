package routes

import (
	"github.com/gofiber/fiber/v2"
)

func HTMLRendering(app *fiber.App) {
	app.Get("/test_template", func(c *fiber.Ctx) error {
		// Render index template
		return c.Render("index", fiber.Map{
			"Title":       "Go Fiber Template Example",
			"Description": "An example template",
			"Greeting":    "Hello, world!",
		})
	})
}
