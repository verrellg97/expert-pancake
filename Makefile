ACCOUNT_BINARY=accountService
BUSINESS_BINARY=businessService
ACCOUNTING_BINARY=accountingService
BUSINESS_RELATION_BINARY=businessRelationService
INVENTORY_BINARY=inventoryService
WAREHOUSE_BINARY=warehouseService
NOTIFICATION_BINARY=notificationService

## up: starts all containers in the background without forcing build
up:
	@echo "Starting Docker images..."
	docker compose up -d
	@echo "Docker images started!"

## up_build: stops docker-compose (if running), builds all projects and starts docker compose
up_build: build_account build_business build_accounting build_business_relation build_inventory build_warehouse
	@echo "Stopping docker images (if running...)"
	docker compose down
	@echo "Building (when required) and starting docker images..."
	docker compose up --build -d
	@echo "Docker images built and started!"

## down: stop docker compose
down:
	@echo "Stopping docker compose..."
	docker compose down
	@echo "Done!"

## build_account: builds the account binary as a linux executable
build_account:
	@echo "Building account binary..."
	cd ./service/account && env GOOS=linux CGO_ENABLED=0 go build -o ${ACCOUNT_BINARY} ./cmd/api
	@echo "Done!"

## build_business: builds the business binary as a linux executable
build_business:
	@echo "Building business binary..."
	cd ./service/business && env GOOS=linux CGO_ENABLED=0 go build -o ${BUSINESS_BINARY} ./cmd/api
	@echo "Done!"

## build_accounting: builds the accounting binary as a linux executable
build_accounting:
	@echo "Building accounting binary..."
	cd ./service/accounting && env GOOS=linux CGO_ENABLED=0 go build -o ${ACCOUNTING_BINARY} ./cmd/api
	@echo "Done!"

## build_business_relation: builds the business_relation binary as a linux executable
build_business_relation:
	@echo "Building business_relation binary..."
	cd ./service/business-relation && env GOOS=linux CGO_ENABLED=0 go build -o ${BUSINESS_RELATION_BINARY} ./cmd/api
	@echo "Done!"

## build_inventory: builds the inventory binary as a linux executable
build_inventory:
	@echo "Building inventory binary..."
	cd ./service/inventory && env GOOS=linux CGO_ENABLED=0 go build -o ${INVENTORY_BINARY} ./cmd/api
	@echo "Done!"

## build_warehouse: builds the warehouse binary as a linux executable
build_warehouse:
	@echo "Building warehouse binary..."
	cd ./service/warehouse && env GOOS=linux CGO_ENABLED=0 go build -o ${WAREHOUSE_BINARY} ./cmd/api
	@echo "Done!"

## build_notification: builds the notification binary as a linux executable
build_notification:
	@echo "Building notification binary..."
	cd ./service/notification && env GOOS=linux CGO_ENABLED=0 go build -o ${NOTIFICATION_BINARY} ./cmd/api
	@echo "Done!"

## up_appwrite: starts appwrite in the background
up_appwrite:
	@echo "Starting Appwrite..."
	cd ./appwrite && docker compose up -d
	@echo "Appwrite started!"

## down_appwrite: stops appwrite
down_appwrite:
	@echo "Stopping Appwrite..."
	cd ./appwrite && docker compose down
	@echo "Appwrite stopped!"