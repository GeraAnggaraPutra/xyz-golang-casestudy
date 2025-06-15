package service

import (
	"context"

	"kredit-plus/src/handler/jwt"
	"kredit-plus/src/model"
	"kredit-plus/toolkit/logger"
)

func (s *Service) LogoutService(
	ctx context.Context,
	claims *jwt.AccessTokenPayload,
) (err error) {
	if err = s.db.Delete(&model.Session{GUID: claims.GUID}).Error; err != nil {
		logger.WithContext(ctx).Error(err, "error delete session by GUID : ", claims.GUID)
		return
	}

	return
}
