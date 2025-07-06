# Docker Setup - Besu Gateway

Este projeto contém configurações Docker para executar o gateway Besu em diferentes ambientes.

## Arquivos Docker

### `docker-compose.yml`
Configuração completa com:
- Gateway da aplicação
- Node Besu local para desenvolvimento
- Volumes para configurações e logs
- Network isolada

### `docker-compose.dev.yml`
Configuração apenas do gateway para desenvolvimento quando você já tem um node Besu externo rodando.

### `Dockerfile`
Imagem otimizada multi-stage para produção.

## Como usar

### 1. Ambiente completo (com node Besu):
```bash
# Sobe gateway + node Besu
docker-compose up -d

# Visualizar logs
docker-compose logs -f

# Parar serviços
docker-compose down
```

### 2. Apenas gateway (node Besu externo):
```bash
# Certificar que tem um node Besu rodando em localhost:8550
# Depois subir apenas o gateway
docker-compose -f docker-compose.dev.yml up -d
```

### 3. Build da imagem apenas:
```bash
docker build -t besu-gateway .
```

## Configuração

### Variáveis de ambiente importantes:
- `PORT`: Porta da aplicação (padrão: 8081)
- `GIN_MODE`: Modo do Gin (release/debug)
- `RPC_URL`: URL do node Besu
- `CHAIN_ID`: ID da blockchain

### Arquivos de configuração:
- `CCInfra/DB/aws_config.yaml`: Configurações AWS DynamoDB
- `CCInfra/RPC_besu/rpc_config.yaml`: Configurações do Besu/Ethereum
- `besu-config/genesis.json`: Genesis block para o node Besu

## Ports

### Gateway:
- **8081**: API REST

### Node Besu (quando usando docker-compose.yml completo):
- **8545**: HTTP RPC
- **8546**: WebSocket RPC  
- **8550**: HTTP RPC (adicional)
- **30303**: P2P

## Volumes

- `besu-logs`: Logs da aplicação
- `besu-data`: Dados do node Besu (blockchain)

## Health Check

O gateway inclui health check em `/health` que verifica:
- Status da aplicação
- Conectividade com banco de dados
- Conectividade com node Besu

## Desenvolvimento

Para desenvolvimento local sem Docker:
1. Ter Go 1.23+ instalado
2. Ter node Besu rodando em localhost:8550
3. Executar: `go run main.go`

## Produção

Para produção, ajustar:
1. Variáveis de ambiente AWS
2. URLs de nodes Besu externos
3. Configurações de segurança
4. Limites de recursos no docker-compose
