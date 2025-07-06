#!/bin/bash

# Script para configurar variáveis de ambiente AWS e executar a aplicação Go
# Arquivo: setup_aws_and_run.sh

# Cores para output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

echo -e "${YELLOW}🔧 Configurando ambiente AWS para a aplicação...${NC}"

# Verifica se o arquivo de configuração existe
AWS_CONFIG_FILE="./CCInfra/DB/aws_config.yaml"

if [ ! -f "$AWS_CONFIG_FILE" ]; then
    echo -e "${RED}❌ Arquivo de configuração AWS não encontrado: $AWS_CONFIG_FILE${NC}"
    echo -e "${YELLOW}📝 Criando arquivo de exemplo...${NC}"
    echo "Por favor, edite o arquivo $AWS_CONFIG_FILE com suas credenciais AWS"
    exit 1
fi

echo -e "${GREEN}✅ Arquivo de configuração AWS encontrado${NC}"

# Função para ler valores do arquivo YAML (versão simplificada)
get_yaml_value() {
    local key=$1
    local file=$2
    grep "$key:" "$file" | sed 's/.*: *"\?\([^"]*\)"\?.*/\1/' | head -1
}

# Lê as configurações do arquivo YAML
echo -e "${YELLOW}📖 Carregando configurações do arquivo YAML...${NC}"

export AWS_ACCESS_KEY_ID=$(get_yaml_value "access_key_id" "$AWS_CONFIG_FILE")
export AWS_SECRET_ACCESS_KEY=$(get_yaml_value "secret_access_key" "$AWS_CONFIG_FILE")
export AWS_REGION=$(get_yaml_value "region" "$AWS_CONFIG_FILE")

# Exporta as variáveis de ambiente se não estiverem vazias
if [ ! -z "$AWS_ACCESS_KEY_ID" ] && [ "$AWS_ACCESS_KEY_ID" != '""' ]; then
    export AWS_ACCESS_KEY_ID
    echo -e "${GREEN}✅ AWS_ACCESS_KEY_ID configurado${NC}"
fi

if [ ! -z "$AWS_SECRET_ACCESS_KEY" ] && [ "$AWS_SECRET_ACCESS_KEY" != '""' ]; then
    export AWS_SECRET_ACCESS_KEY
    echo -e "${GREEN}✅ AWS_SECRET_ACCESS_KEY configurado${NC}"
fi

if [ ! -z "$AWS_REGION" ] && [ "$AWS_REGION" != '""' ]; then
    export AWS_DEFAULT_REGION="$AWS_REGION"
    export AWS_REGION
    echo -e "${GREEN}✅ AWS_REGION configurado para: $AWS_REGION${NC}"
fi

if [ ! -z "$AWS_OUTPUT" ] && [ "$AWS_OUTPUT" != '""' ]; then
    export AWS_DEFAULT_OUTPUT="$AWS_OUTPUT"
    echo -e "${GREEN}✅ AWS_DEFAULT_OUTPUT configurado para: $AWS_OUTPUT${NC}"
fi

if [ ! -z "$AWS_PROFILE" ] && [ "$AWS_PROFILE" != '""' ]; then
    export AWS_PROFILE
    echo -e "${GREEN}✅ AWS_PROFILE configurado para: $AWS_PROFILE${NC}"
fi

if [ ! -z "$AWS_SESSION_TOKEN" ] && [ "$AWS_SESSION_TOKEN" != '""' ]; then
    export AWS_SESSION_TOKEN
    echo -e "${GREEN}✅ AWS_SESSION_TOKEN configurado${NC}"
fi

echo -e "${GREEN}🎯 Todas as variáveis de ambiente AWS foram configuradas!${NC}"
echo -e "${YELLOW}🚀 Iniciando a aplicação Go...${NC}"

# Executa a aplicação Go
if [ "$1" = "build" ]; then
    echo -e "${YELLOW}🔨 Compilando a aplicação...${NC}"
    go build -o main .
    echo -e "${YELLOW}▶️  Executando o binário compilado...${NC}"
    ./main
elif [ "$1" = "test" ]; then
    echo -e "${YELLOW}🧪 Executando testes...${NC}"
    go test ./...
elif [ "$1" = "exec" ]; then
    echo -e "${YELLOW}▶️  Executando binário pré-compilado...${NC}"
    ./main
else
    echo -e "${YELLOW}▶️  Executando go run main.go...${NC}"
    go run main.go
fi

# Captura o código de saída
exit_code=$?

if [ $exit_code -eq 0 ]; then
    echo -e "${GREEN}✅ Aplicação executada com sucesso!${NC}"
else
    echo -e "${RED}❌ Aplicação terminou com erro (código: $exit_code)${NC}"
fi

exit $exit_code
