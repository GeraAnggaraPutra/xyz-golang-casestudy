package cache

import "github.com/redis/go-redis/v9"

func NewCache() (cache redis.UniversalClient, err error) {
	opt, err := newCacheOption()
	if err != nil {
		return
	}

	if opt.driver == "redis" {
		cache, err = NewRedis(opt)
	}

	return
}
