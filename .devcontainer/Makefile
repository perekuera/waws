# Makefile

# Nombre del programa
PROGRAM_NAME := waws

# Compilador Go
GO := go

# Opciones de compilación para Linux y Windows
LDFLAGS := -ldflags="-s -w"
WINDOWS_LDFLAGS := -ldflags="-s -w -H=windowsgui"

# Directorio de salida
BUILD_DIR := build

# Compilación para Linux
linux:
	@mkdir -p $(BUILD_DIR)
	$(GO) build $(LDFLAGS) -o $(BUILD_DIR)/$(PROGRAM_NAME)_linux_amd64

# Compilación para Windows
windows:
	@mkdir -p $(BUILD_DIR)
	$(GO) build $(WINDOWS_LDFLAGS) -o $(BUILD_DIR)/$(PROGRAM_NAME)_windows_amd64.exe

# Compilación para todas las plataformas
all: linux windows

# Eliminar los archivos compilados
clean:
	rm -rf $(BUILD_DIR)

# Ejecutar el programa
run:
	$(GO) run .

# Ayuda
help:
	@echo "Uso:"
	@echo "  make [target]"
	@echo ""
	@echo "Objetivos disponibles:"
	@echo "  linux   Compilar para Linux"
	@echo "  windows Compilar para Windows"
	@echo "  all     Compilar para todas las plataformas"
	@echo "  clean   Eliminar archivos compilados"
	@echo "  run     Ejecutar el programa (requiere tener Go instalado)"
