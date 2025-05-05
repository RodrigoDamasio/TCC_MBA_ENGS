package handler

import (
	"net/http"

	"project/db"

	"github.com/gin-gonic/gin"
)

func CreateTransactionDinamo(c *gin.Context) {
	// Exemplo de dados recebidos da requisição
	transaction := db.TransactionData{
		ID:              "unique-transaction-id-001",
		ProductionOrder: "PO-001",
		FinalResult:     "Aprovado",
	}

	// Salva os dados no DynamoDB
	err := db.SaveTransaction(transaction)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao salvar transação no DynamoDB"})
		return
	}

	// Retorna o sucesso
	c.JSON(http.StatusOK, gin.H{
		"message": "Transação criada com sucesso",
		"data":    transaction,
	})
}
