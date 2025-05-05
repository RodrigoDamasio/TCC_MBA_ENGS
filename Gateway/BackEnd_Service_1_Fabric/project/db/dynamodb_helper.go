package db

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type TransactionData struct {
	ID              string
	ProductionOrder string
	FinalResult     string
}

// SaveTransaction salva os dados da transação no DynamoDB
func SaveTransaction(transaction TransactionData) error {
	// Carrega a configuração padrão da AWS
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("Erro ao carregar configuração da AWS: %v", err)
		return err
	}

	// Cria o cliente do DynamoDB
	client := dynamodb.NewFromConfig(cfg)

	// Define o nome da tabela no DynamoDB
	tableName := "FabricTransactionHashTCC"

	// Prepara os dados para inserção
	item := map[string]types.AttributeValue{
		"ID":              &types.AttributeValueMemberS{Value: transaction.ID},
		"ProductionOrder": &types.AttributeValueMemberS{Value: transaction.ProductionOrder},
		"FinalResult":     &types.AttributeValueMemberS{Value: transaction.FinalResult},
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
