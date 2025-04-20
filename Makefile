.PHONY: all build clean test run

BINARY_NAME=devops-mastery-server
DOCKER_IMAGE=devops-mastery
DOCKER_TAG=latest

all: build

build:
	@echo "Building..."
	go build -o bin/$(BINARY_NAME) src/cmd/main.go

clean:
	@echo "Cleaning..."
	rm -rf bin/

test:
	@echo "Running tests..."
	go test ./... -v

run:
	@echo "Running server..."
	go run src/cmd/main.go

deps:
	@echo "Installing dependencies..."
	go mod tidy
	go get github.com/go-chi/chi/v5
	go get github.com/go-chi/cors
	go get github.com/spf13/viper

fmt:
	@echo "Formatting code..."
	go fmt ./...

vet:
	@echo "Vetting code..."
	go vet ./...

docker-build:
	@echo "Building Docker image..."
	docker build -t $(DOCKER_IMAGE):$(DOCKER_TAG) .

docker-run: docker-build
	@echo "Running Docker container..."
	docker run -p 8081:8081 --name $(BINARY_NAME) $(DOCKER_IMAGE):$(DOCKER_TAG)

docker-stop:
	@echo "Stopping Docker container..."
	docker stop $(BINARY_NAME) || true
	docker rm $(BINARY_NAME) || true

# Docker Compose commands
docker-compose-up:
	@echo "Starting services with Docker Compose..."
	docker-compose up -d

docker-compose-down:
	@echo "Stopping services with Docker Compose..."
	docker-compose down

docker-compose-logs:
	@echo "Viewing logs from Docker Compose services..."
	docker-compose logs -f
