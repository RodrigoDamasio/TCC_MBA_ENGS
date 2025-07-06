package Infra_grpc

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type FabricConfig struct {
	Fabric struct {
		MSPID         string `yaml:"msp_id"`
		CryptoPath    string `yaml:"crypto_path"`
		PeerEndpoint  string `yaml:"peer_endpoint"`
		GatewayPeer   string `yaml:"gateway_peer"`
		ChannelName   string `yaml:"channel_name"`
		ChaincodeName string `yaml:"chaincode_name"`
		Paths         struct {
			CertPath    string `yaml:"cert_path"`
			KeyPath     string `yaml:"key_path"`
			TLSCertPath string `yaml:"tls_cert_path"`
		} `yaml:"paths"`
	} `yaml:"fabric"`
}

var Config *FabricConfig

// LoadConfig carrega as configurações do arquivo YAML
func LoadConfig(configPath string) error {
	if configPath == "" {
		configPath = "./CCInfra/Grpc_Hyperledger/fabric_config.yaml"
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		return fmt.Errorf("erro ao ler arquivo de configuração: %w", err)
	}

	Config = &FabricConfig{}
	err = yaml.Unmarshal(data, Config)
	if err != nil {
		return fmt.Errorf("erro ao fazer parse do arquivo de configuração: %w", err)
	}

	return nil
}

// GetMSPID retorna o MSP ID
func (c *FabricConfig) GetMSPID() string {
	return c.Fabric.MSPID
}

// GetCryptoPath retorna o caminho base dos certificados
func (c *FabricConfig) GetCryptoPath() string {
	return c.Fabric.CryptoPath
}

// GetCertPath retorna o caminho completo dos certificados
func (c *FabricConfig) GetCertPath() string {
	return filepath.Join(c.Fabric.CryptoPath, c.Fabric.Paths.CertPath)
}

// GetKeyPath retorna o caminho completo das chaves privadas
func (c *FabricConfig) GetKeyPath() string {
	return filepath.Join(c.Fabric.CryptoPath, c.Fabric.Paths.KeyPath)
}

// GetTLSCertPath retorna o caminho completo do certificado TLS
func (c *FabricConfig) GetTLSCertPath() string {
	return filepath.Join(c.Fabric.CryptoPath, c.Fabric.Paths.TLSCertPath)
}

// GetPeerEndpoint retorna o endpoint do peer
func (c *FabricConfig) GetPeerEndpoint() string {
	return c.Fabric.PeerEndpoint
}

// GetGatewayPeer retorna o nome do gateway peer
func (c *FabricConfig) GetGatewayPeer() string {
	return c.Fabric.GatewayPeer
}

// GetChannelName retorna o nome do canal
func (c *FabricConfig) GetChannelName() string {
	return c.Fabric.ChannelName
}

// GetChaincodeName retorna o nome do chaincode
func (c *FabricConfig) GetChaincodeName() string {
	return c.Fabric.ChaincodeName
}
