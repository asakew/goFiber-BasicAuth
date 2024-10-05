# goFiber-BasicAuth
**Basic Authentication middleware for Fiber that provides an HTTP basic authentication.**
Full Docs: https://docs.gofiber.io/api/middleware/basicauth

## Features
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

## updates go.mod and go.sum
```bash
go mod tidy
```

## Build
```bash
go mod tidy # updates go.mod and go.sum
go mod download # downloads all the dependencies
go build # builds the binary
./main # runs the binary
```