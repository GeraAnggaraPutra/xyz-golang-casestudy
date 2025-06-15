package application

import (
	"io"
	"strconv"

	"kredit-plus/src/constant"
	"kredit-plus/src/domain/storage/payload"
	"kredit-plus/src/domain/storage/service"
	"kredit-plus/src/kernel"
	"kredit-plus/toolkit/logger"

	"github.com/labstack/echo/v4"
)

func getFileApp(svc *service.Service) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		var request payload.GetFileRequest
		if err = c.Bind(&request); err != nil {
			err = logger.PrintNewError(err, constant.ErrFailedParseRequest)
			return
		}

		if err := c.Validate(request); err != nil {
			logger.PrintError(err, "error validation get file request")
			return kernel.ResponseErrorValidate(c, err)
		}

		rc, err := svc.GetFileService(c.Request().Context(), request)
		if err != nil {
			return kernel.ResponseError(c, err, msgFailedOpenFile)
		}
		defer rc.Close()

		size := rc.Attrs.Size

		sizeStr := strconv.FormatInt(size, 10)

		c.Response().Header().Set(echo.HeaderContentType, rc.Attrs.ContentType)
		c.Response().Header().Set(echo.HeaderContentLength, sizeStr)

		if _, err := io.Copy(c.Response().Writer, rc); err != nil {
			return kernel.ResponseError(c, err, msgFailedOpenFile)
		}

		return
	}
}
