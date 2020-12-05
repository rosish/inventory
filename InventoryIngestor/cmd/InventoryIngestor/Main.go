package main

import (
	"fmt"
	"inventory/InventoryIngestor/pkg/bmcClient"
)
func main() {
	fmt.Printf("\nRunning ingestor application\n")
	fmt.Printf(bmcClient.FetchLiveInventory())
}


