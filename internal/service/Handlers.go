package service

import (
	"github.com/rosish/goModuleTest/pkg/logger"
	"inventory/InventoryIngestor/pkg/bmcClient"
	"inventory/db/cmd/pkg/manager"
	"net/http"

)

func InventoryAccessHandler(w http.ResponseWriter, r *http.Request) {
	logger.Log()

	data := bmcClient.FetchLiveInventory()
	w.Write([]byte(data))

	data1 := manager.FetchCachedInventory()
	w.Write([]byte(data1))
}
