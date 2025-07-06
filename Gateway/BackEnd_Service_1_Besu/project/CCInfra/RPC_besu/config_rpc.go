package Infra_rpc

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

// RPCConfig estrutura para carregar configurações do arquivo YAML
type RPCConfig struct {
	RPC struct {
		ContractAddress string `yaml:"contract_address"`
		PrivateKey      string `yaml:"private_key"`
		ChainID         int64  `yaml:"chain_id"`
		RPCURL          string `yaml:"rpc_url"`
	} `yaml:"rpc"`
}

// LoadRPCConfig carrega as configurações do arquivo YAML
func LoadRPCConfig(configPath string) (*RPCConfig, error) {
	// Se não especificar caminho, usa o padrão na mesma pasta
	if configPath == "" {
		configPath = "rpc_config.yaml"
	}

	// Converte para caminho absoluto se necessário
	absPath, err := filepath.Abs(configPath)
	if err != nil {
		return nil, fmt.Errorf("erro ao resolver caminho do arquivo de configuração: %v", err)
	}

	// Lê o arquivo YAML
	data, err := ioutil.ReadFile(absPath)
	if err != nil {
		return nil, fmt.Errorf("erro ao ler arquivo de configuração %s: %v", absPath, err)
	}

	// Faz o parse do YAML
	var config RPCConfig
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("erro ao fazer parse do arquivo YAML: %v", err)
	}

	// Valida se as configurações necessárias foram carregadas
	if config.RPC.ContractAddress == "" {
		return nil, fmt.Errorf("contract_address não pode estar vazio")
	}
	if config.RPC.PrivateKey == "" {
		return nil, fmt.Errorf("private_key não pode estar vazio")
	}
	if config.RPC.RPCURL == "" {
		return nil, fmt.Errorf("rpc_url não pode estar vazio")
	}
	if config.RPC.ChainID == 0 {
		return nil, fmt.Errorf("chain_id deve ser maior que 0")
	}

	return &config, nil
}

// GetContractAddress retorna o endereço do contrato
func (c *RPCConfig) GetContractAddress() string {
	return c.RPC.ContractAddress
}

// GetPrivateKey retorna a chave privada
func (c *RPCConfig) GetPrivateKey() string {
	return c.RPC.PrivateKey
}

// GetChainID retorna o ID da chain
func (c *RPCConfig) GetChainID() int64 {
	return c.RPC.ChainID
}

// GetRPCURL retorna a URL do RPC
func (c *RPCConfig) GetRPCURL() string {
	return c.RPC.RPCURL
}
