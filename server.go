package main

import (
	"github.com/devnura/go-echo-rest-api/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Hello() string {
	return "hello world"
}

func main() {
	var server = echo.New()
	server.Use(middleware.Logger())
	server.Use(middleware.Recover())
	routes.SetupLogger(server)
	// routes.SetupGlobalErrorHandler(server)
	routes.SetupMiddleware(server)
	routes.SetupRoute(server)

	server.Logger.Fatal(server.Start(":8080"))
}
