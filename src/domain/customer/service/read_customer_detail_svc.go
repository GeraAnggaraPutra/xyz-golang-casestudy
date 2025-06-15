package service

import (
	"context"

	"kredit-plus/src/domain/customer/payload"
	"kredit-plus/src/model"
	"kredit-plus/toolkit/logger"
)

func (s *Service) ReadCustomerDetailService(
	ctx context.Context,
	request payload.ReadCustomerDetailRequest,
) (data model.Customer, err error) {
	data = model.Customer{GUID: request.GUID}

	if err = s.db.Preload("TenorLimits").First(&data).Error; err != nil {
		logger.WithContext(ctx).Error(err, "error find customer by GUID : ", request.GUID)
		return
	}

	return
}
