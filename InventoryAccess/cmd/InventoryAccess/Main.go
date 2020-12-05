package main

import "github.com/rosish/inventory/InventoryAccess/internal/service"

func main() {
	server := service.NewServer()
	server.SetupRoutes()
	server.StartServer()
}

