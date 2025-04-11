BROKER_BINARY=brokerApp
AUTH_BINARY=authApp



up:
	@echo "Starting Docker images..."
	docker compose up -d 
	@echo "Docker images started."

## down: stop docker compose
down:
	@echo "Stopping docker compose..."
	 docker compose down
	@echo "Done!"

up_build: build_broker
	@echo "Stopping docker images..."
	docker compose down 
	@echo "Building and starting docker images"
	docker compose up --build -d 
	@echo "Docker image built and started!"

build_broker:
	@echo "Building broker binary..."
	cd ../broker-service && env GOOS=linux CGO_ENABLED=0 go build -o $(BROKER_BINARY) ./cmd/api
	@echo "done"

build_auth:
	@echo "Building broker binary..."
	cd ../authentication-service && env GOOS=linux CGO_ENABLED=0 go build -o $(AUTH_BINARY) ./cmd/api
	@echo "done"