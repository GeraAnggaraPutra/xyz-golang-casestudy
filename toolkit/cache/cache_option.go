package cache

import (
	"os"
	"strconv"

	"kredit-plus/toolkit/util"

	"github.com/pkg/errors"
)

type cacheOption struct {
	driver   string
	host     string
	port     int
	username string
	password string
	db       int
	*connectionOption
}

func newCacheOption() (opt *cacheOption, err error) {
	driver := os.Getenv("CACHE_DRIVER")
	host := os.Getenv("CACHE_HOST")
	portStr := os.Getenv("CACHE_PORT")
	username := os.Getenv("CACHE_USERNAME")
	password := os.Getenv("CACHE_PASSWORD")
	dbStr := os.Getenv("CACHE_DB")

	if host == "" || portStr == "" || dbStr == "" {
		return nil, errors.Wrapf(errors.New("invalid data source host or port"), "db: host=%s port=%s db=%s", host, portStr, dbStr)
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		return nil, errors.Wrapf(err, "error parse int on port db env : %s", portStr)
	}

	db, err := strconv.Atoi(dbStr)
	if err != nil {
		return nil, errors.Wrapf(err, "error parse int on db env : %s", dbStr)
	}

	connOpt := defaultConnectionOption()
	connOpt.maxRetries = util.ParseInt(connOpt.maxRetries, os.Getenv("CACHE_MAX_RETRIES"))
	connOpt.minRetryBackoff = util.ParseDuration(connOpt.minRetryBackoff, os.Getenv("CACHE_MIN_RETRY_BACKOFF"))
	connOpt.maxRetryBackoff = util.ParseDuration(connOpt.maxRetryBackoff, os.Getenv("CACHE_MAX_RETRY_BACKOFF"))
	connOpt.dialTimeout = util.ParseDuration(connOpt.dialTimeout, os.Getenv("CACHE_DIAL_TIMEOUT"))
	connOpt.readTimeout = util.ParseDuration(connOpt.readTimeout, os.Getenv("CACHE_READ_TIMEOUT"))
	connOpt.writeTimeout = util.ParseDuration(connOpt.writeTimeout, os.Getenv("CACHE_WRITE_TIMEOUT"))
	connOpt.poolFIFO = util.ParseBool(connOpt.poolFIFO, os.Getenv("CACHE_POOL_FIFO"))
	connOpt.poolSize = util.ParseInt(connOpt.poolSize, os.Getenv("CACHE_POOL_SIZE"))
	connOpt.poolTimeout = util.ParseDuration(connOpt.poolTimeout, os.Getenv("CACHE_POOL_TIMEOUT"))
	connOpt.maxIdleConns = util.ParseInt(connOpt.maxIdleConns, os.Getenv("CACHE_MAX_IDLE_CONNS"))
	connOpt.maxActiveConns = util.ParseInt(connOpt.maxActiveConns, os.Getenv("CACHE_MAX_ACTIVE_CONNS"))
	connOpt.connMaxIdleTime = util.ParseDuration(connOpt.connMaxIdleTime, os.Getenv("CACHE_CONN_MAX_IDLE_TIME"))
	connOpt.connMaxLifeTime = util.ParseDuration(connOpt.connMaxLifeTime, os.Getenv("CACHE_CONN_MAX_LIFE_TIME"))
	connOpt.identitySuffix = os.Getenv("CACHE_IDENTITY_SUFFIX")
	connOpt.keepAliveInterval = util.ParseDuration(connOpt.keepAliveInterval, os.Getenv("CACHE_KEEP_ALIVE_INTERVAL"))

	return &cacheOption{
		driver:           driver,
		host:             host,
		port:             port,
		username:         username,
		password:         password,
		db:               db,
		connectionOption: connOpt,
	}, nil
}
