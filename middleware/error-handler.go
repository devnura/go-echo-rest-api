package middleware

import (
	"errors"

	"net/http"
	"strings"

	"github.com/devnura/go-echo-rest-api/common/log"
	"github.com/devnura/go-echo-rest-api/config"
	"github.com/devnura/go-echo-rest-api/constants"
	"github.com/devnura/go-echo-rest-api/transfer"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GlobalErrorHandler(e *echo.Echo) func(error, echo.Context) {
	return func(err error, c echo.Context) {

		if c.Response().Committed {
			return
		}

		if httpError, ok := err.(*transfer.CustomError); ok {
			rejectErrMessage(c, httpError)
			return
		}

		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			httpError := transfer.NewCustomError(constants.ERR002)
			rejectErrMessage(c, httpError)
		case errors.Is(err, gorm.ErrInvalidData):
			httpError := transfer.NewCustomError(constants.ERR003)
			rejectErrMessage(c, httpError)
		case errors.Is(err, gorm.ErrInvalidValueOfLength):
			httpError := transfer.NewCustomError(constants.ERR004)
			rejectErrMessage(c, httpError)
		default:
			log.InfoWithID(c, err.Error())
			httpError := transfer.NewCustomError(constants.ERR999)
			rejectErrMessage(c, httpError)

		}

	}
}

func rejectErrMessage(c echo.Context, err *transfer.CustomError) {
	code := strings.ToLower(err.Code)
	err.Message = config.ERRConfig[code]
	c.Logger().Error(c.JSON(http.StatusOK, err))
}
