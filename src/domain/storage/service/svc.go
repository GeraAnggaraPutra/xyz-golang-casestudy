package service

import (
	"kredit-plus/src/handler/storage"

	"gorm.io/gorm"
)

type Service struct {
	db  *gorm.DB
	gcs *storage.Storage
}

func NewService(db *gorm.DB, gcs *storage.Storage) *Service {
	return &Service{
		db:  db,
		gcs: gcs,
	}
}
