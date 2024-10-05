# goFiber-BasicAuth
* Basic Authentication middleware for Fiber that provides an HTTP: https://docs.gofiber.io/api/middleware/basicauth
* Session: https://docs.gofiber.io/api/middleware/session
* Cache: https://docs.gofiber.io/api/middleware/cache

## Features
* Standard Go Project Layout: https://github.com/golang-standards/project-layout
* Database PostgreSQL + GORM + .env + users_db, posts_db
#### Create .env file:
```ENV
DB_USER=postgres
DB_PASSWORD=postgres
DB_HOST=localhost
DB_PORT=5432
DB_NAME=postgres
DB_SSLMODE=disable
```

## authConfig
```go
authConfig := basicauth.Config{
	Users: map[string]string{
		"admin": "password123",},
}
```
file: /internal/handlers/user.go
```go
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
```

## Getting Started
Github: https://github.com/asakew/goFiber-BasicAuth
```bash
git clone https://github.com/asakew/goFiber-BasicAuth
```

## Updates go.mod and go.sum and running 
```bash
go mod tidy
go run app/main.go
```

## Stop Running
```bash
CTRL + C # Windows
CTRL + Break # Linux
Control + C # Mac
```

## Build
```bash
go mod tidy # updates go.mod and go.sum
go mod download # downloads all the dependencies
go build # builds the binary
```