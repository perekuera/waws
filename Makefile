# Makefile
PROGRAM_NAME := waws

SOURCE_DIR := ./src

INSTALL_DIR := /usr/local/bin

GO_VERSION := $(shell go version 2>/dev/null)

# Verificar si Go está instalado
ifeq ($(strip $(GO_VERSION)),)
$(error "Go no está instalado en el sistema. Por favor, instale Go antes de continuar.")
endif

build:
	@mkdir -p build
	@go build -o build/$(PROGRAM_NAME) $(SOURCE_DIR)/main.go
	@echo "Build process finished"

install: build
	@mkdir -p $(INSTALL_DIR)
	@sudo cp build/$(PROGRAM_NAME) $(INSTALL_DIR)/$(PROGRAM_NAME)
	@echo "Install process finished"

uninstall:
	@sudo rm -f $(INSTALL_DIR)/$(PROGRAM_NAME)
	@echo "Uninstall process finished"

clean:
	@rm -rf build
	@echo "Clean process finished"

run:
	@go run $(SOURCE_DIR)/main.go

help:
	@echo "Use:"
	@echo "  make [target]"
	@echo ""
	@echo "Available targets:"
	@echo "  build               Build program"
	@echo "  install             Build and installs program"
	@echo "  uninstall           Uninstall program"
	@echo "  clean               Clean build folder"
	@echo "  run                 Runs program"
