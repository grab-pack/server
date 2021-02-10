package configs

import "github.com/sharpvik/log-go"

type Database struct {
	Host       string
	Port       string
	User       string
	Password   string
	Name       string
	Migrations string
}

func mustInitDB() (db Database) {
	log.Debug("config database")
	db.Host = mustGet("POSTGRES_HOST")
	db.Port = mustGet("POSTGRES_PORT")
	db.User = mustGet("POSTGRES_USER")
	db.Password = mustGet("POSTGRES_PASSWORD")
	db.Name = mustGet("POSTGRES_DB")
	db.Migrations = mustGet("POSTGRES_MIGRATIONS")
	return
}
