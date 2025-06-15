package service

import (
	"context"

	"kredit-plus/src/domain/user/payload"
	"kredit-plus/src/model"
	"kredit-plus/toolkit/logger"
)

func (s *Service) ReadUserDetailService(
	ctx context.Context,
	request payload.ReadUserDetailRequest,
) (data model.User, err error) {
	data = model.User{GUID: request.GUID}

	if err = s.db.First(&data).Error; err != nil {
		logger.WithContext(ctx).Error(err, "error find user by GUID : ", request.GUID)
		return
	}

	return
}
