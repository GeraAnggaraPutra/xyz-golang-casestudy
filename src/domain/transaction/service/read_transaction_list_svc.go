package service

import (
	"context"

	"kredit-plus/src/domain/transaction/payload"
	"kredit-plus/src/model"
	"kredit-plus/toolkit/logger"
)

func (s *Service) ReadTransactionListService(
	ctx context.Context,
	request payload.ReadTransactionListRequest,
) (data []model.Transaction, totalData int64, err error) {
	statement := s.db.Model(&model.Transaction{}).Where("deleted_at IS NULL")

	if request.SetSearch {
		statement = statement.Where(`
			contract_no ILIKE ? 
			OR CAST(otr AS VARCHAR) ILIKE ?
			OR CAST(admin_fee AS VARCHAR) ILIKE ?
			OR CAST(installment_amount AS VARCHAR) ILIKE ?
			OR CAST(interest_amount AS VARCHAR) ILIKE ?
			OR asset_name ILIKE ?
			OR asset_type ILIKE ?
			`, request.Search, request.Search, request.Search,
			request.Search, request.Search, request.Search, request.Search)
	}

	statement = statement.Order(request.Order)

	if err = statement.Count(&totalData).Error; err != nil {
		logger.WithContext(ctx).Error(err, "error count transaction", "request", request)
		return
	}

	if request.SetPaginate {
		statement = statement.Limit(request.Limit).Offset(request.Offset)
	}

	if err = statement.Find(&data).Error; err != nil {
		logger.WithContext(ctx).Error(err, "error find transaction", "request", request)
		return
	}

	return
}
