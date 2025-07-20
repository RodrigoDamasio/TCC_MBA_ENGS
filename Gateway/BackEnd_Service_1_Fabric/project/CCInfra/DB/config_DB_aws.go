package Infra_db

import (
	"log"
	"os"
	"path/filepath"
	entitys "project/CCEntitys"

	"gopkg.in/yaml.v2"
)

// Estrutura do arquivo aws_config.yaml
type AWSConfig struct {
	AWS entitys.AWS `yaml:"aws"`
}

type ConfigManager struct {
	config *AWSConfig
}

// Nova instância do ConfigManager
func NewConfigManager() *ConfigManager {
	return &ConfigManager{}
}

// LoadAWSConfig carrega a configuração do arquivo aws_config.yaml
func (cm *ConfigManager) LoadAWSConfig() error {
	configPath := filepath.Join("CCInfra", "DB", "aws_config.yaml")

	yamlFile, err := os.ReadFile(configPath)
	if err != nil {
		log.Printf("Erro ao ler arquivo de configuração: %v", err)
		return err
	}

	var config AWSConfig

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Printf("Erro ao fazer unmarshal do YAML: %v", err)
		return err
	}

	cm.config = &config
	return nil
}

func (cm *ConfigManager) GetTableName() string {
	if cm.config == nil {
		log.Printf("Configuração não carregada. Usando valor padrão.")
		return "FabricTransactionHashTCC"
	}
	return cm.config.AWS.DynamoDB.TablePrefix
}

func (cm *ConfigManager) GetRegion() string {
	if cm.config == nil {
		log.Printf("Configuração não carregada. Usando valor padrão.")
		return "us-east-1"
	}
	return cm.config.AWS.Config.Region
}

func (cm *ConfigManager) GetDynamoDBEndpoint() string {
	if cm.config == nil {
		return ""
	}
	return cm.config.AWS.DynamoDB.Endpoint
}

func (cm *ConfigManager) GetProfile() string {
	if cm.config == nil {
		return "default"
	}
	return cm.config.AWS.Session.Profile
}

func (cm *ConfigManager) GetConfig() *AWSConfig {
	return cm.config
}
