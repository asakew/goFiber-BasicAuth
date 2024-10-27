package handlers

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

func FrontPage(c *fiber.Ctx) error {
	return c.Render("index", nil)
}

func Login(c *fiber.Ctx) error {
	return c.Render("login", nil)
}

func HandleLogin(c *fiber.Ctx) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	// Authentication logic (use database to validate user)
	// Placeholder logic for demonstration
	if username == "admin" && password == "password123" {
		return c.Redirect("/", fiber.StatusSeeOther)
	}
	return c.SendStatus(fiber.StatusUnauthorized)
}

func Logout(c *fiber.Ctx) error {
	// Clear the session cookie
	c.ClearCookie("session") // Adjust "session" to your actual cookie name

	// Optionally log the logout action
	log.Println("User logged out")

	// Redirect to the login page
	return c.Redirect("/login", fiber.StatusSeeOther)
}
