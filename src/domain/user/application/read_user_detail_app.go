package application

import (
	"net/http"

	"kredit-plus/src/constant"
	"kredit-plus/src/domain/user/payload"
	"kredit-plus/src/domain/user/service"
	"kredit-plus/src/kernel"
	"kredit-plus/toolkit/logger"

	"github.com/labstack/echo/v4"
)

func readUserDetailApp(svc *service.Service) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		var request payload.ReadUserDetailRequest
		if err = c.Bind(&request); err != nil {
			err = logger.PrintNewError(err, constant.ErrFailedParseRequest)
			return
		}

		if err := c.Validate(request); err != nil {
			logger.PrintError(err, "error validation get user detail request")
			return kernel.ResponseErrorValidate(c, err)
		}

		data, err := svc.ReadUserDetailService(c.Request().Context(), request)
		if err != nil {
			return kernel.ResponseError(c, err, msgFailedGetUserDetail)
		}

		return kernel.ResponseData(c, kernel.ResponsePayload{
			Code:    http.StatusOK,
			Data:    payload.ToUserResponse(data),
			Message: msgSuccessGetUserDetail,
		})
	}
}
