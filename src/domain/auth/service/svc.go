package service

import (
	"gorm.io/gorm"
	"github.com/redis/go-redis/v9"
)

type Service struct {
	db *gorm.DB
	cache redis.UniversalClient
}

func NewService(db *gorm.DB, cache redis.UniversalClient) *Service {
	return &Service{
		db: db,
		cache: cache,
	}
}
