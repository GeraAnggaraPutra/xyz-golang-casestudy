package kernel

import (
	"errors"
	"net/http"

	"kredit-plus/src/constant"
	"kredit-plus/src/handler/validator"
	"kredit-plus/src/util"

	"github.com/labstack/echo/v4"
)

type responseErrorPayload struct {
	Error   interface{} `json:"error,omitempty"`
	Message string      `json:"message"`
}

func ResponseErrorValidate(c echo.Context, err error) error {
	return c.JSON(http.StatusBadRequest, responseErrorPayload{
		Error:   validator.ValidationErrors(err),
		Message: constant.ErrMsgValidate,
	})
}

func ResponseError(c echo.Context, err error, msg string) error {
	var echoError *echo.HTTPError

	// if *echo.HTTPError, let echo middleware handles it
	if errors.As(err, &echoError) {
		return err
	}

	e := formatError(err)
	if e != nil {
		return c.JSON(http.StatusBadRequest, responseErrorPayload{
			Error:   e,
			Message: msg,
		})
	}

	return err
}

func formatError(err error) (e map[string]string) {
	switch {
	case errors.Is(err, constant.ErrPasswordIncorrect):
		e = map[string]string{"password": util.CapitalFirstLetter(err.Error())}
	case errors.Is(err, constant.ErrAccountNotFound) || errors.Is(err, constant.ErrEmailAlreadyExists):
		e = map[string]string{"email": util.CapitalFirstLetter(err.Error())}
	}

	return
}
