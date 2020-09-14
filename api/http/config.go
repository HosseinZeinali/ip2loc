package http

import (
	"os"
	"strconv"
)

type Config struct {
	Port int
}

func InitConfig() (*Config, error) {
	apiPort, _ := strconv.Atoi(os.Getenv("API_PORT"))
	config := &Config{
		Port: apiPort,
	}

	return config, nil
}
