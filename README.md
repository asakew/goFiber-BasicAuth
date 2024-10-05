# goFiber-BasicAuth
* Basic Authentication middleware for Fiber that provides an HTTP: https://docs.gofiber.io/api/middleware/basicauth/
* Session: https://docs.gofiber.io/api/middleware/session/
* Cache: https://docs.gofiber.io/api/middleware/cache/

## Features
* Standard Go Project Layout: https://github.com/golang-standards/project-layout/
* Database PostgreSQL GORM: + users_db, posts_db

## authConfig
```go
authConfig := basicauth.Config{
	Users: map[string]string{
		"admin": "password123",
		"testUser": "testUser",
	},
}
```

## Getting Started
Github: https://github.com/asakew/goFiber-BasicAuth/
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