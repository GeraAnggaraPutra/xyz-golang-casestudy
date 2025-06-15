package query

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Query struct {
	cache redis.UniversalClient
	db    *gorm.DB
}

func NewQuery(cache redis.UniversalClient, db *gorm.DB) *Query {
	return &Query{
		cache: cache,
		db:    db,
	}
}
