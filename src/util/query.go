package util

import (
	"context"
	"database/sql"

	"kredit-plus/toolkit/logger"

	"gorm.io/gorm"
)

// Wrapper for query transaction.
func Transaction(
	ctx context.Context,
	db *gorm.DB,
	txFunc func(db *gorm.DB) (err error),
) (err error) {
	db = db.WithContext(ctx)

	tx := db.Begin(&sql.TxOptions{})

	err = txFunc(db)
	if err != nil {
		if errRollback := tx.Rollback().Error; errRollback != nil {
			logger.WithContext(ctx).Error(errRollback, "error rollback")
			return
		}

		return
	}

	if err = tx.Commit().Error; err != nil {
		logger.WithContext(ctx).Error(err, "error commit")
		return
	}

	return
}
