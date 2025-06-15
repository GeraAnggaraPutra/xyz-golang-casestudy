package cache

import (
	"time"

	"kredit-plus/toolkit/constant"
)

// ConnectionOption is cache connection option.
type connectionOption struct {
	maxRetries        int           // 3
	minRetryBackoff   time.Duration // 8ms
	maxRetryBackoff   time.Duration // 512ms
	dialTimeout       time.Duration // 5s
	readTimeout       time.Duration // 3s
	writeTimeout      time.Duration // 3s
	poolFIFO          bool          // true
	poolSize          int           // 10
	poolTimeout       time.Duration // 4s
	maxIdleConns      int           // 0
	maxActiveConns    int           // 0
	connMaxIdleTime   time.Duration // 30m
	connMaxLifeTime   time.Duration // 0
	identitySuffix    string        // ""
	keepAliveInterval time.Duration
}

// DefaultConnectionOption returns sensible conn setting.
func defaultConnectionOption() *connectionOption {
	return &connectionOption{
		maxRetries:        constant.DefaultCacheMaxRetries,
		minRetryBackoff:   constant.DefaultCacheMinRetryBackoff,
		maxRetryBackoff:   constant.DefaultCacheMaxRetryBackoff,
		dialTimeout:       constant.DefaultCacheDialTimeout,
		readTimeout:       constant.DefaultCacheReadTimeout,
		writeTimeout:      constant.DefaultCacheWriteTimeout,
		poolFIFO:          constant.DefaultCachePoolFIFO,
		poolSize:          constant.DefaultCachePoolSize,
		poolTimeout:       constant.DefaultCachePoolTimeout,
		maxIdleConns:      constant.DefaultCacheMaxIdleConns,
		maxActiveConns:    constant.DefaultCacheMaxActiveConns,
		connMaxIdleTime:   constant.DefaultCacheConnMaxIdleTime,
		connMaxLifeTime:   constant.DefaultCacheConnMaxLifeTime,
		identitySuffix:    constant.DefaultCacheIdentitySuffix,
		keepAliveInterval: constant.DefaultCacheKeepAliveInterval,
	}
}
