package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4/middleware"
)

type Config struct {
	PORT          string
	LOGGER_CONFIG *middleware.LoggerConfig
}

func SetupConfig() *Config {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	loggerConfig := &middleware.LoggerConfig{
		Format:           "[${time_custom}]:  Method ${method}  Status ${status}  Path ${uri}\n",
		CustomTimeFormat: "1/2/2006 15:04:05",
	}

	return &Config{
		PORT:          port,
		LOGGER_CONFIG: loggerConfig,
	}
}
