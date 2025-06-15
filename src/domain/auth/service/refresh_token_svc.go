package service

import (
	"context"

	"kredit-plus/src/domain/auth/helper"
	"kredit-plus/src/domain/auth/payload"
	"kredit-plus/src/handler/jwt"
	"kredit-plus/src/model"
	"kredit-plus/toolkit/logger"
)

func (s *Service) RefreshTokenService(
	ctx context.Context,
	request payload.RefreshTokenRequest,
) (data model.Session, user model.User, err error) {
	refreshTokenClaims, err := jwt.ClaimsRefreshToken(request.RefreshToken)
	if err != nil {
		logger.PrintError(err, "error claims refresh token : "+request.RefreshToken)
		return
	}

	session := model.Session{GUID: refreshTokenClaims.GUID}

	if err = s.db.First(&session).Error; err != nil {
		logger.WithContext(ctx).Error(err, "error find session by GUID : ", refreshTokenClaims.GUID)
		return
	}

	data, err = helper.GenerateSessionModel(ctx, request.ToSessionPayload(session))
	if err != nil {
		logger.PrintError(err, "error generate session model", "session payload", request.ToSessionPayload(session))
		return
	}

	if err = s.db.Where("guid = ? AND deleted_at IS NULL", session.UserGUID).First(&user).Error; err != nil {
		logger.WithContext(ctx).Error(err, "error find user", "guid", session.UserGUID)
		return
	}

	if err = s.db.Updates(&data).Error; err != nil {
		logger.WithContext(ctx).Error(err, "error update session", "session model", data)
		return
	}

	return
}
