package handlers

import (
	"BasicAuth/internal/database"
	"BasicAuth/internal/models"
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

func Register(c *fiber.Ctx) error {
	return c.Render("register", nil)
}

func HandleRegister(c *fiber.Ctx) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	if username == "" || password == "" {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	// Save User to Database using GORM
	user := models.User{Username: username, Password: password}
	result := database.UsersDB.Create(&user)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": result.Error.Error(),
		})
	}

	return c.Redirect("/login", fiber.StatusSeeOther)
}

func Logout(c *fiber.Ctx) error {
	// Clear the session cookie
	c.ClearCookie("session") // Adjust "session" to your actual cookie name

	// Optionally log the logout action
	log.Println("User logged out")

	// Redirect to the login page
	return c.Redirect("/login", fiber.StatusSeeOther)
}
