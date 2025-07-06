package handler

import (
	"encoding/json"
	"net/http"
	entity "project/CCEntitys"
	db "project/CCInfra/DB"
	grpc "project/CCInfra/Grpc_Hyperledger"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/hyperledger/fabric-gateway/pkg/client"
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

	contract := grpc.GetContract()

	TransactionFields, commit, err := contract.SubmitAsync(
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

	ExtractBatchInfo, err := ExtractBatchInfo(string(TransactionFields))

	ID_dinamo, err := strconv.Atoi(ExtractBatchInfo.ID)

	// Salva os dados no DynamoDB APÓS commit confirmado
	DinamoPopulate := entity.TransactionData{
		ID:              ID_dinamo,
		ProductionOrder: ExtractBatchInfo.ProductionOrder,
		FinalResult:     ExtractBatchInfo.FinalResult,
		Hash:            commit.TransactionID(),
		Data_daytime:    ExtractBatchInfo.Data_daytime,
	}

	db.SaveTransaction(DinamoPopulate)

	// Não sobrescreve o erro anterior
	// (erro do DynamoDB deve ser tratado individualmente, mas a transação já está confirmada no ledger)

	c.JSON(http.StatusOK, gin.H{
		"message": "Lote de teste de qualidade criado com sucesso e confirmado no ledger",
		"Dados":   ExtractBatchInfo,
		"hash":    DinamoPopulate.Hash,
		"data":    DinamoPopulate.ProductionOrder,
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

func QueryBatchByID(c *gin.Context) {

	id := c.Param("id")

	contract := grpc.GetContract()
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
