# Project information
VERSION? := 0# $(shell git describe --tags)
BUILD := $(shell git rev-parse --short HEAD)
PROJECTNAME := $(shell basename "$(PWD)")

# Go build variables
GOBASE := $(shell pwd)
GOPATH := $(GOBASE)/vendor:$(GOBASE)
GOBIN := $(GOBASE)/bin
GOFILES := $(wildcard *.go)

CMD := $(GOBASE)/cmd

# Linker flags
LDFLAGS=-v -ldflags "-X=main.Version=$(VERSION) -X=main.Build=$(BUILD)"

.PHONY: all build clean lint

all: suite 

full-suite: test cover bench ## Run all tests
fast-suite: test cover ## Just test the lib and generate coverage

lint: ## Lint the files
	@printf "\033[36m%-30s\033[0m\n" "Lint source code"
	@golint pkg/...
	@golint cmd/...

build: test ## Build the binary file
	@printf "\033[36m%-30s\033[0m\n" "Build binaries"
	@mkdir -p ./bin/limbo
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go build $(LDFLAGS) -o $(GOBIN)/limbo/limbo $(CMD)/$(PROJECTNAME)

bench: lint ## Run the benchmarks
	@printf "\033[36m%-30s\033[0m\n" "Run benchmarks"
	@cd pkg/life
	@go test ./... -bench=. -benchmem

test: lint ## Test the library
	@printf "\033[36m%-30s\033[0m\n" "Perform covered tests"
	@go test -race ./... -coverprofile artifacts/cover.out

cover: test ## Generate coverage
	@printf "\033[36m%-30s\033[0m\n" "Output coverage"
	@go tool cover -html=artifacts/cover.out -o artifacts/cover.html
	@go tool cover -func=artifacts/cover.out

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'