package service

import (
	"context"
	"time"

	"kredit-plus/src/constant"
	"kredit-plus/src/domain/customer/payload"
	"kredit-plus/src/model"
	"kredit-plus/src/util"
	"kredit-plus/toolkit/logger"

	"github.com/pkg/errors"
)

func (s *Service) CreateCustomerService(
	ctx context.Context,
	request payload.CreateCustomerRequest,
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
		GUID:        util.GenerateUUID(),
		NIK:         request.NIK,
		FullName:    request.FullName,
		LegalName:   request.LegalName,
		BirthPlace:  request.BirthPlace,
		BirthDate:   date,
		Salary:      request.Salary,
		PhotoKTP:    util.ExtractFileURL(request.PhotoKTP),
		PhotoSelfie: util.ExtractFileURL(request.PhotoSelfie),
		CreatedBy:   userGUID,
	}

	if err = tx.Omit("updated_at").Create(&customer).Error; err != nil {
		logger.WithContext(ctx).Error(err, "error create customer", "customer", customer)
		return
	}

	var tenors []model.CustomerTenorLimit
	for _, tenor := range request.Tenors {
		tenors = append(tenors, model.CustomerTenorLimit{
			CustomerGUID: customer.GUID,
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

func (s *Service) IsNIKExistsService(ctx context.Context, nik string) (bool, error) {
	var count int64
	err := s.db.Model(&model.Customer{}).
		Where("nik = ? AND deleted_at IS NULL", nik).
		Count(&count).Error

	if err != nil {
		logger.WithContext(ctx).Error(err, "error checking nik existence", "nik", nik)
		return false, err
	}

	return count > 0, nil
}
