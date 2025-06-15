package api

import (
	"net/http"
	"os"
	"strconv"

	"kredit-plus/src/constant"
	"kredit-plus/toolkit/logger"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

func errorHandler() echo.HTTPErrorHandler {
	return func(err error, c echo.Context) {
		var echoError *echo.HTTPError

		// if *echo.HTTPError, let echokit middleware handles it
		if errors.As(err, &echoError) {
			_ = c.JSON(echoError.Code, echoError)
		} else {
			appDebug, _ := strconv.ParseBool(os.Getenv("APP_DEBUG"))

			switch {
			case !appDebug:
				_ = c.JSON(constant.ErrUnknownSource.Code, constant.ErrUnknownSource)
			default:
				_ = c.JSON(http.StatusInternalServerError, map[string]string{
					"message": err.Error(),
				})
			}
		}

		mappingErrorLogger(err, c)
	}
}

func mappingErrorLogger(err error, c echo.Context) {
	if c.Response().Status >= http.StatusInternalServerError {
		logger.PrintWarn(err.Error(),
			"path", c.Request().URL.Path,
		)
	}
}
