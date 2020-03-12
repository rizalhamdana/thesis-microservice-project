package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config configuration struct
type Config struct {
	DB *DBConfig
}

// DBConfig configuration for databae struct
type DBConfig struct {
	Host     string
	Dialect  string
	Username string
	Password string
	Name     string
	Charset  string
}

// GetConfig is used for getting some configuration for our system
func GetConfig() *Config {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf(err.Error())
	}
	return &Config{
		DB: &DBConfig{
			Host:     os.Getenv("DB_HOST"),
			Dialect:  os.Getenv("DB_DRIVER"),
			Username: os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Name:     os.Getenv("DB_NAME"),
			Charset:  "utf8",
		},
	}
}
