package routes

import (
	"BasicAuth/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func UserHandlers(app *fiber.App) {
	app.Get("/", handlers.FrontPage)
	app.Get("/login", handlers.Login)
	app.Post("/login", handlers.HandleLogin)
	app.Get("/register", handlers.Register)
	app.Post("/register", handlers.HandleRegister)
	app.Get("/logout", handlers.Logout)
}
