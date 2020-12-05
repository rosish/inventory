package main

import "github.com/rosish/inventory/internal/service"

func main() {
	server := service.NewServer()
	server.SetupRoutes()
	server.StartServer()
}

