package helper

import (
	"context"

	"kredit-plus/src/domain/auth/payload"
	"kredit-plus/src/handler/jwt"
	"kredit-plus/src/model"
	"kredit-plus/toolkit/logger"
)

func GenerateSessionModel(
	ctx context.Context,
	request payload.SessionPayload,
) (data model.Session, err error) {
	accessToken, err := jwt.GenerateAccessToken(request.ToAccessTokenRequest())
	if err != nil {
		logger.WithContext(ctx).Error(err, "error generate access token")
		return
	}

	refreshToken, err := jwt.GenerateRefreshToken(request.ToRefreshTokenRequest())
	if err != nil {
		logger.WithContext(ctx).Error(err, "error generate refresh token")
		return
	}

	data = model.Session{
		GUID:                  request.SessionGUID,
		UserGUID:              request.UserGUID,
		AccessToken:           accessToken.Token,
		AccessTokenExpiredAt:  accessToken.ExpiresAt,
		RefreshToken:          refreshToken.Token,
		RefreshTokenExpiredAt: refreshToken.ExpiresAt,
		IPAddress:             request.IPAddress,
		UserAgent:             request.UserAgent,
	}

	return
}
