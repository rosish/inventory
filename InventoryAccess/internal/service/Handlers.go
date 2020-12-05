package service

import (
	"github.com/rosish/goModuleTest/pkg/logger"
	"github.com/rosish/inventory/InventoryIngestor/pkg/bmcClient"
	"net/http"
)

func InventoryAccessHandler(w http.ResponseWriter, r *http.Request) {
	inventory := bmcClient.FetchInventory()
	w.Write([]byte(inventory))
	logger.Log()
	//w.Write([]byte("hello 1"))
}
