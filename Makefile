# Git tag, if any and commit hash
VERSION := $(shell git describe --tags --long 2>/dev/null || git rev-parse --short HEAD)

# Directory for output, compiled files will go in $DIR/bin
BIN_DIR := "$(shell pwd)/bin"
LDFLAGS=-ldflags "-X main.Version=$(VERSION)"

# Go commands
GOINSTALL=go install -v $(LDFLAGS) ./...

# Check reflex installation
HAS_REFLEX := $(shell command -v reflex;)

.PHONY: all get build clean install image container destroy tools help

all: build

get: ## Import packages
	go get -v ./...

build: ## Build Go binary
	GOBIN=$(BIN_DIR) $(GOINSTALL)

clean: ## Clean bin directory
	rm -f $(BIN_DIR)/*
	rmdir $(BIN_DIR)

test: ## Run tests
	go test -v ./...

install: ## Install Go binary
	$(GOINSTALL)

image: ## Build docker image
	@docker-compose build

container: ## Run docker container
	@docker-compose up -d

destroy: ## Remove docker
	@docker-compose down

tools: reflex ## Download extra tools
	@echo 'Setup has completed.'

reflex:
ifndef HAS_REFLEX
	@sh -c 'GO111MODULE=off go get -u github.com/cespare/reflex'
else
	@echo 'Reflex has already installed.'
endif

help: ## Show usage
	@grep -E '^[ a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'
