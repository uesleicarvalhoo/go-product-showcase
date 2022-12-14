COVERAGE_OUTPUT=coverage.output
COVERAGE_HTML=coverage.html
GO_PACKAGES=cmd internal pkg
GO_ENTRYPOINT=cmd/api/*.go

## @ Help
.PHONY: help
help: ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make [target]\033[36m\033[0m\n\nTargets:\n"} /^[a-zA-Z_-]+:.*?##/ { printf "\033[36m%-10s\033[0m %s\n", $$1, $$2 }' $(MAKEFILE_LIST)

## @ Application
.PHONY: run compose swagger
run: swagger ## Run app
	@DEBUG=TRUE go run $(GO_ENTRYPOINT)

docs/*: $(wildcard cmd/api/server.go) $(wildcard cmd/api/services/*.go) $(wildcard cmd/api/services/*/*.go)
	@swag init -generalInfo $(GO_ENTRYPOINT) -output ./docs

swagger: docs/*  ## Generate Swagger content

compose:  ## Init containers with dev dependencies
	@docker compose build && docker compose up -d

release:  ## Create a new release
	@echo "Input version[$(shell git describe --abbrev=0 --tags --always)]:"
	@read INPUT_VERSION; \
	echo "Creating a new release version: $$INPUT_VERSION" \
	&& git tag "$$INPUT_VERSION" \
	&& git push origin "$$INPUT_VERSION" \
	&& git push origin -u "$(shell git rev-parse --abbrev-ref HEAD)"


## @ Linter
.PHONY: lint format
lint:
	@golangci-lint run -v

format:
	@gofumpt -e -l -w $(GO_PACKAGES)

## @ Tests
.PHONY: test coverage
test:  ## Run tests of project
	@go test ./... -race -v -count=1

coverage: ## Run tests, make report and open into browser
	@go test ./... -race -v -cover  -covermode=atomic -coverprofile=$(COVERAGE_OUTPUT)
	@go tool cover -html=$(COVERAGE_OUTPUT) -o $(COVERAGE_HTML)
	@wslview ./$(COVERAGE_HTML) || xdg-open ./$(COVERAGE_HTML) || powershell.exe Invoke-Expression ./$(COVERAGE_HTML)

## @ Clean
.PHONY: clean clean_coverage_cache
clean_coverage_cache:
	@rm -rf $(COVERAGE_OUTPUT)
	@rm -rf $(COVERAGE_HTML)

clean: clean_coverage_cache ## Remove Cache files
