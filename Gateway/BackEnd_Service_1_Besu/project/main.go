package main

import (
	"log"
	"net/http"
	"os"
	_ "project/CCDocs" // Geração de documentação Swagger
	handler "project/CCHandler"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Quality Test API
// @version 1.0
// @description API para criar e consultar lotes de testes de qualidade usando contratos Solidity.
// @host localhost:8082
// @BasePath /

func main() {
	router := gin.Default()

	// Initialize the batch controller
	batchController := handler.NewBatchController()

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "healthy",
			"service": "besu-gateway",
			"version": "1.0.0",
		})
	})

	// Rotas principais - using controller methods
	router.POST("/batch", batchController.CreateBatch)
	router.GET("/batch/:id", batchController.QueryBatchByID)
	router.GET("/batch/tx/:txHash", batchController.QueryBatchByTransaction)

	// Documentação Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Obter porta das variáveis de ambiente
	port := os.Getenv("PORT")
	if port == "" {
		port = "8082" // porta padrão
	}

	log.Printf("Iniciando o servidor na porta %s...", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
