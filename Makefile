.PHONY: install build build-docker start test test-full help

GO = GO111MODULE=on go

##@ Dependencies

install: ## Install dependencies.
	$(GO) run build.go setup


##@ Build

build: ## Build Go binaries.
	@echo "build go files"
	$(GO) run scripts/build.go build

build-docker: ## Build Docker image for development.
	@echo "build docker container"
	docker build --tag grumblechat/server:dev -f build/docker/Dockerfile .

start: build ## start grumble-server locally
	@echo "starting grumble-server"
	./bin/*/grumble-server

##@ Testing

test: ## short test-suite
	@echo "running minimal tests"
	$(GO) test --short ./...

test-full: ## full test-suite
	@echo "running all tests"
	$(GO) test -v ./...

##@ Helpers

help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)
