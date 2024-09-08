# Simple Makefile for a Go project

all: build

# Build the application


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

# Changesets tasks
change-add:
	@pnpm dlx @changesets/cli add

change-chore:
	@pnpm dlx @changesets/cli add --empty

change-status:
	@pnpm dlx @changesets/cli status --since=origin/master

change-version:
	@pnpm dlx @changesets/cli version

change-tag:
	@pnpm dlx @changesets/cli tag

.PHONY: all build run test clean change-add change-empty change-status change-version change-tag
