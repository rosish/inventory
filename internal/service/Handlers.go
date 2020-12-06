package service

import (
	"github.com/rosish/goModuleTest/pkg/logger"
	"github.com/rosish/inventory/InventoryDb/pkg/bmcDbClient"
	"github.com/rosish/inventory/InventoryIngestor/pkg/bmcClient"
	"net/http"

)

func InventoryAccessHandler(w http.ResponseWriter, r *http.Request) {
	logger.Log()

	data := bmcClient.FetchLiveInventory()
	w.Write([]byte(data))

	data1 := bmcDbClient.FetchCachedInventory()
	w.Write([]byte(data1))
}
