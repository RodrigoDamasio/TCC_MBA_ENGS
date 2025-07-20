package handler

import (
	"encoding/json"
	"net/http"
	db "project/CCInfra/DB"
	grpc "project/CCInfra/Grpc_Hyperledger"

	"github.com/gin-gonic/gin"

	"github.com/hyperledger/fabric-gateway/pkg/client"
)

type BatchControllerInterface interface {
	CreateBatch(c *gin.Context)
	QueryBatchByID(c *gin.Context)
}

type BatchController struct {
}

func NewBatchController() BatchControllerInterface {
	return &BatchController{}
}

// CreateBatch godoc
// @Summary Cria um novo lote de teste de qualidade
// @Description Envia uma transação para criar um novo lote de teste de qualidade no Hyperledger Fabric.
// @Tags batch
// @Accept json
// @Produce json
// @Param batch body map[string]interface{} true "Dados do lote de teste de qualidade"
// @Success 200 {string} string "Lote de teste de qualidade criado com sucesso"
// @Failure 500 {string} string "Erro ao criar lote de teste de qualidade"
// @Router /batch [post]
func (bc *BatchController) CreateBatch(c *gin.Context) {

	var data map[string]interface{}

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	//Envia transação para o Hyperledger Fabric

	TransactionFields, commit, err := grpc.ContractPointer.SubmitAsync(
		"PerformBatchQualityTest",
		client.WithArguments(
			data["id"].(string),
			data["minLimit"].(string),
			data["maxLimit"].(string),
			data["values"].(string),
			data["productionOrder"].(string),
			data["testDate"].(string),
			data["testDescription"].(string),
			data["tokenURI"].(string),
		),
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar lote de teste de qualidade"})
		return
	}

	// Aguarda a confirmação do commit da transação no ledger (commit definitivo)

	status, err := commit.Status()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao aguardar confirmação do commit"})
		return
	}
	if !status.Successful {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Transação não foi confirmada no ledger"})
		return
	}

	// Organiza dados a serem salvos no DynamoDB

	DinamoPopulate, err := ExtractBatchInfo(string(TransactionFields), commit.TransactionID())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao processar informações do lote para o DynamoDB"})
		return
	}

	// Salva os dados no DynamoDB APÓS commit confirmado

	db.SaveTransaction(*DinamoPopulate)

	c.JSON(http.StatusOK, gin.H{
		"message": "Lote de teste de qualidade criado com sucesso e confirmado no ledger",
		"Dados": gin.H{
			"ID":              DinamoPopulate.ID,
			"ProductionOrder": DinamoPopulate.ProductionOrder,
			"FinalResult":     DinamoPopulate.FinalResult,
			"Data_daytime":    DinamoPopulate.Data_daytime,
		},
		"hash": DinamoPopulate.Hash,
		"data": DinamoPopulate.ProductionOrder,
	})
}

// QueryBatchByID godoc
// @Summary Consulta um lote de teste de qualidade
// @Description Consulta os detalhes de um lote de teste de qualidade pelo ID.
// @Tags batch
// @Accept json
// @Produce json
// @Param id path string true "ID do lote de teste"
// @Success 200 {object} map[string]interface{} "Dados do lote de teste de qualidade"
// @Failure 404 {string} string "Lote não encontrado"
// @Failure 500 {string} string "Erro ao consultar lote"
// @Router /batch/{id} [get]
func (bc *BatchController) QueryBatchByID(c *gin.Context) {

	id := c.Param("id")

	evaluateResult, err := grpc.ContractPointer.EvaluateTransaction("QueryQualityTestBatch", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao consultar lote de teste de qualidade"})
		return
	}

	var result map[string]interface{}
	if err := json.Unmarshal(evaluateResult, &result); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao processar dados do lote"})
		return
	}

	c.JSON(http.StatusOK, result)
}

// Backward compatibility wrapper functions
// These can be removed once all code is updated to use the controller

// CreateBatch is a wrapper function for backward compatibility
func CreateBatch(c *gin.Context) {
	controller := NewBatchController()
	controller.CreateBatch(c)
}

// QueryBatchByID is a wrapper function for backward compatibility
func QueryBatchByID(c *gin.Context) {
	controller := NewBatchController()
	controller.QueryBatchByID(c)
}
