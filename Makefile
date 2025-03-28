.PHONY: run build test coverage lint clean help

# Variáveis
APP_NAME := script-spotify-youtube
BIN_DIR := ./bin
MAIN_FILE := ./main.go
COVERAGE_FILE := coverage.out
COVERAGE_HTML := coverage.html

# Comandos
## Rodar a aplicação
run:
	@echo ">> Rodando a aplicação..."
	go run $(MAIN_FILE)

## Compilar a aplicação
build:
	@echo ">> Compilando o aplicativo para o diretório '$(BIN_DIR)'..."
	mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/$(APP_NAME) $(MAIN_FILE)
	@echo ">> Build concluído: $(BIN_DIR)/$(APP_NAME)"

## Rodar testes
test:
	@echo ">> Rodando testes..."
	go test ./...

## Gerar relatório de cobertura
coverage:
	@echo ">> Gerando relatório de cobertura..."
	go test ./... -coverprofile=$(COVERAGE_FILE)
	go tool cover -html=$(COVERAGE_FILE) -o $(COVERAGE_HTML)
	@echo ">> Relatório gerado: $(COVERAGE_HTML)"

## Rodar análise estática (lint)
lint:
	@echo ">> Rodando análise estática..."
	golangci-lint run || echo ">> Instale o golangci-lint para usar esse comando: https://golangci-lint.run/"

## Limpar arquivos gerados
clean:
	@echo ">> Limpando arquivos gerados..."
	rm -rf $(BIN_DIR) $(COVERAGE_FILE) $(COVERAGE_HTML)
	@echo ">> Limpeza concluída."

## Exibir ajuda
help:
	@echo "Comandos disponíveis:"
	@echo "  run        - Rodar a aplicação"
	@echo "  build      - Compilar a aplicação"
	@echo "  test       - Rodar testes"
	@echo "  coverage   - Gerar relatório de cobertura"
	@echo "  lint       - Rodar análise estática (lint)"
	@echo "  clean      - Limpar arquivos gerados"
	@echo "  help       - Exibir esta ajuda"
