# Simple Makefile for a Go project

# Build the application
all: build

build:
	@echo "Building..."

	@CGO_ENABLED=0 go build -o build/birthday cmd/birthday/main.go

# Run the application
run:
	@go run cmd/birthday/main.go

# Test the application
test:
	@echo "Testing..."
	@go test ./test/... -v

# Clean the binary
clean:
	@echo "Cleaning..."
	@rm -rfd build

.PHONY: all build run test clean
