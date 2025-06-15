package application

import (
	"net/http"

	"kredit-plus/src/domain/auth/service"
	"kredit-plus/src/handler/auth"
	"kredit-plus/src/kernel"

	"github.com/labstack/echo/v4"
)

func logoutApp(svc *service.Service) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		ah, err := auth.GetAuth(c)
		if err != nil {
			return
		}

		err = svc.LogoutService(c.Request().Context(), ah.GetClaims())
		if err != nil {
			return kernel.ResponseError(c, err, msgFailedLogout)
		}

		return kernel.ResponseData(c, kernel.ResponsePayload{
			Code:    http.StatusOK,
			Data:    nil,
			Message: msgSuccessLogout,
		})
	}
}
