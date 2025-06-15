package kernel

import (
	"kredit-plus/src/handler/storage"

	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Kernel struct {
	cache redis.UniversalClient
	db    *gorm.DB
	dbx   *sqlx.DB
	gcs   *storage.Storage
}

func NewKernel(
	cache redis.UniversalClient,
	db *gorm.DB,
	dbx *sqlx.DB,
	gcs *storage.Storage,
) *Kernel {
	return &Kernel{
		cache: cache,
		db:    db,
		dbx:   dbx,
		gcs:   gcs,
	}
}

func (k *Kernel) GetCache() redis.UniversalClient {
	return k.cache
}

func (k *Kernel) GetDB() *gorm.DB {
	return k.db
}

func (k *Kernel) GetDBX() *sqlx.DB {
	return k.dbx
}

func (k *Kernel) GetGCS() *storage.Storage {
	return k.gcs
}
