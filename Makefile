
.PHONY: build
build: ## Builds the application for production
	go build -ldflags="-w -s" -v -o .

.PHONY: clean
clean: ## Runs mod tidy
	go mod tidy

.PHONY: update
update: ## Update go modules
	go get -t -u ./...


.PHONY: run
run: ## Execute the application locally
	go run -race ./cmd/pokemon-red-study

.PHONY: format
format: ## Format code and organize imports
	go fmt ./...
	fieldalignment -fix ./...

.PHONY: lint
lint: ## Runs golangci-lint
	golangci-lint run

.PHONY: mocks
mocks: ## Generate mocks
	find internal -type d -name "mocks" -exec rm -rf {} +
	mockery
