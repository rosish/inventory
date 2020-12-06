module github.com/rosish/inventory

go 1.14


replace (
	github.com/rosish/inventory/InventoryIngestor v0.0.0 => ./InventoryIngestor
	github.com/rosish/inventory/InventoryDb v0.0.0 => ./InventoryDb
)

require (
	github.com/gorilla/mux v1.8.0
	github.com/rosish/goModuleTest v0.0.0-20201204165957-0f6a9cf7476f
	github.com/rosish/inventory/InventoryIngestor v0.0.0
	github.com/rosish/inventory/InventoryDb v0.0.0
)


