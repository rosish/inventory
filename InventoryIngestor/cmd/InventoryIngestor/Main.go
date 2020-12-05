package main

import (
	"fmt"
	"github.com/rosish/inventory/InventoryIngestor/pkg/bmcClient"
)
func main() {
	fmt.Printf("\nRunning ingestor application\n")
	fmt.Printf(bmcClient.FetchInventory())
}


