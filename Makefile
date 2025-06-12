# Variables
BINARY_NAME = munchies
CMD_DIR = ./cmd/munchies
BUILD_DIR = ./bin

MUNCHIES_COMMAND ?= help
MUNCHIES_INSTALL_DIR ?= /usr/local/bin
MUNCHIES_UNINSTALL_DIR ?= /usr/local/bin

.PHONY: all
all: help ## Default target

.PHONY: build
build: ## Build the binary
	@echo "Building the binary..."
	@mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(BINARY_NAME) $(CMD_DIR)

.PHONY: run
run: build ## Run the application use MUNCHIES_COMMAND to specify the command to run
	@echo "Running the application..."
	$(BUILD_DIR)/$(BINARY_NAME) $(MUNCHIES_COMMAND)

.PHONY: test
test: ## Run tests
	@echo "Running tests..."
	go test ./... -v

.PHONY: clean
clean: ## Clean up build artifacts
	@echo "Cleaning up..."
	@rm -rf $(BUILD_DIR)

.PHONY: deep-clean
deep-clean: ## clean .munchies directory and saved data
	@read -p "Are you sure you want to delete all munchies data? [y/N] " ans; \
	if [ "$$ans" = "y" ] || [ "$$ans" = "Y" ]; then \
		echo "Deleting .munchies directory and saved data..."; \
		rm -rf $(HOME)/.munchies; \
	else \
		echo "Aborted."; \
	fi

.PHONY: install
install: build ## Install the binary to MUNCHIES_INSTALL_DIR. Defaults to /usr/local/bin.
	@echo "Installing the binary to $(MUNCHIES_INSTALL_DIR)..."
	@mv $(BUILD_DIR)/$(BINARY_NAME) $(MUNCHIES_INSTALL_DIR)

.PHONY: uninstall
uninstall: ## unstall the binary from specified directory. Defaults to /usr/local/bin
	@echo "Uninstalling the binary from $(MUNCHIES_UNINSTALL_DIR)..."
	@rm $(MUNCHIES_UNINSTALL_DIR)/$(BINARY_NAME) || true

.PHONY: fmt
fmt: ## Format the code
	@echo "Formatting code..."
	go fmt ./...

.PHONY: help
help: ## Show this help.
	@egrep -h '\s##\s' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m  %-30s\033[0m %s\n", $$1, $$2}'
