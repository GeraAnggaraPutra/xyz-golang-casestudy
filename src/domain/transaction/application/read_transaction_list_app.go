package application

import (
	"net/http"

	"kredit-plus/src/constant"
	"kredit-plus/src/domain/transaction/payload"
	"kredit-plus/src/domain/transaction/service"
	"kredit-plus/src/kernel"
	"kredit-plus/toolkit/logger"

	"github.com/labstack/echo/v4"
)

func readTransactionListApp(svc *service.Service) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		var request payload.ReadTransactionListRequest
		if err = c.Bind(&request); err != nil {
			err = logger.PrintNewError(err, constant.ErrFailedParseRequest)
			return
		}

		if err := c.Validate(request); err != nil {
			logger.PrintError(err, "error validation get transaction list request")
			return kernel.ResponseErrorValidate(c, err)
		}

		request.Init()

		data, totalData, err := svc.ReadTransactionListService(c.Request().Context(), request)
		if err != nil {
			return kernel.ResponseError(c, err, msgFailedGetTransactionList)
		}

		return kernel.ResponsePaginate(c, request.PaginationPayload, totalData, kernel.ResponsePayload{
			Code:    http.StatusOK,
			Data:    payload.ToReadListTransactionResponses(data),
			Message: msgSuccessGetTransactionList,
		})
	}
}
