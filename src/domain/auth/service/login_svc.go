package service

import (
	"context"

	"kredit-plus/src/constant"
	"kredit-plus/src/domain/auth/helper"
	"kredit-plus/src/domain/auth/payload"
	"kredit-plus/src/model"
	"kredit-plus/src/query"
	"kredit-plus/src/util"
	"kredit-plus/toolkit/logger"

	"github.com/pkg/errors"
)

func (s *Service) LoginService(
	ctx context.Context,
	request payload.LoginRequest,
) (data model.Session, user model.User, err error) {
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

	if err = s.db.Where("email = ? AND deleted_at IS NULL", request.Email).First(&user).Error; err != nil {
		logger.WithContext(ctx).Error(err, "error find user by email : "+request.Email)
		return
	}

	if err = util.CompareHashPassword(request.Password, user.Password); err != nil {
		err = logger.PrintNewError(err, constant.ErrPasswordIncorrect)
		return
	}

	data, err = helper.GenerateSessionModel(ctx, request.ToSessionPayload(user.GUID))
	if err != nil {
		logger.PrintError(err, "error generate session model", "session payload", request.ToSessionPayload(user.GUID))
		return
	}

	q := query.NewQuery(s.cache, tx)

	err = q.CreateSessionQuery(ctx, data)
	if err != nil {
		logger.WithContext(ctx).Error(err, "error create session", "session model", data)
		return
	}

	if err = tx.Commit().Error; err != nil {
		logger.WithContext(ctx).Error(err, "error commit")
		err = errors.WithStack(constant.ErrUnknownSource)
	}

	return
}
