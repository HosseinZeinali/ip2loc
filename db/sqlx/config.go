package sqlx

import (
	"fmt"
	"os"
)

type Config struct {
	DatabaseURI string
	Driver      string
}

func InitConfig() (*Config, error) {
	var dbHost string = "localhost"
	var dbName string = os.Getenv("POSTGRES_DB")
	var dbUser string = os.Getenv("POSTGRES_USER")
	var dbPassword string = os.Getenv("POSTGRES_PASSWORD")

	connection := "host=" + dbHost + " port=2345 user=" + dbUser + " dbname=" + dbName + " password=" + dbPassword + " sslmode=disable"

	config := &Config{
		DatabaseURI: connection,
	}
	if config.DatabaseURI == "" {
		return nil, fmt.Errorf("DatabaseURI must be set")
	}
	return config, nil
}
