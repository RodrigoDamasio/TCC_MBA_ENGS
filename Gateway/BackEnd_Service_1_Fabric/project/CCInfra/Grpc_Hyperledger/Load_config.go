package Infra_grpc

import (
	"fmt"
	"os"
	"path/filepath"
	entitys "project/CCEntitys"

	"gopkg.in/yaml.v2"
)

type FabricConfig struct {
	Fabric entitys.Fabric `yaml:"fabric"`
}

var Config *FabricConfig

// Carrega arquivo YAML
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

func (c *FabricConfig) GetMSPID() string {
	return c.Fabric.MSPID
}

func (c *FabricConfig) GetCryptoPath() string {
	return c.Fabric.CryptoPath
}

func (c *FabricConfig) GetCertPath() string {
	return filepath.Join(c.Fabric.CryptoPath, c.Fabric.Paths.CertPath)
}

func (c *FabricConfig) GetKeyPath() string {
	return filepath.Join(c.Fabric.CryptoPath, c.Fabric.Paths.KeyPath)
}

func (c *FabricConfig) GetTLSCertPath() string {
	return filepath.Join(c.Fabric.CryptoPath, c.Fabric.Paths.TLSCertPath)
}

func (c *FabricConfig) GetPeerEndpoint() string {
	return c.Fabric.PeerEndpoint
}

func (c *FabricConfig) GetGatewayPeer() string {
	return c.Fabric.GatewayPeer
}

func (c *FabricConfig) GetChannelName() string {
	return c.Fabric.ChannelName
}

func (c *FabricConfig) GetChaincodeName() string {
	return c.Fabric.ChaincodeName
}
