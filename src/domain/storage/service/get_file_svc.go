package service

import (
	"context"

	"kredit-plus/src/domain/storage/payload"
	"kredit-plus/toolkit/logger"

	"cloud.google.com/go/storage"
)

func (s *Service) GetFileService(
	ctx context.Context,
	request payload.GetFileRequest,
) (sr *storage.Reader, err error) {
	sr, err = s.gcs.GetFile(ctx, request.Path)
	if err != nil {
		logger.WithContext(ctx).Error(err, "error get file", "path", request.Path)
		return
	}

	return
}
