ACCOUNT_BINARY=accountService
BUSINESS_BINARY=businessService

## up: starts all containers in the background without forcing build
up:
	@echo "Starting Docker images..."
	docker compose up -d
	@echo "Docker images started!"

## up_build: stops docker-compose (if running), builds all projects and starts docker compose
up_build: build_account build_business
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
	@echo "Building acccount binary..."
	cd ./service/account && env GOOS=linux CGO_ENABLED=0 go build -o ${ACCOUNT_BINARY} ./cmd/api
	@echo "Done!"

## build_business: builds the business binary as a linux executable
build_business:
	@echo "Building business binary..."
	cd ./service/business && env GOOS=linux CGO_ENABLED=0 go build -o ${BUSINESS_BINARY} ./cmd/api
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