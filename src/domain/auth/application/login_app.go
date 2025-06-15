package application

import (
	"net/http"

	"kredit-plus/src/constant"
	"kredit-plus/src/domain/auth/payload"
	"kredit-plus/src/domain/auth/service"
	"kredit-plus/src/kernel"
	"kredit-plus/toolkit/logger"

	"github.com/labstack/echo/v4"
)

func loginApp(svc *service.Service) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		var request payload.LoginRequest
		if err = c.Bind(&request); err != nil {
			err = logger.PrintNewError(err, constant.ErrFailedParseRequest)
			return
		}

		if err := c.Validate(request); err != nil {
			logger.PrintError(err, "error validation login request")
			return kernel.ResponseErrorValidate(c, err)
		}

		data, user, err := svc.LoginService(c.Request().Context(), request)
		if err != nil {
			return kernel.ResponseError(c, err, msgFailedLogin)
		}

		return kernel.ResponseData(c, kernel.ResponsePayload{
			Code:    http.StatusOK,
			Data:    payload.ToSessionResponse(data, user),
			Message: msgSuccessLogin,
		})
	}
}
