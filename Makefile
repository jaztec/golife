.PHONY: all lint

all: full-suite 

full-suite: fast-suite bench ## Run all tests
fast-suite: test cover ## Just test the lib and generate coverage

lint: ## Lint the files
	@printf "\033[36m%-30s\033[0m\n" "Lint source code"
	@golint ./...
	@printf "done\n"

bench: lint ## Run the benchmarks
	@printf "\033[36m%-30s\033[0m\n" "Run benchmarks"
	@go test ./... -bench=. -benchmem

test: lint ## Test the library
	@printf "\033[36m%-30s\033[0m\n" "Perform covered tests"
	@go test -race ./... -coverprofile artifacts/cover.out

cover: test ## Generate coverage
	@printf "\033[36m%-30s\033[0m\n" "Output coverage"
	@mkdir -p artifacts
	@go tool cover -html=artifacts/cover.out -o artifacts/cover.html
	@go tool cover -func=artifacts/cover.out

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'