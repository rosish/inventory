module github.com/rosish/inventory

go 1.14

replace (
	inventory/InventoryIngestor v0.0.0 => ./InventoryIngestor
	inventory/db v0.0.0 => ./InventoryDb
)

require (
	github.com/gorilla/mux v1.8.0
	github.com/rosish/goModuleTest v0.0.0-20201204165957-0f6a9cf7476f
	inventory/InventoryIngestor v0.0.0
	inventory/db v0.0.0
)


