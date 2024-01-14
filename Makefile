SHELL:=/usr/bin/env bash

GOPROXY := https://proxy.golang.com.cn
export GOPROXY
export GO111MODULE := on

IMG ?= demo:$(shell git describe --always --dirty --tags)

.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

.PHONY: build
build: ## Build target
	go build -o output/demo

.PHONY: docker-build
docker-build: ## Build the docker image
	docker build -t $(IMG) -f Dockerfile --build-arg LD_FLAGS="$(LD_FLAGS)" .

.PHONY: fmt
fmt: ## Run go fmt against code.
	go fmt ./...

lint: ## Run linting checks
	docker run --rm -v $(PWD):/app -w /app -e GOPROXY=${GOPROXY} golangci/golangci-lint:v1.47.2 golangci-lint run -v --timeout 5m

clean: ## Clean tmp files
	rm -fr output applog