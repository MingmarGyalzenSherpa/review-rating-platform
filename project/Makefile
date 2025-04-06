BROKER_BINARY=brokerApp


up:
	@echo "Starting Docker images..."
	docker compose up -d 
	@echo "Docker images started."

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