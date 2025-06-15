package service

import (
	"context"
	"database/sql"

	"kredit-plus/src/domain/user/payload"
	"kredit-plus/src/model"
	"kredit-plus/src/util"
	"kredit-plus/toolkit/logger"
)

func (s *Service) CreateUserService(
	ctx context.Context,
	request payload.CreateUserRequest,
	userGUID string,
) (err error) {
	password, err := util.GenerateHashPassword(request.Password)
	if err != nil {
		logger.WithContext(ctx).Error(err, "error generate hash password : "+request.Password)
		return
	}

	user := model.User{
		GUID:      util.GenerateUUID(),
		Email:     request.Email,
		Password:  password,
		CreatedBy: sql.NullString{String: userGUID, Valid: true},
	}

	if err = s.db.Omit("updated_at").Create(&user).Error; err != nil {
		logger.WithContext(ctx).Error(err, "error create user", "user", user)
		return
	}

	return
}

func (s *Service) IsEmailExistsService(ctx context.Context, email string) (bool, error) {
	var count int64
	err := s.db.Model(&model.User{}).
		Where("email = ? AND deleted_at IS NULL", email).
		Count(&count).Error

	if err != nil {
		logger.WithContext(ctx).Error(err, "error checking email existence", "email", email)
		return false, err
	}

	return count > 0, nil
}
