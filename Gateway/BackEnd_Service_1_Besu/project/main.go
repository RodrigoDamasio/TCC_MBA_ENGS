package main

import (
	"log"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "project/docs" // Geração de documentação Swagger
	"project/handler"
)

// @title Quality Test API
// @version 1.0
// @description API para criar e consultar lotes de testes de qualidade usando contratos Solidity.
// @host localhost:8081
// @BasePath /
func main() {
	router := gin.Default()

	// Rotas
	router.POST("/batch", handler.CreateBatch)
	router.GET("/batch/:id", handler.QueryBatchByID)
	router.GET("/ws/latest-block", handler.LatestBlockHandler)

	// Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Println("Iniciando o servidor na porta 8082...")
	if err := router.Run(":8082"); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
