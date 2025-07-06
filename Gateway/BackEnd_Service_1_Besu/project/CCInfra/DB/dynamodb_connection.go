package Infra_db

import (
	"context"
	"log"
	entity "project/CCEntitys"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// SaveTransaction salva os dados da transação no DynamoDB
func SaveTransaction(transaction entity.TransactionData) error {
	// Cria uma instância do ConfigManager e carrega a configuração
	configManager := NewConfigManager()
	err := configManager.LoadAWSConfig()
	if err != nil {
		log.Fatalf("Erro ao carregar configuração do arquivo YAML: %v", err)
		return err
	}

	// Carrega a configuração padrão da AWS
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("Erro ao carregar configuração da AWS: %v", err)
		return err
	}

	// Cria o cliente do DynamoDB
	client := dynamodb.NewFromConfig(cfg)

	// Define o nome da tabela no DynamoDB a partir do ConfigManager
	tableName := configManager.GetTableName()

	// Prepara os dados para inserção
	item := map[string]types.AttributeValue{
		"ID_Besu":         &types.AttributeValueMemberN{Value: strconv.FormatInt(int64(transaction.ID), 10)},
		"ProductionOrder": &types.AttributeValueMemberS{Value: transaction.ProductionOrder},
		"FinalResult":     &types.AttributeValueMemberS{Value: transaction.FinalResult},
		"Hash":            &types.AttributeValueMemberS{Value: transaction.Hash},
		"Data_daytime":    &types.AttributeValueMemberS{Value: transaction.Data_daytime},
	}

	// Realiza a operação de inserção no DynamoDB
	_, err = client.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item:      item,
	})

	if err != nil {
		log.Printf("Erro ao salvar transação no DynamoDB: %v", err)
		return err
	}

	log.Printf("Transação salva com sucesso no DynamoDB: %+v", transaction)
	return nil
}
