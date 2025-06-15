package cache

import (
	"os"
	"strconv"

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
