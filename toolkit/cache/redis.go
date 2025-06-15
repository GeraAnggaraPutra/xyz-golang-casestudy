package cache

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func NewRedis(opt *cacheOption) (client *redis.Client, err error) {
	client = redis.NewClient(&redis.Options{
		Addr:            fmt.Sprintf("%s:%d", opt.host, opt.port),
		Username:        opt.username,
		Password:        opt.password,
		DB:              opt.db,
		MaxRetries:      opt.maxRetries,
		MinRetryBackoff: opt.minRetryBackoff,
		MaxRetryBackoff: opt.maxRetryBackoff,
		DialTimeout:     opt.dialTimeout,
		ReadTimeout:     opt.readTimeout,
		WriteTimeout:    opt.writeTimeout,
		PoolFIFO:        opt.poolFIFO,
		PoolSize:        opt.poolSize,
		PoolTimeout:     opt.poolTimeout,
		MaxIdleConns:    opt.maxIdleConns,
		MaxActiveConns:  opt.maxActiveConns,
		ConnMaxIdleTime: opt.connMaxIdleTime,
		ConnMaxLifetime: opt.connMaxLifeTime,
		IdentitySuffix:  opt.identitySuffix,
	})

	go keepAlive(context.Background(), client, opt.keepAliveInterval)

	return
}
