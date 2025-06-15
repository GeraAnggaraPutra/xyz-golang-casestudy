package service

import (
	"context"

	"kredit-plus/src/domain/user/payload"
	"kredit-plus/src/model"
	"kredit-plus/toolkit/logger"
)

func (s *Service) ReadUserListService(
	ctx context.Context,
	request payload.ReadUserListRequest,
) (data []model.User, totalData int64, err error) {
	statement := s.db.Model(&model.User{}).Where("deleted_at IS NULL")

	if request.SetSearch {
		statement = statement.Where("email ILIKE ?", request.Search)
	}

	statement = statement.Order(request.Order)

	if err = statement.Count(&totalData).Error; err != nil {
		logger.WithContext(ctx).Error(err, "error count user", "request", request)
		return
	}

	if request.SetPaginate {
		statement = statement.Limit(request.Limit).Offset(request.Offset)
	}

	if err = statement.Find(&data).Error; err != nil {
		logger.WithContext(ctx).Error(err, "error find user", "request", request)
		return
	}

	return
}
