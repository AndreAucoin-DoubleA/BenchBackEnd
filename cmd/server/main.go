package main

import (
	"backend_bench/internal/config"
	"backend_bench/internal/server"
	wikiconsumer "backend_bench/internal/service/wikiconsumer"
	"fmt"
)

func main() {
	config := config.GetConfig()
	go wikiconsumer.StartWikiConsumer(config.Stream)
	fmt.Printf("Server is running on port: %s\n", config.Port)
	server.StartServer(config.Port)
}
