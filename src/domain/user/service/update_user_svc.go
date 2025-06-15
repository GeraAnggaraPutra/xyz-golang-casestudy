package service

import (
	"context"
	"database/sql"

	"kredit-plus/src/domain/user/payload"
	"kredit-plus/src/model"
	"kredit-plus/src/util"
	"kredit-plus/toolkit/logger"
)

func (s *Service) UpdateUserService(
	ctx context.Context,
	request payload.UpdateUserRequest,
	userGUID string,
) (err error) {
	var password string

	if request.Password != "" {
		password, err = util.GenerateHashPassword(request.Password)
		if err != nil {
			logger.WithContext(ctx).Error(err, "error generate hash password : "+request.Password)
			return
		}
	}

	user := model.User{
		GUID:      request.GUID,
		Email:     request.Email,
		Password:  password,
		UpdatedBy: sql.NullString{String: userGUID, Valid: true},
	}

	if err = s.db.Updates(&user).Error; err != nil {
		logger.WithContext(ctx).Error(err, "error update user", "user model", user)
		return
	}

	return
}

func (s *Service) IsEmailExistsExcludeUserService(ctx context.Context, email, userGUID string) (bool, error) {
	var count int64
	err := s.db.Model(&model.User{}).
		Where("email = ? AND guid != ? AND deleted_at IS NULL", email, userGUID).
		Count(&count).Error
	
	if err != nil {
		logger.WithContext(ctx).Error(err, "error checking email existence", "email", email, "userGUID", userGUID)
		return false, err
	}
	
	return count > 0, nil
}
