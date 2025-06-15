package main

import (
	"log"
	"os"
	"time"

	"kredit-plus/src/api"
	"kredit-plus/src/handler/storage"
	"kredit-plus/src/kernel"
	"kredit-plus/toolkit/cache"
	"kredit-plus/toolkit/config"
	"kredit-plus/toolkit/db"
	"kredit-plus/toolkit/logger"

	"github.com/joho/godotenv"
)

func main() {
	var err error

	// load .env file
	if os.Getenv("APP_ENV") == "" {
		err = godotenv.Load(".env")
		if err != nil {
			log.Fatalf("ERROR load env file : %s", err.Error())
		}
	}

	ctx, cancel := config.NewRuntimeContext()
	defer func() {
		cancel()

		if err != nil {
			log.Printf("found error : %s", err.Error())
		}
	}()

	setDefaultTimezone()

	// setup logger
	logger.NewLogger()

	// setup cache
	cache, err := cache.NewCache()
	if err != nil {
		log.Printf("ERROR setup cache : %s", err.Error())
		return
	}

	// setup database
	dbx, database, err := db.NewDatabase()
	if err != nil {
		log.Printf("ERROR setup database : %s", err.Error())
		return
	}

	// setup gcs
	gcs, err := storage.NewStorage(ctx)
	if err != nil {
		log.Printf("ERROR setup GCS : %s", err.Error())
		return
	}

	// setup module
	k := kernel.NewKernel(cache, database, dbx, &gcs)

	// run echo http
	api.RunEchoServer(ctx, k)
}

func setDefaultTimezone() {
	loc, err := time.LoadLocation("UTC")
	if err != nil {
		loc = time.Now().Location()
	}

	time.Local = loc
}
