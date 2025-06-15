package cache

import (
	"context"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

func keepAlive(ctx context.Context, redis *redis.Client, interval time.Duration) {
	for {
		cmd := redis.Ping(ctx)
		if cmd.Err() != nil {
			log.Printf("ERROR redis.keepAlive \n%s \n\n", cmd.Err())
		}

		time.Sleep(interval)
	}
}
