package Infra_db

import (
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

// AWSConfig representa a estrutura do arquivo aws_config.yaml
type AWSConfig struct {
	AWS struct {
		Credentials struct {
			AccessKeyID     string `yaml:"access_key_id"`
			SecretAccessKey string `yaml:"secret_access_key"`
		} `yaml:"credentials"`
		Config struct {
			Region string `yaml:"region"`
			Output string `yaml:"output"`
		} `yaml:"config"`
		DynamoDB struct {
			Endpoint    string `yaml:"endpoint"`
			TablePrefix string `yaml:"table_prefix"`
		} `yaml:"dynamodb"`
		Session struct {
			Token   string `yaml:"token"`
			Profile string `yaml:"profile"`
		} `yaml:"session"`
	} `yaml:"aws"`
}

// ConfigManager gerencia as configurações AWS
type ConfigManager struct {
	config *AWSConfig
}

// NewConfigManager cria uma nova instância do ConfigManager
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

// GetTableName retorna o nome da tabela DynamoDB configurada
func (cm *ConfigManager) GetTableName() string {
	if cm.config == nil {
		log.Printf("Configuração não carregada. Usando valor padrão.")
		return "FabricTransactionHashTCC"
	}
	return cm.config.AWS.DynamoDB.TablePrefix
}

// GetRegion retorna a região AWS configurada
func (cm *ConfigManager) GetRegion() string {
	if cm.config == nil {
		log.Printf("Configuração não carregada. Usando valor padrão.")
		return "us-east-1"
	}
	return cm.config.AWS.Config.Region
}

// GetDynamoDBEndpoint retorna o endpoint do DynamoDB configurado
func (cm *ConfigManager) GetDynamoDBEndpoint() string {
	if cm.config == nil {
		return ""
	}
	return cm.config.AWS.DynamoDB.Endpoint
}

// GetProfile retorna o perfil AWS configurado
func (cm *ConfigManager) GetProfile() string {
	if cm.config == nil {
		return "default"
	}
	return cm.config.AWS.Session.Profile
}

// GetConfig retorna a configuração completa (caso necessário)
func (cm *ConfigManager) GetConfig() *AWSConfig {
	return cm.config
}
