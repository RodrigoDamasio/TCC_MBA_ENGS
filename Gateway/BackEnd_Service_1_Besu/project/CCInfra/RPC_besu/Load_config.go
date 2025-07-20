package Infra_rpc

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	entity "project/CCEntitys"

	"gopkg.in/yaml.v2"
)

// RPCConfig estrutura para carregar configurações do arquivo YAML
type RPCConfig struct {
	RPC entity.RPC `yaml:"rpc"`
}

// LoadRPCConfig carrega as configurações do arquivo YAML
func LoadRPCConfig(configPath string) (*RPCConfig, error) {

	if configPath == "" {
		configPath = "rpc_config.yaml"
	}

	absPath, err := filepath.Abs(configPath)
	if err != nil {
		return nil, fmt.Errorf("erro ao resolver caminho do arquivo de configuração: %v", err)
	}

	data, err := ioutil.ReadFile(absPath)
	if err != nil {
		return nil, fmt.Errorf("erro ao ler arquivo de configuração %s: %v", absPath, err)
	}

	var config RPCConfig
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("erro ao fazer parse do arquivo YAML: %v", err)
	}

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

func (c *RPCConfig) GetContractAddress() string {
	return c.RPC.ContractAddress
}

func (c *RPCConfig) GetPrivateKey() string {
	return c.RPC.PrivateKey
}

func (c *RPCConfig) GetChainID() int64 {
	return c.RPC.ChainID
}

func (c *RPCConfig) GetRPCURL() string {
	return c.RPC.RPCURL
}
