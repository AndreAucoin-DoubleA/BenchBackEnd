package main

import (
	"backend_bench/internal/config"
	"backend_bench/internal/db"
	"backend_bench/internal/model"
	"backend_bench/internal/server"
	"backend_bench/internal/service/wikiconsumer"
	"fmt"
)

func main() {
	config := config.GetConfig()
	session := db.ConnectToCassandra(config.CassandraHost, config.CassandraPort, config.KeyspaceKey)
	defer session.Close()
	go wikiconsumer.StartWikiConsumer(config.Stream)
	fmt.Printf("Server is running on port: %s\n", config.Port)
	server.StartServer(config.Port, &model.UserRepository{Session: session}, config.JWTSecret)
}
