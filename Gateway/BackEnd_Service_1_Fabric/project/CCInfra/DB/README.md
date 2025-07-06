# Configuração AWS

Este diretório contém os arquivos necessários para configurar automaticamente as variáveis de ambiente AWS quando você executar a aplicação.

## Arquivos

- `aws_config.yaml` - Arquivo de configuração com as credenciais e configurações AWS
- `aws_config.go` - Código Go para carregar as configurações e definir variáveis de ambiente
- `setup_aws_and_run.sh` - Script bash para configurar ambiente e executar a aplicação

## Como configurar

### 1. Editar o arquivo de configuração

Abra o arquivo `aws_config.yaml` e preencha com suas credenciais AWS:

```yaml
aws:
  credentials:
    access_key_id: "sua_access_key_aqui"
    secret_access_key: "sua_secret_key_aqui"
  
  config:
    region: "us-east-1"  # Sua região preferida
    output: "json"
  
  dynamodb:
    endpoint: ""  # Deixe vazio para usar o endpoint padrão
    table_prefix: "tcc_"
    
  session:
    token: ""  # Para sessões temporárias (opcional)
    profile: "default"
```

### 2. Como executar a aplicação

#### Opção 1: Usando o script (Recomendado)
```bash
# Executar a aplicação
./CCInfra/DB/setup_aws_and_run.sh

# Compilar apenas
./CCInfra/DB/setup_aws_and_run.sh build

# Executar testes
./CCInfra/DB/setup_aws_and_run.sh test
```

#### Opção 2: Go run normal
```bash
# A aplicação irá carregar automaticamente as configurações AWS
go run main.go
```

## Configurações automáticas

Quando você executar a aplicação, as seguintes variáveis de ambiente serão configuradas automaticamente:

- `AWS_ACCESS_KEY_ID`
- `AWS_SECRET_ACCESS_KEY`
- `AWS_DEFAULT_REGION`
- `AWS_REGION`
- `AWS_DEFAULT_OUTPUT`
- `AWS_PROFILE`
- `AWS_SESSION_TOKEN` (se fornecido)

## Segurança

⚠️ **IMPORTANTE**: 
- Nunca commite o arquivo `aws_config.yaml` com suas credenciais reais no Git
- Adicione o arquivo ao `.gitignore`
- Use IAM roles em produção sempre que possível
- Para desenvolvimento local, considere usar `aws configure` ao invés de hardcodar credenciais

## Exemplo de .gitignore

Adicione estas linhas ao seu `.gitignore`:

```
# Configurações AWS sensíveis
CCInfra/DB/aws_config.yaml
```

## Troubleshooting

### Erro: "Arquivo de configuração AWS não encontrado"
- Verifique se o arquivo `aws_config.yaml` existe na pasta `CCInfra/DB/`
- Verifique se o caminho está correto

### Erro: "Credenciais AWS inválidas"
- Verifique se `access_key_id` e `secret_access_key` estão corretos
- Verifique se a região está correta
- Teste as credenciais com `aws configure list`

### Script não executa
- Verifique se o script tem permissão de execução: `chmod +x CCInfra/DB/setup_aws_and_run.sh`
- Execute do diretório raiz do projeto
