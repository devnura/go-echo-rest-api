package routes

import (
	"github.com/devnura/go-echo-rest-api/config/db"
	"github.com/devnura/go-echo-rest-api/handler"
	_middleware "github.com/devnura/go-echo-rest-api/middleware"
	"github.com/devnura/go-echo-rest-api/repository"
	"github.com/devnura/go-echo-rest-api/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func SetupGlobalErrorHandler(e *echo.Echo) {
	e.HTTPErrorHandler = _middleware.GlobalErrorHandler(e)
}

func SetupLogger(e *echo.Echo) {
	e.Logger.SetLevel(log.DEBUG)
}

func SetupMiddleware(e *echo.Echo) {
	// global or root middleware
	e.Use(middleware.RequestID())

}

func SetupRoute(e *echo.Echo) {

	// repository
	var (
		gormDB = db.NewMysqlDB()

		authRepo repository.AuthRepository = repository.NewAuthRepository(gormDB)
	)

	// service
	var (
		authService = service.NewAuthService(authRepo)
	)

	// handler
	var (
		authHandler = handler.NewAuthHandler(authService)
	)

	g := e.Group("/api/v1/auth")
	g.POST("/login", authHandler.Login)
}
