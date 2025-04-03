# Variáveis
GO = go
CMD_DIR = cmd
MAIN_FILE = $(CMD_DIR)/main.go

# Target para rodar o projeto
run:
	$(GO) run $(MAIN_FILE)

# Target para rodar as migrações do banco de dados
migrate:
	# Adapte para o comando de migração que você estiver usando
	$(GO) run ./db/migrate.go

# Target para rodar o servidor com o Gin
start:
	$(GO) run $(MAIN_FILE)

# Target para rodar testes (caso queira rodar testes no seu projeto)
test:
	$(GO) test ./...

# Target para compilar o binário do projeto
build:
	$(GO) build -o myproject ./cmd

# Target para limpar arquivos gerados (como binários ou arquivos de cache)
clean:
	rm -f myproject

# Alvo default que será executado ao rodar "make" sem parâmetros
default: run
