package service

import (
	"context"

	"kredit-plus/src/domain/customer/payload"
	"kredit-plus/src/model"
	"kredit-plus/toolkit/logger"
)

func (s *Service) ReadCustomerListService(
	ctx context.Context,
	request payload.ReadCustomerListRequest,
) (data []model.Customer, totalData int64, err error) {
	statement := s.db.Model(&model.Customer{}).Where("deleted_at IS NULL")

	if request.SetSearch {
		statement = statement.Where(`
			nik ILIKE ? 
			OR full_name ILIKE ?
			OR legal_name ILIKE ?
			OR birth_place ILIKE ?
			`, request.Search, request.Search, request.Search, request.Search)
	}

	statement = statement.Order(request.Order)

	if err = statement.Count(&totalData).Error; err != nil {
		logger.WithContext(ctx).Error(err, "error count customer", "request", request)
		return
	}

	if request.SetPaginate {
		statement = statement.Limit(request.Limit).Offset(request.Offset)
	}

	if err = statement.Find(&data).Error; err != nil {
		logger.WithContext(ctx).Error(err, "error find customer", "request", request)
		return
	}

	return
}
