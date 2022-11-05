package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/devnura/go-echo-rest-api/config"
	"github.com/devnura/go-echo-rest-api/controller"
	"github.com/devnura/go-echo-rest-api/helper"
	"github.com/devnura/go-echo-rest-api/repository"
	"github.com/devnura/go-echo-rest-api/service"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                  = config.DatabaseConnection()
	authRepository repository.AuthRepository = repository.NewAuthRepository(db)

	authService    service.AuthService       = service.NewAuthService(authRepository)
	authController controller.AuthController = controller.NewAuthController(authService)
)

func main() {
	defer config.CloseConnection(db)
	var (
		NAME = "DevNura"
	)

	e := echo.New()

	e.Use(
		middleware.Logger(),
		middleware.Recover(),
		middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"*"},
			AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
		}),
		middleware.LoggerWithConfig(middleware.LoggerConfig{
			Format: fmt.Sprintf("\n%s | ${host} | ${time_custom} | ${status} | ${latency_human} | ${remote_ip} | ${method} | ${uri} ",
				NAME,
			),
			CustomTimeFormat: "2006/01/02 15:04:05",
			Output:           os.Stdout,
		}),
	)

	e.Validator = &helper.CustomValidator{Validator: validator.New()}

	// g := e.Group("/v1/auth")
	// g.POST("/login", authController)
	e.Logger.Fatal(e.Start(":1323"))
}
