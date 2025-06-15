package service

import (
	"context"
	"time"

	"kredit-plus/src/domain/user/payload"
	"kredit-plus/src/model"
	"kredit-plus/toolkit/logger"
)

func (s *Service) DeleteUserService(
	ctx context.Context,
	request payload.DeleteUserRequest,
	userGUID string,
) (err error) {
	updates := map[string]interface{}{
		"deleted_at": time.Now(),
		"deleted_by": userGUID,
	}

	if err = s.db.Model(&model.User{}).
		Where("guid = ?", request.GUID).
		Updates(updates).Error; err != nil {
		logger.WithContext(ctx).Error(err, "error soft delete user by GUID : ", request.GUID)
		return
	}

	return
}
