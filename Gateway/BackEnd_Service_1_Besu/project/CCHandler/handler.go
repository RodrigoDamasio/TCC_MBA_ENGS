package handler

import (
	"context"
	"log"
	"math/big"
	"net/http"
	entity "project/CCEntitys"
	db "project/CCInfra/DB"
	infrastructure "project/CCInfra/RPC_besu"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
)

// Handler para criar um batch (usecase + handler)
// Estrutura para requisição de criação de batch

// CreateBatch godoc
// @Summary      Cria um novo lote de teste de qualidade
// @Description  Envia um lote de dados para o contrato inteligente e grava no DynamoDB
// @Tags         batch
// @Accept       json
// @Produce      json
// @Param        batch  body      entity.CreateBatchRequest  true  "Dados do lote de teste"
// @Success      200    {object}  map[string]interface{}
// @Failure      400    {object}  map[string]string
// @Failure      500    {object}  map[string]string
// @Router       /batch [post]
func CreateBatch(c *gin.Context) {

	var Besu_Requisition_data entity.CreateBatchRequest

	if err := c.ShouldBindJSON(&Besu_Requisition_data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	blockchain, err := infrastructure.NewBlockchain()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Infra error: " + err.Error()})
		return
	}

	values := make([]*big.Int, len(Besu_Requisition_data.Values))
	for i, v := range Besu_Requisition_data.Values {
		values[i] = big.NewInt(v)
	}

	tx, err := blockchain.ContractInstance.PerformBatchQualityTest(
		blockchain.Auth, Besu_Requisition_data.ID, big.NewInt(Besu_Requisition_data.MinLimit), big.NewInt(Besu_Requisition_data.MaxLimit), values,
		Besu_Requisition_data.ProductionOrder, Besu_Requisition_data.TestDate, Besu_Requisition_data.TestDescription, Besu_Requisition_data.ImagesJson,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to submit transaction"})
		return
	}

	receipt, err := bind.WaitMined(context.Background(), blockchain.Client, tx)
	if err != nil {
		// Erro ao aguardar mineração
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao aguardar confirmação da transação"})
		return
	}

	log.Printf("Log info: %v", receipt.Logs[1].Data)

	statusLog := ""
	if len(receipt.Logs) > 0 {
		// Converte os dados do log para string ASCII
		logDataASCII := string(receipt.Logs[1].Data)

		if logDataASCII == "" {
			statusLog = "Log vazio"
		} else if contains := !(stringContains(logDataASCII, "N")); contains {
			statusLog = "OK"
		} else if contains := (stringContains(logDataASCII, "N")); contains {
			statusLog = "NOK"
		} else {
			statusLog = "Outro"
		}
		log.Printf("Log info (ASCII): %s", logDataASCII)
		log.Printf("Status detectado: %s", statusLog)
	}

	// ...existing code...

	// Função auxiliar para busca case-insensitive

	//Busca resultado do log

	ID_Besu, err := strconv.Atoi(Besu_Requisition_data.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	// GRAVAÇÃO NO DYNAMODB

	DinamoPopulate := entity.TransactionData{
		ID:              ID_Besu,
		ProductionOrder: Besu_Requisition_data.ProductionOrder,
		FinalResult:     statusLog,
		Hash:            string(tx.Hash().Hex()),
		Data_daytime:    Besu_Requisition_data.TestDate,
	}

	db.SaveTransaction(DinamoPopulate)

	//CONFIRMAÇÃO GRAVAÇÃO NO BESU

	c.JSON(http.StatusOK, gin.H{
		"message": "Lote de teste de qualidade criado com sucesso e confirmado no ledger",
		"Dados":   DinamoPopulate,
		"hash":    DinamoPopulate.Hash,
		"data":    DinamoPopulate.ProductionOrder,
	})

}

// Handler para consultar batch por ID (usecase + handler)

// QueryBatchByID godoc
// @Summary      Consulta lote de teste de qualidade por ID
// @Description  Retorna os dados do lote de teste de qualidade pelo ID informado
// @Tags         batch
// @Produce      json
// @Param        id   path      string  true  "ID do lote"
// @Success      200  {object}  map[string]interface{}
// @Failure      500  {object}  map[string]string
// @Router       /batch/{id} [get]
func QueryBatchByID(c *gin.Context) {
	id := c.Param("id")
	blockchain, err := infrastructure.NewBlockchain()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Infra error: " + err.Error()})
		return
	}

	// Chama o novo método de infraestrutura para criar CallOpts
	callOpts := blockchain.NewCallOpts(c.Request.Context())

	batchID, minLimit, maxLimit, values, productionOrder, testDate, testDescription, sampleResults, finalResult, imagesJson, err :=
		blockchain.ContractInstance.QueryQualityTestBatch(callOpts, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query contract"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"batchID":         batchID,
		"minLimit":        minLimit.String(),
		"maxLimit":        maxLimit.String(),
		"values":          values,
		"productionOrder": productionOrder,
		"testDate":        testDate,
		"testDescription": testDescription,
		"sampleResults":   sampleResults,
		"finalResult":     finalResult,
		"imagesJson":      imagesJson,
	})
}

// QueryBatchByTransaction godoc
// @Summary      Consulta lote de teste de qualidade por hash da transação
// @Description  Retorna os logs e informações do recibo da transação na rede Besu
// @Tags         batch
// @Produce      json
// @Param        txHash   path      string  true  "Hash da transação"
// @Success      200  {object}  map[string]interface{}
// @Failure      404  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /batch/tx/{txHash} [get]
func QueryBatchByTransaction(c *gin.Context) {
	txHash := c.Param("txHash")
	blockchain, err := infrastructure.NewBlockchain()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Infra error: " + err.Error()})
		return
	}

	// Busca o recibo da transação pelo hash
	receipt, err := blockchain.Client.TransactionReceipt(c.Request.Context(), common.HexToHash(txHash))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Recibo da transação não encontrado"})
		return
	}

	// Retorna informações do recibo e logs
	c.JSON(http.StatusOK, gin.H{
		"txHash":  receipt.TxHash.Hex(),
		"status":  receipt.Status,
		"logs":    receipt.Logs,
		"block":   receipt.BlockNumber.String(),
		"gasUsed": receipt.GasUsed,
	})
}

func stringContains(s, substr string) bool {
	return strings.Contains(strings.ToUpper(s), strings.ToUpper(substr))
}
