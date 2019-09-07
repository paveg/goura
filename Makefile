export GO111MODULE=on
APP_NAME := goura
REVISION := $(shell git rev-parse --short HEAD)
LDFLAGS := "-X github.com/paveg/goura/cmd.revision=$(REVISION)"
.DEFAULT_GOAL := help

GO_FILES_CMD := find . -name 'vendor' -prune -o -name '*.go' -print
GO_PATHS_CMD := $(GO_FILES_CMD) | awk -F/ '{ print $$1 "/" $$2 }' | uniq
PACKAGES_CMD := $(GO_PATHS_CMD) | grep -v '\.go$$' | awk -F/ '{ print $$0 "/..." }'
GO_TEST_CMD := go test -v -race -p 1 $(GO_TEST_FLAGS)
GO_PATHS := $(shell $(GO_PATHS_CMD))
PACKAGES := . $(shell $(PACKAGES_CMD))

ARTIFACT := ./bin/$(APP_NAME)
BIN_DIR := $(CURDIR)/bin

.PHONY: tools.setup
tools.setup: ## Set up tools
	@./init.sh

.PHONY: check
check: vet lint ## Run static code check

.PHONY: vet
vet: ## Run vet
	@go vet $(PACKAGES)

.PHONY: lint
lint: ## Run static lint for local
	@echo $(PACKAGES) | xargs -n 1 golint

.PHONY: docker.run
docker.run: ## Run on docker
	@docker run -it --rm $(APP_NAME):latest

.PHONY: ci.lint
ci.lint: tools.setup ## Run static lint for CI
	@$(BIN_DIR)/golangci-lint run --tests --disable-all --enable=goimports --enable=golint --enable=govet --enable=errcheck --enable=staticcheck --enable=gosec $(PACKAGES)

.PHONY: test
test: ## Run code test
	@$(GO_TEST_CMD) $(PACKAGES)

.PHONY: build
build: ## Build go binary
	@go build -ldflags $(LDFLAGS) -o $(ARTIFACT)

.PHONY: docker.build
docker.build: ## Build docker image
	@docker build -f ./Dockerfile -t $(APP_NAME):latest .

.PHONY: install
install: build ## Install go binary
	@echo "export $(PWD)/bin/goura into $(HOME)/bin/goura"
	@ln -sf $(PWD)/bin/goura $(HOME)/bin/goura

.PHONY: help
help: ## Show options
	 @grep -E '^[a-zA-Z_-{\.}]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'
