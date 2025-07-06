# Fabric Gateway Service

Serviço Gateway para integração com Hyperledger Fabric e AWS DynamoDB.

## 🚀 Características

- **Hyperledger Fabric**: Integração com blockchain Fabric
- **AWS DynamoDB**: Armazenamento de dados de transações
- **Docker**: Containerização completa
- **Hot Reload**: Desenvolvimento com recarga automática
- **Health Check**: Monitoramento de saúde do serviço
- **Swagger**: Documentação automática da API

## 📋 Pré-requisitos

- Docker e Docker Compose
- Go 1.23+ (para desenvolvimento local)
- AWS credentials configuradas
- Rede Hyperledger Fabric configurada

## 🛠️ Configuração

### 1. Configuração dos arquivos YAML

Edite os arquivos de configuração conforme necessário:

#### `CCInfra/DB/aws_config.yaml`
```yaml
aws:
  credentials:
    access_key_id: "sua_access_key"
    secret_access_key: "sua_secret_key"
  config:
    region: "us-east-1"
    output: "json"
  dynamodb:
    endpoint: ""
    table_prefix: "FabricTransactionHashTCC"
```

#### `CCInfra/Grpc_Hyperledger/fabric_config.yaml`
```yaml
fabric:
  msp_id: "Org1MSP"
  crypto_path: "./organizations/peerOrganizations/org1.example.com"
  peer_endpoint: "dns:///localhost:7051"
  gateway_peer: "peer0.org1.example.com"
  channel_name: "channeltcc"
  chaincode_name: "chaincodetcc"
```

### 2. Variáveis de Ambiente

Crie o arquivo `.env`:
```bash
cp .env.example .env
# Edite o arquivo .env com suas configurações
```

## 🚀 Execução

### Produção

```bash
# Verificar configurações
make check-config

# Construir e iniciar
make build
make up

# Ver logs
make logs

# Verificar status
make health
```

### Desenvolvimento (Hot Reload)

```bash
# Iniciar em modo desenvolvimento
make dev-up

# Ver logs de desenvolvimento
make dev-logs

# Parar desenvolvimento
make dev-down
```

### Comandos úteis

```bash
# Ver todos os comandos disponíveis
make help

# Reiniciar serviços
make restart

# Limpeza completa
make clean

# Deploy completo
make deploy
```

## 📚 API Endpoints

### Health Check
```
GET /health
```

### Batch Operations
```
POST /batch          # Criar novo batch
GET  /batch/:id      # Consultar batch por ID
```

### Documentação
```
GET /swagger/*       # Swagger UI
```

## 🏗️ Estrutura do Projeto

```
project/
├── CCDocs/                     # Documentação Swagger
├── CCEntitys/                  # Entidades do domínio
├── CCHandler/                  # Handlers HTTP
├── CCInfra/
│   ├── DB/
│   │   ├── aws_config.yaml     # Configuração AWS
│   │   ├── config_DB_aws.go    # Gerenciador de config AWS
│   │   └── dynamodb_connection.go
│   └── Grpc_Hyperledger/
│       ├── fabric_config.yaml  # Configuração Fabric
│       └── grpc_connection.go
├── organizations/              # Certificados Fabric
├── docker-compose.yml          # Produção
├── docker-compose.dev.yml      # Desenvolvimento
├── Dockerfile                  # Build produção
├── Dockerfile.dev             # Build desenvolvimento
├── Makefile                   # Comandos automatizados
└── main.go                    # Aplicação principal
```

## 🔧 Desenvolvimento

### Estrutura de Configuração

O projeto utiliza uma arquitetura modular com:

- **ConfigManager**: Classe para gerenciar configurações AWS
- **Separação de Responsabilidades**: Configurações isoladas do código de negócio
- **Hot Reload**: Mudanças no código refletidas automaticamente
- **Multi-stage Builds**: Otimização de imagens Docker

### Debugging

```bash
# Acessar shell do container
make shell

# Ver logs detalhados
make debug-logs

# Desenvolvimento com debugging
make dev-up
# Conectar debugger na porta 2345
```

## 🐳 Docker

### Imagens Disponíveis

- **Produção**: Imagem otimizada com Alpine Linux
- **Desenvolvimento**: Imagem com ferramentas de desenvolvimento

### Volumes Configurados

- Arquivos de configuração (`aws_config.yaml`, `fabric_config.yaml`)
- Certificados e organizações Fabric
- Logs da aplicação
- Cache do Go modules (desenvolvimento)

## 📊 Monitoramento

### Health Checks

- **Endpoint**: `GET /health`
- **Docker Health Check**: Automatizado
- **Intervalo**: 30 segundos
- **Timeout**: 10 segundos

### Logs

```bash
# Logs em tempo real
make logs

# Logs de desenvolvimento
make dev-logs

# Logs para debugging
make debug-logs
```

## 🔧 Troubleshooting

### Problemas Comuns

1. **Erro de conexão AWS**: Verificar `aws_config.yaml`
2. **Erro de conexão Fabric**: Verificar `fabric_config.yaml` e certificados
3. **Porta ocupada**: Alterar porta no `docker-compose.yml`

### Verificação de Configuração

```bash
# Verificar arquivos de configuração
make check-config

# Verificar status dos containers
make health

# Ver logs para debugging
make debug-logs
```

## 📝 Licença

[Especificar licença do projeto]

## 🤝 Contribuição

[Instruções para contribuição]
