# Makefile para Go Fear the Foca

APP_NAME := go-fear-the-foca
CMD_PATH := ./cmd/$(APP_NAME)
BIN_PATH := ./bin

.PHONY: all install dev build run clean test lint fmt

all: build

install:
	@echo "Instalando dependencias..."
	@go mod download
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@sudo apt install exiftool

dev:
	@echo "Instalando dependencias de desarrollo..."
	@go run cmd/go-fear-the-foca/main.go -domain ejemplo.com


build:
	@echo "Compilando la aplicación..."
	@mkdir -p $(BIN_PATH)
	@go build -o $(BIN_PATH)/$(APP_NAME) $(CMD_PATH)

run: build
	@echo "Ejecutando la aplicación..."
	@$(BIN_PATH)/$(APP_NAME) $(ARGS)

clean:
	@echo "Limpiando archivos binarios..."
	@rm -rf $(BIN_PATH)

test:
	@echo "Ejecutando pruebas..."
	@go test ./...

lint:
	@echo "Ejecutando linter..."
	@golangci-lint run

fmt:
	@echo "Formateando el código..."
	@go fmt ./...
