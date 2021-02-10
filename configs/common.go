package configs

import (
	"os"

	"github.com/sharpvik/log-go"
)

// Config contains configuration information for the whole app.
type Config struct {
	Database Database
}

// MustInit attempts to initialise Config and panics in case of failure.
func MustInit() (config Config) {
	log.Debug("config common")
	config.Database = mustInitDB()
	return
}

func mustGet(key string) (r string) {
	r = os.Getenv(key)
	if r == "" {
		log.Fatal("failed to get '%s'", key)
	}
	return
}
