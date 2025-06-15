package application

import (
	"net/http"

	"kredit-plus/src/constant"
	"kredit-plus/src/domain/user/payload"
	"kredit-plus/src/domain/user/service"
	"kredit-plus/src/kernel"
	"kredit-plus/src/handler/auth"
	"kredit-plus/toolkit/logger"

	"github.com/labstack/echo/v4"
)

func deleteUserApp(svc *service.Service) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		var request payload.DeleteUserRequest
		if err = c.Bind(&request); err != nil {
			err = logger.PrintNewError(err, constant.ErrFailedParseRequest)
			return
		}

		if err := c.Validate(request); err != nil {
			logger.PrintError(err, "error validation delete user request")
			return kernel.ResponseErrorValidate(c, err)
		}

		ah, err := auth.GetAuth(c)
		if err != nil {
			logger.WithContext(c.Request().Context()).Error(err, "error get auth handler")
			return
		}

		err = svc.DeleteUserService(c.Request().Context(), request, ah.GetClaims().UserGUID)
		if err != nil {
			return kernel.ResponseError(c, err, msgFailedDeleteUser)
		}

		return kernel.ResponseData(c, kernel.ResponsePayload{
			Code:    http.StatusOK,
			Data:    nil,
			Message: msgSuccessDeleteUser,
		})
	}
}
