package config

import (
	"backend_bench/internal/model"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetConfig() model.Config {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file could not be loaded")
	}

	port := os.Getenv("PORT")
	stream := os.Getenv("STREAM_URL")
	return model.Config{
		Port:   port,
		Stream: stream,
	}
}
