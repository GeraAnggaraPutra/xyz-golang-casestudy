package application

import (
	"net/http"

	"kredit-plus/src/constant"
	"kredit-plus/src/domain/storage/payload"
	"kredit-plus/src/domain/storage/service"
	"kredit-plus/src/kernel"
	"kredit-plus/toolkit/logger"

	"github.com/labstack/echo/v4"
)

func createFileUploadApp(svc *service.Service) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		var request payload.CreateFileUploadRequest
		if err = c.Bind(&request); err != nil {
			err = logger.PrintNewError(err, constant.ErrFailedParseRequest)
			return
		}

		request.File, err = c.FormFile("file")
		if err != nil {
			return kernel.ResponseError(c, err, "Failed to get file")
		}

		if err := c.Validate(request); err != nil {
			logger.PrintError(err, "error validation upload file request")
			return kernel.ResponseErrorValidate(c, err)
		}

		data, err := svc.CreateFileUploadService(c.Request().Context(), request)
		if err != nil {
			return kernel.ResponseError(c, err, msgFailedFileUpload)
		}

		return kernel.ResponseData(c, kernel.ResponsePayload{
			Code:    http.StatusOK,
			Data:    payload.ToReadFileUploadResponse(data),
			Message: msgSuccessFileUpload,
		})
	}
}
