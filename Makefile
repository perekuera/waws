# Makefile

# Nombre del programa
PROGRAM_NAME := waws

# Directorio de instalación
INSTALL_DIR := /usr/local/bin

# Compilación
build:
	@mkdir -p build
	go build -o build/$(PROGRAM_NAME)

# Instalación (depende de build)
install: build
	@mkdir -p $(INSTALL_DIR)
	cp build/$(PROGRAM_NAME) $(INSTALL_DIR)/$(PROGRAM_NAME)

# Desinstalación
uninstall:
	rm -f $(INSTALL_DIR)/$(PROGRAM_NAME)

# Limpieza
clean:
	rm -rf build

# Ejecutar el programa
run:
	go run .

# Ayuda
help:
	@echo "Uso:"
	@echo "  make [target]"
	@echo ""
	@echo "Objetivos disponibles:"
	@echo "  build               Compilar el programa"
	@echo "  install             Instalar el programa"
	@echo "  uninstall           Desinstalar el programa"
	@echo "  clean               Eliminar archivos compilados"
	@echo "  run                 Ejecutar el programa (requiere tener Go instalado)"
