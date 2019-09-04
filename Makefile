APP_NAME := goura
.DEFAULT_GOAL := help

GO_FILES_CMD := find . -name 'vendor' -prune -o -name '*.go' -print
GO_PATHS_CMD := $(GO_FILES_CMD) | awk -F/ '{ print $$1 "/" $$2 }' | uniq
PACKAGES_CMD := $(GO_PATHS_CMD) | grep -v '\.go$$' | awk -F/ '{ print $$0 "/..." }'
GO_TEST_CMD := go test -v -race -p 1 $(GO_TEST_FLAGS)
GO_PATHS := $(shell $(GO_PATHS_CMD))
PACKAGES := . $(shell $(PACKAGES_CMD))

ARTIFACT := ./bin/$(APP_NAME)
BIN_DIR := $(CURDIR)/bin

.PHONY: setup
setup: ## Set up dependencies
	@GOBIN=$(BIN_DIR) go install github.com/golangci/golangci-lint/cmd/golangci-lint
	@GOBIN=$(BIN_DIR) go install github.com/kyoh86/richgo

.PHONY: build
build: ## Build go binary
	@go build -o $(ARTIFACT)

.PHONY: docker.build
docker.build: ## Build docker image
	@docker build -f ./Dockerfile -t $(APP_NAME):latest .

.PHONY: docker.run
docker.run: ## Run on docker
	@docker run -it --rm $(APP_NAME):latest

.PHONY: lint
lint: ## Run static code analysis
	@bin/golangci-lint run --tests --disable-all --enable=goimports --enable=golint --enable=govet --enable=errcheck --enable=staticcheck --enable=gosec $(PACKAGES)

.PHONY: test
test: ## Run code test
	@$(GO_TEST_CMD) $(PACKAGES) | bin/richgo testfilter

.PHONY: help
help: ## Show options
	 @grep -E '^[a-zA-Z_-{\.}]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'
