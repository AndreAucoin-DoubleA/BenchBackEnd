package main

import (
	"backend_bench/internal/config"
	"backend_bench/internal/server"
	"fmt"
)

func main() {
	config := config.GetConfig()
	server.StartServer(config.Port)
	fmt.Printf("Server is running on port: %s\n", config.Port)
}
