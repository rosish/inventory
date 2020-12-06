package service

import (
	"github.com/rosish/goModuleTest/pkg/logger"
	"github.com/rosish/inventory/InventoryDb/pkg/bmcDbClient"
	"github.com/rosish/inventory/InventoryIngestor/pkg/bmcClient"
	"github.com/rosish/inventory/pkg/manager"
	"net/http"
)

func InventoryAccessHandler(w http.ResponseWriter, r *http.Request) {
	logger.Log()
	manager.Discover()
	data := bmcClient.FetchLiveInventory()
	w.Write([]byte(data))
	w.Write([]byte("\n"))
	data1 := bmcDbClient.FetchCachedInventory()
	w.Write([]byte(data1))
}
