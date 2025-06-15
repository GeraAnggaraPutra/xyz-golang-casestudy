package service

import (
	"context"
	"fmt"
	"strings"
	"time"

	"kredit-plus/src/domain/storage/payload"
	"kredit-plus/toolkit/logger"
)

func (s *Service) CreateFileUploadService(
	ctx context.Context,
	request payload.CreateFileUploadRequest,
) (path string, err error) {
	var (
		date     = time.Now().Format(time.DateOnly)
		now      = time.Now().Unix()
		filename = fmt.Sprintf("%d_%s", now, strings.ReplaceAll(request.File.Filename, " ", "_"))
	)

	path = fmt.Sprintf("%s/%s/%s", request.Path, date, filename)

	err = s.gcs.UploadFile(ctx, request.File, path)
	if err != nil {
		logger.WithContext(ctx).Error(err, "error upload file to cloud storage", "request", request)
		return
	}

	return
}
