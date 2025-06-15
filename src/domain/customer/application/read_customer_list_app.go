package application

import (
	"net/http"

	"kredit-plus/src/constant"
	"kredit-plus/src/domain/customer/payload"
	"kredit-plus/src/domain/customer/service"
	"kredit-plus/src/kernel"
	"kredit-plus/toolkit/logger"

	"github.com/labstack/echo/v4"
)

func readCustomerListApp(svc *service.Service) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		var request payload.ReadCustomerListRequest
		if err = c.Bind(&request); err != nil {
			err = logger.PrintNewError(err, constant.ErrFailedParseRequest)
			return
		}

		if err := c.Validate(request); err != nil {
			logger.PrintError(err, "error validation get customer list request")
			return kernel.ResponseErrorValidate(c, err)
		}

		request.Init()

		data, totalData, err := svc.ReadCustomerListService(c.Request().Context(), request)
		if err != nil {
			return kernel.ResponseError(c, err, msgFailedGetCustomerList)
		}

		return kernel.ResponsePaginate(c, request.PaginationPayload, totalData, kernel.ResponsePayload{
			Code:    http.StatusOK,
			Data:    payload.ToReadListCustomerResponses(data),
			Message: msgSuccessGetCustomerList,
		})
	}
}
