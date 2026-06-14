.PHONY: build test docker fleet-status lint clean fmt vet

# Binary name
BINARY = openagent

# Go parameters
GOCMD = go
GOBUILD = $(GOCMD) build
GOTEST = $(GOCMD) test
GOVET = $(GOCMD) vet

# Fleet edge worker endpoint
FLEET_EDGE_URL ?= https://fleet-edge-worker.casey-digennaro.workers.dev

# Docker
DOCKER_CMD = docker
DOCKER_IMAGE = superinstance/openagent:latest

## build: Compile the openagent binary
build:
	$(GOBUILD) -o $(BINARY) .

## test: Run all tests
test:
	$(GOTEST) ./... -v -count=1

## test-short: Run only short tests
test-short:
	$(GOTEST) ./... -short -count=1

## docker: Build Docker image
docker:
	$(DOCKER_CMD) build -t $(DOCKER_IMAGE) .

## fleet-status: Query the fleet-edge-worker API for live fleet status
fleet-status:
	@echo "🔍 Querying fleet-edge-worker at $(FLEET_EDGE_URL)..."
	@curl -sS --connect-timeout 10 --max-time 30 "$(FLEET_EDGE_URL)/api/fleet/status" | python3 -m json.tool 2>/dev/null || \
	curl -sS --connect-timeout 10 --max-time 30 "$(FLEET_EDGE_URL)/fleet/status" | python3 -m json.tool 2>/dev/null || \
	echo "⚠️  Could not reach fleet-edge-worker. Is the worker deployed?"

## lint: Run golangci-lint
lint:
	golangci-lint run ./...

## vet: Run go vet
vet:
	$(GOVET) ./...

## fmt: Format all Go code
fmt:
	$(GOCMD) fmt ./...

## clean: Remove build artifacts
clean:
	rm -f $(BINARY)
	rm -f coverage.out
	$(GOCMD) clean -testcache

## all: Build and test
all: build test

## help: Show this help
help:
	@echo "Available targets:"
	@grep -E '^## ' $(MAKEFILE_LIST) | sed 's/## //' | column -t
