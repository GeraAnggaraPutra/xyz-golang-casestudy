package auth

import (
	"kredit-plus/src/constant"
	"kredit-plus/src/handler/jwt"

	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Auth struct {
	cache  redis.UniversalClient
	db     *gorm.DB
	claims *jwt.AccessTokenPayload
}

func NewAuth(cache redis.UniversalClient, db *gorm.DB) *Auth {
	return &Auth{
		cache: cache,
		db:    db,
	}
}

func GetAuth(c echo.Context) (*Auth, error) {
	a, ok := c.Get("auth").(Auth)
	if !ok {
		return nil, constant.ErrTokenUnauthorized
	}

	return &a, nil
}

func (a *Auth) GetClaims() *jwt.AccessTokenPayload {
	return a.claims
}

func (a *Auth) SetClaims(claims *jwt.AccessTokenPayload) {
	a.claims = claims
}
