package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func Create() Config {
	if err := godotenv.Load("../../.env"); err != nil {
		log.Fatal("Error: .env file not found")
	}
	return Config{
		Port: os.Getenv("PORT"),
	}
}

func DBConfigCreate() DBConfig {
	if err := godotenv.Load("../../.env"); err != nil {
		log.Fatal("Error: .env file not found")
	}
	return DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DBName"),
		SSLMode:  os.Getenv("SSLMode"),
	}
}
