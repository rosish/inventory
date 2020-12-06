package manager

import (
	"fmt"
	"github.com/rosish/inventory/InventoryDb/pkg/bmcDbClient"
	"github.com/rosish/inventory/InventoryIngestor/pkg/bmcClient"
)

func Discover() {

	fmt.Println("inventory discovered : !")

	fmt.Println(bmcClient.FetchLiveInventory())
	fmt.Println(bmcDbClient.FetchCachedInventory())
}