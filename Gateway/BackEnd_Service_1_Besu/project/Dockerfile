# Use a imagem oficial do Golang como base
FROM golang:1.20 as builder

# Configura o diretório de trabalho dentro do container
WORKDIR /app

# Copia os arquivos go.mod e go.sum para o container
COPY go.mod go.sum ./

# Baixa as dependências
RUN go mod download

# Copia o restante do código-fonte do projeto para o container
COPY . .

# Compila o binário
RUN go build -o main .

# Cria uma nova imagem mínima para executar o binário
FROM debian:bullseye-slim

# Instala dependências necessárias para execução do binário
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

# Define o diretório de trabalho na imagem final
WORKDIR /app

# Copia o binário compilado da etapa anterior
COPY --from=builder /app/main .

# Copia os arquivos necessários para execução (como os arquivos de Swagger)
COPY --from=builder /app/docs ./docs

# Define a porta que será exposta
EXPOSE 8081

# Comando para rodar o binário
CMD ["./main"]