package main

import (
	"log"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files" // Import correto para o Swagger UI handler
	ginSwagger "github.com/swaggo/gin-swagger"

	// Swagger docs gerados (substitua por onde estão seus docs gerados)
	_ "project/docs" // Certifique-se de gerar a documentação antes de rodar
	"project/handler"
)

// @title Quality Test API
// @version 1.0
// @description API para criar e consultar lotes de testes de qualidade usando Hyperledger Fabric.
// @host localhost:8081
// @BasePath /
func main() {
	router := gin.Default()

	// Rotas
	router.POST("/batch", handler.CreateBatch)
	router.GET("/batch/:id", handler.QueryBatchByID)

	// Documentação Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Println("Iniciando o servidor na porta 8081...")
	if err := router.Run(":8081"); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
