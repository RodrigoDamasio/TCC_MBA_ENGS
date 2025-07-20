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

var ClientPointer *dynamodb.Client
var TablePointer *string

func init() {

	configManager := NewConfigManager()

	// Carrega a configuração do arquivo YAML
	err := configManager.LoadAWSConfig()
	if err != nil {
		log.Fatalf("Erro ao carregar configuração do arquivo YAML: %v", err)
	}

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("Erro ao carregar configuração da AWS: %v", err)
	}

	// Novo cliente DynamoDB
	client := dynamodb.NewFromConfig(cfg)

	tableName := configManager.GetTableName()

	ClientPointer = client
	TablePointer = &tableName

	log.Printf("Conexão com DynamoDB estabelecida com sucesso. Tabela: %s", tableName)

}

func SaveTransaction(transaction entity.TransactionData) error {

	// Prepara os dados para inserção
	item := map[string]types.AttributeValue{
		"ID_Besu":         &types.AttributeValueMemberN{Value: strconv.FormatInt(int64(transaction.ID), 10)},
		"ProductionOrder": &types.AttributeValueMemberS{Value: transaction.ProductionOrder},
		"FinalResult":     &types.AttributeValueMemberS{Value: transaction.FinalResult},
		"Hash":            &types.AttributeValueMemberS{Value: transaction.Hash},
		"Data_daytime":    &types.AttributeValueMemberS{Value: transaction.Data_daytime},
	}

	// Insere no DynamoDB
	_, err := ClientPointer.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(*TablePointer),
		Item:      item,
	})

	if err != nil {
		log.Printf("Erro ao salvar transação no DynamoDB: %v", err)
		return err
	}

	log.Printf("Transação salva com sucesso no DynamoDB: %+v", transaction)
	return nil
}
