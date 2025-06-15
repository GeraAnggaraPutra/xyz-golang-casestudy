package service

import (
	"context"

	"kredit-plus/src/domain/transaction/payload"
	"kredit-plus/src/model"
	"kredit-plus/toolkit/logger"
)

func (s *Service) ReadTransactionDetailService(
	ctx context.Context,
	request payload.ReadTransactionDetailRequest,
) (data model.Transaction, err error) {
	data = model.Transaction{GUID: request.GUID}

	if err = s.db.Preload("Customer").First(&data).Error; err != nil {
		logger.WithContext(ctx).Error(err, "error find transaction by GUID : ", request.GUID)
		return
	}

	return
}
