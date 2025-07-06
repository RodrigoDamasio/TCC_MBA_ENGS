package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files" // Import correto para o Swagger UI handler
	ginSwagger "github.com/swaggo/gin-swagger"

	// Swagger docs gerados (substitua por onde estão seus docs gerados)
	_ "project/CCDocs" // Certifique-se de gerar a documentação antes de rodar
	handler "project/CCHandler"
)

// @title Quality Test API
// @version 1.0
// @description API para criar e consultar lotes de testes de qualidade usando Hyperledger Fabric.
// @host localhost:8081
// @BasePath /
func main() {

	router := gin.Default()

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "healthy",
			"service": "fabric-gateway",
			"version": "1.0.0",
		})
	})

	// Rotas principais
	router.POST("/batch", handler.CreateBatch)
	router.GET("/batch/:id", handler.QueryBatchByID)

	// Documentação Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Obter porta das variáveis de ambiente
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081" // porta padrão
	}

	log.Printf("Iniciando o servidor na porta %s...", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
