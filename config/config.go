package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

type BrokerConfig struct {
	Driver   string
	Host     string
	Port     int
	Username string
	Password string
}

type Config struct {
	Broker BrokerConfig
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func New() *Config {
	return &Config{
		Broker: BrokerConfig{
			Driver:   getStringEnv("BROKER_DRIVER", "rabbitmq"),
			Host:     getStringEnv("BROKER_HOST", "localhost"),
			Port:     getIntEnv("BROKER_PORT", 5672),
			Username: getStringEnv("BROKER_USERNAME", "guest"),
			Password: getStringEnv("BROKER_PASSWORD", "guest"),
		},
	}
}

// getStringEnv read an environment as string or return a default value
func getStringEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}

// getIntEnv read an environment as int or return a default value
func getIntEnv(key string, defaultValue int) int {
	str := getStringEnv(key, "")

	if value, err := strconv.Atoi(str); err == nil {
		return value
	}

	return defaultValue
}
