PROJECT_DIR = $(shell pwd)
PROJECT_BIN = $(PROJECT_DIR)/bin
$(shell [ -f bin ] || mkdir -p $(PROJECT_BIN))
PATH := $(PROJECT_BIN):$(PATH)

GOLANGCI_LINT = $(PROJECT_BIN)/golangci-lint
RELEASE_STR = $(shell git rev-parse --short HEAD)

.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/obsidian-bot-linux-amd64 .
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o bin/obsidian-bot-darwin-amd64 .

.PHONY: .install-linter
.install-linter:
	@ [ -f $(PROJECT_BIN)/golangci-lint ] || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(PROJECT_BIN) v1.55.2

.PHONY: lint
lint: .install-linter
	$(GOLANGCI_LINT) run ./... --config=./.golangci.yml

.PHONY: modules
modules:
	go mod tidy
	go mod vendor

.PHONY: test
test:
	go test ./...