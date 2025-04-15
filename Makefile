.PHONY: all build clean test run

BINARY_NAME=devops-mastery-server

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