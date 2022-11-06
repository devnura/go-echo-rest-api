package handler

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/devnura/go-echo-rest-api/dto"
	"github.com/devnura/go-echo-rest-api/helper"
	"github.com/devnura/go-echo-rest-api/service"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	service service.AuthService
}

func NewAuthHandler(s service.AuthService) *AuthHandler {
	return &AuthHandler{service: s}
}

func (h *AuthHandler) Login(c echo.Context) error {
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var req dto.LoginDTO
	defer cancel()

	validate := validator.New()
	// validate request body
	if err := c.Bind(&req); err != nil {
		fmt.Print(err)
		response := helper.BuildErrorResponse("Bad Request", err.Error(), helper.EmptyObj{})
		return c.JSON(http.StatusBadRequest, response)
	}

	//use the validator library to validate required fields
	if err := validate.Struct(&req); err != nil {
		response := helper.BuildErrorResponse("Bad Request", err.Error(), helper.EmptyObj{})
		return c.JSON(http.StatusBadRequest, response)
	}

	res, err := h.service.FindByEmail(c, &req)
	if err != nil {
		fmt.Print(err)
		response := helper.BuildErrorResponse("Bad Request", err.Error(), helper.EmptyObj{})
		return c.JSON(http.StatusNotFound, response)
	}

	response := helper.BuildResponse(true, "Success", res)
	return c.JSON(http.StatusOK, response)
}
