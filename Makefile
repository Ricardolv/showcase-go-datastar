# Makefile para Showcase Go Datastar

# Variáveis
BINARY_NAME=showcase
BUILD_DIR=bin
TMP_DIR=tmp

# Comandos padrão
.PHONY: help dev build clean install test lint format deps

# Help
help: ## Mostra esta ajuda
	@echo "Showcase Go Datastar - Comandos disponíveis:"
	@echo ""
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

# Desenvolvimento
dev: ## Inicia o servidor de desenvolvimento com hot reload
	@echo "🚀 Iniciando desenvolvimento..."
	@npm run dev

# Setup inicial
install: ## Instala todas as dependências
	@echo "📦 Instalando dependências..."
	@go mod download
	@npm install
	@go install github.com/air-verse/air@latest
	@go install github.com/a-h/templ/cmd/templ@latest

# Build CSS
css-dev: ## Build CSS para desenvolvimento (com watch)
	@echo "🎨 Building CSS (dev mode)..."
	@npm run css:dev

css-prod: ## Build CSS para produção (minificado)
	@echo "🎨 Building CSS (prod mode)..."
	@npm run css:prod

# Geração de templates
templ-generate: ## Gera templates templ
	@echo "📝 Gerando templates..."
	@templ generate

templ-watch: ## Observa mudanças nos templates
	@echo "👀 Observando templates..."
	@templ generate --watch

# Build
build: css-prod templ-generate ## Build completo para produção
	@echo "🏗️ Building aplicação..."
	@mkdir -p $(BUILD_DIR)
	@go build -ldflags="-s -w" -o $(BUILD_DIR)/$(BINARY_NAME) ./cmd/server
	@echo "✅ Build completo! Binary: $(BUILD_DIR)/$(BINARY_NAME)"

# Build com informações de debug
build-debug: css-prod templ-generate ## Build com símbolos de debug
	@echo "🔍 Building com debug..."
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(BINARY_NAME) ./cmd/server

# Cross compilation
build-all: css-prod templ-generate ## Build para todas as plataformas
	@echo "🌍 Building para múltiplas plataformas..."
	@mkdir -p $(BUILD_DIR)
	@GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o $(BUILD_DIR)/$(BINARY_NAME)-linux-amd64 ./cmd/server
	@GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-amd64 ./cmd/server
	@GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-arm64 ./cmd/server
	@GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o $(BUILD_DIR)/$(BINARY_NAME)-windows-amd64.exe ./cmd/server
	@echo "✅ Builds completos para todas as plataformas!"

# Execução
run: ## Executa a aplicação
	@echo "▶️ Executando aplicação..."
	@go run ./cmd/server

run-prod: build ## Executa o binary de produção
	@echo "🚀 Executando binary de produção..."
	@./$(BUILD_DIR)/$(BINARY_NAME)

# Limpeza
clean: ## Remove arquivos de build
	@echo "🧹 Limpando arquivos de build..."
	@rm -rf $(BUILD_DIR)
	@rm -rf $(TMP_DIR)
	@rm -rf web/static/css/styles.css
	@rm -f **/*_templ.go

# Testes
test: ## Executa todos os testes
	@echo "🧪 Executando testes..."
	@go test -v ./...

test-coverage: ## Executa testes com coverage
	@echo "📊 Executando testes com coverage..."
	@go test -v -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage disponível em coverage.html"

# Qualidade de código
lint: ## Executa linting
	@echo "🔍 Executando lint..."
	@go vet ./...

format: ## Formata o código
	@echo "✨ Formatando código..."
	@go fmt ./...
	@templ fmt .

# Dependências
deps: ## Atualiza dependências Go
	@echo "🔄 Atualizando dependências..."
	@go mod tidy
	@go mod download

deps-upgrade: ## Upgrade das dependências
	@echo "⬆️ Fazendo upgrade das dependências..."
	@go get -u ./...
	@go mod tidy

# Deployment
deploy-prep: build ## Prepara arquivos para deploy
	@echo "📦 Preparando deploy..."
	@mkdir -p deploy
	@cp $(BUILD_DIR)/$(BINARY_NAME) deploy/
	@cp -r web/static deploy/
	@echo "✅ Arquivos de deploy prontos em ./deploy/" 