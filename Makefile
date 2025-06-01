# Variables
BINARY_NAME = munchies
CMD_DIR = ./cmd/munchies
BUILD_DIR = ./bin

.PHONY: all
all: help ## Default target

.PHONY: build
build: ## Build the binary
	@echo "Building the binary..."
	@mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(BINARY_NAME) $(CMD_DIR)

.PHONY: run
run: build ## Run the application
	@echo "Running the application..."
	$(BUILD_DIR)/$(BINARY_NAME)

.PHONY: test
test: ## Run tests
	@echo "Running tests..."
	go test ./... -v

.PHONY: clean
clean: ## Clean up build artifacts
	@echo "Cleaning up..."
	@rm -rf $(BUILD_DIR)

deep-clean: ## clean .munchies directory and saved data.
	@echo "Deep cleaning..."
	@rm -rf ${HOME}/.munchies

install: build ## Install the binary to /usr/local/bin
	@echo "Installing the binary..."
	@mv $(BUILD_DIR)/$(BINARY_NAME) /usr/local/bin/

.PHONY: fmt
fmt: ## Format the code
	@echo "Formatting code..."
	go fmt ./...

.PHONY: cmd
cmd: build ## Run the application with a specific command (e.g., make cmd ARGS="help")
	@echo "Running with arguments: $(ARGS)"
	$(BUILD_DIR)/$(BINARY_NAME) $(ARGS)

.PHONY: help
help: ## Show this help.
	@egrep -h '\s##\s' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m  %-30s\033[0m %s\n", $$1, $$2}'
