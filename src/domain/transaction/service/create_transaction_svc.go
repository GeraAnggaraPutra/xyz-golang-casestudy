package service

import (
	"context"
	"fmt"
	"time"

	"kredit-plus/src/constant"
	"kredit-plus/src/domain/transaction/payload"
	"kredit-plus/src/model"
	"kredit-plus/src/util"
	"kredit-plus/toolkit/logger"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm/clause"
)

const (
	adminFeeRate         = constant.DefaultAdminFee
	interestRatePerMonth = constant.DefaultInterestRatePerMonth
)

func (s *Service) CreateTransactionService(
	ctx context.Context,
	request payload.CreateTransactionRequest,
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

	var customerLimit model.CustomerTenorLimit
	err = tx.Clauses(clause.Locking{Strength: "UPDATE"}).
		Where("customer_guid = ? AND tenor_months = ?",
			request.CustomerGUID, request.TenorMonths).First(&customerLimit).Error
	if err != nil {
		logger.WithContext(ctx).Error(err, "error find customer tenor limit",
			"customer_guid", request.CustomerGUID, "tenor_months", request.TenorMonths)
		return
	}

	transaction := model.Transaction{
		GUID:         util.GenerateUUID(),
		CustomerGUID: request.CustomerGUID,
		AssetName:    request.AssetName,
		AssetType:    request.AssetType,
		OTR:          request.OTR,
		TenorMonths:  request.TenorMonths,
		CreatedBy:    userGUID,
	}

	transaction.AdminFee = request.OTR * adminFeeRate

	principalForLoan := request.OTR + transaction.AdminFee

	if principalForLoan > customerLimit.LimitAmount {
		err = constant.ErrLimitExceeded
		logger.WithContext(ctx).Warn("Transaction OTR + AdminFee exceeds customer limit",
			"customer_guid", request.CustomerGUID, "tenor_months", request.TenorMonths,
			"otr_plus_admin_fee", principalForLoan, "limit_amount", customerLimit.LimitAmount)
		return
	}

	transaction.InterestAmount = principalForLoan * interestRatePerMonth * float64(request.TenorMonths)
	totalAmountToRepay := principalForLoan + transaction.InterestAmount
	transaction.InstallmentAmount = totalAmountToRepay / float64(request.TenorMonths)
	transaction.ContractNo = fmt.Sprintf("KP-%s-%s-%s",
		request.AssetType, time.Now().Format("20060102"), uuid.New().String()[:6])

	if err = tx.Omit("updated_at").Create(&transaction).Error; err != nil {
		logger.WithContext(ctx).Error(err, "error create transaction", "transaction", transaction)
		return
	}

	if err = tx.Commit().Error; err != nil {
		logger.WithContext(ctx).Error(err, "error commit")
		err = errors.WithStack(constant.ErrUnknownSource)
	}

	return
}
