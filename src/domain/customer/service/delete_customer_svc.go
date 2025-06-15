package service

import (
	"context"
	"time"

	"kredit-plus/src/domain/customer/payload"
	"kredit-plus/src/model"
	"kredit-plus/toolkit/logger"
)

func (s *Service) DeleteCustomerService(
	ctx context.Context,
	request payload.DeleteCustomerRequest,
	userGUID string,
) (err error) {
	updates := map[string]interface{}{
		"deleted_at": time.Now(),
		"deleted_by": userGUID,
	}

	if err = s.db.Model(&model.Customer{}).
		Where("guid = ?", request.GUID).
		Updates(updates).Error; err != nil {
		logger.WithContext(ctx).Error(err, "error soft delete customer by GUID : ", request.GUID)
		return
	}

	return
}
