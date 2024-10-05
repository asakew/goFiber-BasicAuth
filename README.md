# goFiber-BasicAuth
**Basic Authentication middleware for Fiber that provides an HTTP basic authentication.**
Full Docs: https://docs.gofiber.io/api/middleware/basicauth

## Features
* Database PostgreSQL + GORM + .env 
#### Create .env file:
```ENV
DB_USER=postgres
DB_PASSWORD=postgres
DB_HOST=localhost
DB_PORT=5432
DB_NAME=postgres
DB_SSLMODE=disable
```

## Getting Started
Github: https://github.com/asakew/goFiber-BasicAuth
```bash
git clone https://github.com/asakew/goFiber-BasicAuth
```

## Server
```bash
go mod tidy # updates go.mod and go.sum
go run app/main.go # runs the binary
```



## Build
```bash
go mod tidy # updates go.mod and go.sum
go mod download # downloads all the dependencies
go build # builds the binary
./main # runs the binary
```