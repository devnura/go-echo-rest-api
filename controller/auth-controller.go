package controller

import (
	"net/http"

	"github.com/devnura/go-echo-rest-api/dto"
	"github.com/devnura/go-echo-rest-api/helper"
	"github.com/devnura/go-echo-rest-api/service"
	"github.com/labstack/echo/v4"
)

type AuthController interface {
	Login(ctx echo.Context)
	// Register(ctx *echo.Context)
}

type authController struct {
	authService service.AuthService
}

func NewAuthController(authService service.AuthService) AuthController {
	return &authController{
		authService: authService,
	}
}

func (controller *authController) Login(ctx echo.Context) {
	var loginDTO dto.LoginDTO
	errDTO := ctx.Bind(&loginDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
	}
}
