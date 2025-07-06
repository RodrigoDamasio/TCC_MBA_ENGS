# Configuração RPC - Blockchain

Este módulo fornece funcionalidades para carregar configurações de blockchain a partir de arquivos YAML.

## Arquivos

### `rpc_config.yaml`
Arquivo de configuração contendo:
- `contract_address`: Endereço do contrato inteligente
- `private_key`: Chave privada para transações
- `chain_id`: ID da blockchain
- `rpc_url`: URL do node RPC

### `config_rpc.go`
Contém a estrutura `RPCConfig` e funções para carregar configurações do arquivo YAML.

### `rpc_config.go`
Contém a estrutura `Blockchain` e funções para criar instâncias de conexão com a blockchain.

## Como Usar

### 1. Carregando configuração padrão:
```go
blockchain, err := NewBlockchain()
if err != nil {
    log.Fatal(err)
}
```

### 2. Carregando configuração de arquivo específico:
```go
blockchain, err := NewBlockchainWithConfig("caminho/para/config.yaml")
if err != nil {
    log.Fatal(err)
}
```

### 3. Carregando apenas as configurações:
```go
config, err := LoadRPCConfig("rpc_config.yaml")
if err != nil {
    log.Fatal(err)
}

contractAddr := config.GetContractAddress()
chainID := config.GetChainID()
rpcURL := config.GetRPCURL()
privateKey := config.GetPrivateKey()
```

## Estrutura do arquivo YAML

```yaml
rpc:
  contract_address: "0x16fb14f423821b20c5da2b14a6b2cfbb061b71cf"
  private_key: "0xf8060dd8211fd80d62023e74aff54d870e13835c630865d6250761394efddc98"
  chain_id: 1337
  rpc_url: "http://localhost:8550"
```

## Funcionalidades

- **Validação**: O sistema valida se todas as configurações obrigatórias estão presentes
- **Flexibilidade**: Permite usar diferentes arquivos de configuração
- **Segurança**: As configurações sensíveis ficam separadas do código
- **Facilidade**: Interface simples para acessar as configurações

## Dependências

- `gopkg.in/yaml.v2`: Para parsing do arquivo YAML
- `github.com/ethereum/go-ethereum`: Para interação com blockchain Ethereum
