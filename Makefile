BINARY_NAME=twitter_clone
BINARY_PATH=./bin/${BINARY_NAME}
GOCMD=go

GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
CYAN   := $(shell tput -Txterm setaf 6)
RESET  := $(shell tput -Txterm sgr0)

all: help

install: ## Install dependencies
	@$(GOCMD) install github.com/a-h/templ/cmd/templ@latest
	@$(GOCMD) get ./...
	@$(GOCMD) mod vendor
	@$(GOCMD) mod tidy
	@$(GOCMD) mod download

run: build ## Run application
	@$(BINARY_PATH)

build: ## Build application
	@echo "building template files..."
	@templ generate -path views/
	@echo "building app..."
	@npx tailwindcss -c ./views/tailwind.config.js -i ./views/app.css -o ./views/public/app.css
	@$(GOCMD) build -o ${BINARY_PATH} ./cmd/app/

test: ## Run the tests of the project
	@make vet
	@make lint
	@make test_race_coverage

test_race_coverage: ## Runs the tests with race and coverage
	$(GOCMD) test -race ./... -coverprofile=coverage.out

vet: ## Vets the project
	$(GOCMD) vet -v

lint: ## Lints the project
	golines -w -l .
	goimports -w -l .
	gofumpt -l -w .

help: ## Show this help.
	@echo ''
	@echo 'Usage:'
	@echo '  ${YELLOW}make${RESET} ${GREEN}<target>${RESET}'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} { \
		if (/^[a-zA-Z_-]+:.*?##.*$$/) {printf "    ${YELLOW}%-20s${GREEN}%s${RESET}\n", $$1, $$2} \
		else if (/^## .*$$/) {printf "  ${CYAN}%s${RESET}\n", substr($$1,4)} \
		}' $(MAKEFILE_LIST)

