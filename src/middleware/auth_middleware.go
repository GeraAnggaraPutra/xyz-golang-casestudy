package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"kredit-plus/src/constant"
	"kredit-plus/src/handler/auth"
	"kredit-plus/src/handler/jwt"
	"kredit-plus/toolkit/logger"
)

type EnsureToken struct {
	auth *auth.Auth
}

func NewEnsureToken(cache redis.UniversalClient, db *gorm.DB) *EnsureToken {
	ah := auth.NewAuth(cache, db)

	return &EnsureToken{
		auth: ah,
	}
}

func (et *EnsureToken) ValidateToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token, err := parseHeaderToken(c.Request().Header)
		if err != nil {
			logger.PrintError(err, "error parse header token")
			return err
		}

		accessTokenClaims, err := jwt.ClaimsAccessToken(token)
		if err != nil {
			logger.PrintError(err, "error claims access token")
			return err
		}

		et.auth.SetClaims(&accessTokenClaims)

		err = et.auth.ValidateSession(c.Request().Context())
		if err != nil {
			return err
		}

		c.Set("auth", *et.auth)

		return next(c)
	}
}

func parseHeaderToken(h http.Header) (token string, err error) {
	headerDataToken := h.Get(constant.DefaultMdwHeaderToken)
	if !strings.Contains(headerDataToken, "Bearer") {
		err = constant.ErrHeaderTokenNotFound
		return
	}

	splitToken := strings.Split(headerDataToken, fmt.Sprintf("%s ", constant.DefaultMdwHeaderBearer))
	if len(splitToken) <= 1 {
		err = constant.ErrHeaderTokenInvalid
		return
	}

	token = splitToken[1]

	return
}
