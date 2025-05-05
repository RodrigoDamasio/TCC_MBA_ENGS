package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

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
func CreateBatch(c *gin.Context) {
	var data map[string]interface{}
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	contract := GetContract()
	_, err := contract.SubmitTransaction(
		"PerformBatchQualityTest",
		data["id"].(string),
		data["minLimit"].(string),
		data["maxLimit"].(string),
		data["values"].(string),
		data["productionOrder"].(string),
		data["testDate"].(string),
		data["testDescription"].(string),
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar lote de teste de qualidade"})
		return
	}

	c.JSON(http.StatusOK, "Lote de teste de qualidade criado com sucesso")

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
func QueryBatchByID(c *gin.Context) {
	id := c.Param("id")

	contract := GetContract()
	evaluateResult, err := contract.EvaluateTransaction("QueryQualityTestBatch", id)
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
