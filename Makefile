# binary aliases
GO = GO111MODULE=on go
BUF = buf

##@ Dependencies
.PHONY: install install-go install-buf

install: install-go install-buf ## Install all dependencies

install-go: ## Install Go dependencies.
	$(GO) run build.go setup

install-proto: ## Install protobuf dependencies
	$(BUF) beta mod update

##@ Lint
.PHONY: lint lint-go lint-protobuf

lint: lint-go lint-proto ## Run all linters

lint-go: ## Lint go code
	@echo "TODO"

lint-proto: ## Lint protobuf definitions
	$(BUF) lint

##@ Build
.PHONY: build build-docker build-protobuf

build: ## Build Go binaries.
	@echo "build go files"
	$(GO) run scripts/build.go build

build-docker: ## Build Docker image for development.
	@echo "build docker container"
	docker build --tag grumblechat/server:dev -f build/docker/Dockerfile .

build-proto: ## Compile protobuf definitions
#	./scripts/protobuf-check.sh
	$(BUF) generate
	$(BUF) build -o buf.snapshot

##@ Testing
.PHONY: test test-full

test: ## short test-suite
	@echo "running minimal tests"
	$(GO) test --short ./...

test-full: ## full test-suite
	@echo "running all tests"
	$(GO) test -v ./...

test-proto: ## check for protobuf breaking changes
	$(BUF) breaking --against buf.snapshot

##@ Helpers
.PHONY: start proto-snapshot help

start: build ## start grumble-server locally
	@echo "starting grumble-server"
	./bin/*/grumble-server

help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)
