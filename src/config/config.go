package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
}

func Create() Config {
	if err := godotenv.Load("../../.env"); err != nil {
		log.Fatal("Error: .env file not found")
	}
	return Config{
		Port: os.Getenv("PORT"),
	}
}
