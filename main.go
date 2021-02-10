package main

import (
	"github.com/joho/godotenv"
	"github.com/sharpvik/log-go"

	"github.com/grab-pack/server/configs"
	"github.com/grab-pack/server/database"
)

func mustInit() (config configs.Config, db *database.Database) {
	log.Debug("config godotenv")
	if err := godotenv.Load(); err != nil {
		log.Fatal("failed to load env")
	}
	config = configs.MustInit()
	db = database.MustInit(config.Database)
	return
}

func main() {
	_, _ = mustInit()
	log.Debug("init successfull")
}
