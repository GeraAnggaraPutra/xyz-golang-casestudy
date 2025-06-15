package service

import (
	"context"
	"database/sql"
	"time"

	"kredit-plus/src/constant"
	"kredit-plus/src/domain/customer/payload"
	"kredit-plus/src/model"
	"kredit-plus/src/util"
	"kredit-plus/toolkit/logger"

	"github.com/pkg/errors"
)

func (s *Service) UpdateCustomerService(
	ctx context.Context,
	request payload.UpdateCustomerRequest,
	userGUID string,
) (err error) {
	tx := s.db.Begin()
	if tx.Error != nil {
		logger.WithContext(ctx).Error(tx.Error, "failed to begin transaction")
		err = errors.WithStack(constant.ErrUnknownSource)
		return
	}

	defer func() {
		if err != nil {
			if errRollback := tx.Rollback(); errRollback != nil {
				logger.WithContext(ctx).Error(err, "error rollback", errRollback)
				err = errors.WithStack(constant.ErrUnknownSource)
				return
			}
		}
	}()

	date, err := time.Parse(time.RFC3339, request.BirthDate)
	if err != nil {
		return
	}

	customer := model.Customer{
		GUID:        request.GUID,
		NIK:         request.NIK,
		FullName:    request.FullName,
		LegalName:   request.LegalName,
		BirthPlace:  request.BirthPlace,
		BirthDate:   date,
		Salary:      request.Salary,
		PhotoKTP:    util.ExtractFileURL(request.PhotoKTP),
		PhotoSelfie: util.ExtractFileURL(request.PhotoSelfie),
		UpdatedBy:   sql.NullString{String: userGUID, Valid: true},
	}

	if err = tx.Updates(&customer).Error; err != nil {
		logger.WithContext(ctx).Error(err, "error update customer", "customer model", customer)
		return
	}

	if err := tx.Where("customer_guid = ?", request.GUID).
		Delete(&model.CustomerTenorLimit{}).Error; err != nil {
		logger.WithContext(ctx).Error(err, "error deleting old tenor", "GUID", request.GUID)
		return err
	}

	var tenors []model.CustomerTenorLimit
	for _, tenor := range request.Tenors {
		tenors = append(tenors, model.CustomerTenorLimit{
			CustomerGUID: request.GUID,
			TenorMonths:  tenor.TenorMonths,
			LimitAmount:  tenor.LimitAmount,
			CreatedBy:    userGUID,
		})
	}

	if len(tenors) > 0 {
		if err := tx.Create(&tenors).Error; err != nil {
			logger.WithContext(ctx).Error(err, "error creating tenors", "tenors", tenors)
			return err
		}
	}

	if err = tx.Commit().Error; err != nil {
		logger.WithContext(ctx).Error(err, "error commit")
		err = errors.WithStack(constant.ErrUnknownSource)
	}

	return
}

func (s *Service) IsNIKExistsExcludeCustomerService(ctx context.Context, nik, customerGUID string) (bool, error) {
	var count int64
	err := s.db.Model(&model.Customer{}).
		Where("nik = ? AND guid != ? AND deleted_at IS NULL", nik, customerGUID).
		Count(&count).Error

	if err != nil {
		logger.WithContext(ctx).Error(err, "error checking nik existence", "nik", nik, "customerGUID", customerGUID)
		return false, err
	}

	return count > 0, nil
}
