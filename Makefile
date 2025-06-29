# Makefile para Design Patterns
# Comandos para executar exemplos específicos

.PHONY: help builder strategy clean

# Comando padrão - mostra ajuda
help:
	@echo "Comandos disponíveis:"
	@echo "  make builder    - Executa o exemplo do padrão Builder"
	@echo "  make strategy   - Executa o exemplo do padrão Strategy"
	@echo "  make clean      - Remove arquivos compilados"
	@echo "  make help       - Mostra esta ajuda"

# Executa o exemplo do padrão Builder
builder:
	@echo "Executando exemplo do padrão Builder..."
	@cd creational/builder/example && go run main.go

# Executa o exemplo do padrão Strategy
strategy:
	@echo "Executando exemplo do padrão Strategy..."
	@cd behavioural/strategy && go run main.go

# Remove arquivos compilados
clean:
	@echo "Limpando arquivos compilados..."
	@find . -name "*.exe" -delete
	@find . -name "*.out" -delete
	@find . -name "*.bin" -delete
	@find . -name "*.tmp" -delete

# Comando para executar todos os exemplos
all: builder strategy
	@echo "Todos os exemplos foram executados!"

# Comando para listar todos os padrões disponíveis
list:
	@echo "Padrões de Design disponíveis:"
	@echo "  - Builder (creational/builder/example/)"
	@echo "  - Strategy (behavioural/strategy/)"
	@echo ""
	@echo "Use 'make <padrão>' para executar um exemplo específico"
	@echo "Exemplo: make builder" 