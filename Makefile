# CS-GO Project Makefile

# Variables
BINARY_NAME=cs-go.out
VERSION=1.0.0
BUILD_DIR=./
MAIN_FILE=main.go
STORAGE_DIR=./storage

# Go commands
GO=go
GOBUILD=$(GO) build
GOCLEAN=$(GO) clean
GOTEST=$(GO) test
GOGET=$(GO) get
GOVET=$(GO) vet

# Determine the operating system
UNAME_S := $(shell uname -s)
ifeq ($(UNAME_S),Darwin)
	# macOS settings
	OPEN=open
else
	# Linux/other settings
	OPEN=xdg-open
endif

# Default target
.PHONY: all
all: build

# Build the project
.PHONY: build
build:
	@./scripts/build.sh

# Clean the project
.PHONY: clean
clean:
	@echo "Cleaning..."
	@$(GOCLEAN)
	@rm -f $(BUILD_DIR)/$(BINARY_NAME)
	@echo "Clean successful!"

# Run tests
.PHONY: test
test:
	@./scripts/test.sh

# Run the application
.PHONY: run
run: build
	@./scripts/run.sh

# Install dependencies
.PHONY: deps
deps:
	@echo "Installing dependencies..."
	@$(GOGET) -v ./...
	@echo "Dependencies installed!"

# Format the code
.PHONY: fmt
fmt:
	@echo "Formatting code..."
	@$(GO) fmt ./...
	@echo "Formatting complete!"

# Run go vet
.PHONY: vet
vet:
	@echo "Running go vet..."
	@$(GOVET) ./...
	@echo "Vet complete!"

# Clean database
.PHONY: clean-db
clean-db:
	@echo "Cleaning database..."
	@mkdir -p $(STORAGE_DIR)
	@rm -f $(STORAGE_DIR)/db.json
	@echo "Database cleaned!"

# Build with version
.PHONY: release
release:
	@echo "Building release version $(VERSION)..."
	@$(GOBUILD) -o $(BUILD_DIR)/cs-go-v$(VERSION).out $(MAIN_FILE)
	@chmod +x $(BUILD_DIR)/cs-go-v$(VERSION).out
	@echo "Release build successful!"

# Help target
.PHONY: help
help:
	@echo "CS-GO Makefile Help"
	@echo "===================="
	@echo "make               - Build the application"
	@echo "make build         - Build the application"
	@echo "make clean         - Remove built binary and clean Go cache"
	@echo "make test          - Run tests"
	@echo "make run           - Build and run the application"
	@echo "make deps          - Install dependencies"
	@echo "make fmt           - Format the code"
	@echo "make vet           - Run go vet"
	@echo "make clean-db      - Clean the database files"
	@echo "make release       - Build a release version"
	@echo "make help          - Show this help message" 