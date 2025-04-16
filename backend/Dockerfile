# Usa a imagem oficial do Golang para construir e executar a aplicação
FROM golang:1.24.1 AS builder

# Define o diretório de trabalho dentro do container
WORKDIR /app

# Copia os arquivos do projeto para dentro do container
COPY . .

# Instala as dependências necessárias (com go mod tidy)
RUN go mod tidy 

# Usa uma imagem mais leve para rodar a aplicação
FROM golang:1.24.1

# Define o diretório de trabalho dentro do container
WORKDIR /app

# Copia os arquivos do projeto para dentro do container
COPY --from=builder /app .

# Expõe a porta usada pela aplicação
EXPOSE 8080

# Comando para rodar a aplicação diretamente com `go run`
CMD ["go", "run", "cmd/main.go"]

# Definir variável de ambiente (se necessário)
ENV GOPROXY=direct
