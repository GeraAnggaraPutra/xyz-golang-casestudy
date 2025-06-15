package application

import (
	"net/http"

	"kredit-plus/src/constant"
	"kredit-plus/src/domain/customer/payload"
	"kredit-plus/src/domain/customer/service"
	"kredit-plus/src/kernel"
	"kredit-plus/src/handler/auth"
	"kredit-plus/toolkit/logger"

	"github.com/labstack/echo/v4"
)

func deleteCustomerApp(svc *service.Service) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		var request payload.DeleteCustomerRequest
		if err = c.Bind(&request); err != nil {
			err = logger.PrintNewError(err, constant.ErrFailedParseRequest)
			return
		}

		if err := c.Validate(request); err != nil {
			logger.PrintError(err, "error validation delete customer request")
			return kernel.ResponseErrorValidate(c, err)
		}

		ah, err := auth.GetAuth(c)
		if err != nil {
			logger.WithContext(c.Request().Context()).Error(err, "error get auth handler")
			return
		}

		err = svc.DeleteCustomerService(c.Request().Context(), request, ah.GetClaims().UserGUID)
		if err != nil {
			return kernel.ResponseError(c, err, msgFailedDeleteCustomer)
		}

		return kernel.ResponseData(c, kernel.ResponsePayload{
			Code:    http.StatusOK,
			Data:    nil,
			Message: msgSuccessDeleteCustomer,
		})
	}
}
