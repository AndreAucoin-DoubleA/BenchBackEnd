package config

import (
	"backend_bench/internal/model"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func GetConfig() model.Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file could not be loaded in local environment")
	}

	port := os.Getenv("PORT")
	stream := os.Getenv("STREAM_URL")
	cassandraHost := os.Getenv("CASSANDRA_HOST")
	cassandraPortStr := os.Getenv("CASSANDRA_PORT")
	cassandraPort, err := strconv.Atoi(cassandraPortStr)
	keyspaceKey := os.Getenv("KEYSPACE")
	if err != nil {
		log.Fatalf("Invalid CASSANDRA_PORT: %v", err)
	}

	return model.Config{
		Port:          port,
		Stream:        stream,
		CassandraPort: cassandraPort,
		CassandraHost: cassandraHost,
		KeyspaceKey:   keyspaceKey,
	}
}
