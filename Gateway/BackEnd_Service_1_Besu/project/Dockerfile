# Build stage
FROM golang:1.23-alpine AS builder

# Instalar dependências necessárias
RUN apk add --no-cache git ca-certificates tzdata

# Definir diretório de trabalho
WORKDIR /app

# Copiar arquivos de dependências primeiro (para melhor cache do Docker)
COPY go.mod go.sum ./

# Baixar dependências
RUN go mod download

# Copiar código fonte
COPY . .

# Limpar e organizar dependências
RUN go mod tidy

# Build da aplicação com otimizações
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags='-w -s -extldflags "-static"' \
    -a -installsuffix cgo \
    -o main .

# Final stage
FROM alpine:latest

# Instalar certificados CA, timezone e bash
RUN apk --no-cache add ca-certificates tzdata bash curl

# Criar usuário não-root
RUN adduser -D -s /bin/sh appuser

# Criar diretórios necessários
RUN mkdir -p /app/CCInfra/DB \
    && mkdir -p /app/CCInfra/RPC_besu \
    && mkdir -p /app/logs \
    && mkdir -p /app/config

# Definir diretório de trabalho
WORKDIR /app

# Copiar binário da aplicação
COPY --from=builder /app/main .

# Copiar script de inicialização
COPY --from=builder /app/Main_run.sh .

# Criar script wrapper para debug
RUN echo '#!/bin/bash' > /app/start.sh && \
    echo 'echo "=== Debug Info ===" ' >> /app/start.sh && \
    echo 'echo "Current directory: $(pwd)"' >> /app/start.sh && \
    echo 'echo "Working directory: $PWD"' >> /app/start.sh && \
    echo 'echo "Files in /app:"' >> /app/start.sh && \
    echo 'ls -la /app/' >> /app/start.sh && \
    echo 'echo "Main_run.sh permissions:"' >> /app/start.sh && \
    echo 'ls -la /app/Main_run.sh' >> /app/start.sh && \
    echo 'echo "=== Starting Main_run.sh ===" ' >> /app/start.sh && \
    echo 'cd /app' >> /app/start.sh && \
    echo 'exec /app/Main_run.sh exec' >> /app/start.sh && \
    chmod +x /app/start.sh

# Copiar arquivos de configuração
COPY --from=builder /app/CCInfra/DB/aws_config.yaml ./CCInfra/DB/
COPY --from=builder /app/CCInfra/RPC_besu/rpc_config.yaml ./CCInfra/RPC_besu/

# Dar permissões apropriadas
RUN chmod +x /app/main \
    && chmod +x /app/Main_run.sh \
    && chown -R appuser:appuser /app

# Mudar para usuário não-root
USER appuser

# Expor porta
EXPOSE 8081

# Definir variáveis de ambiente padrão
ENV GIN_MODE=release
ENV PORT=8081

# Garantir que o working directory seja /app
WORKDIR /app

# Comando de inicialização usando script wrapper com debug
CMD ["/app/start.sh"]