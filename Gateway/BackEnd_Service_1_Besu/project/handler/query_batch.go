package handler

import (
	"net/http"

	"project/contract"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
)

// QueryBatchByID godoc
// @Summary Consulta os detalhes de um lote de testes de qualidade
// @Description Faz uma consulta ao contrato Solidity para obter os detalhes de um lote pelo ID
// @Tags Batch
// @Accept json
// @Produce json
// @Param id path string true "ID do lote"
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]string
// @Router /batch/{id} [get]
func QueryBatchByID(c *gin.Context) {
	id := c.Param("id")

	// Conectando ao nó Ethereum
	client, err := ethclient.Dial("http://localhost:8550")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to Ethereum client"})
		return
	}

	// Instanciando o contrato
	contractInstance, err := contract.NewContract(common.HexToAddress(contractAddress), client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load contract"})
		return
	}

	// Realizando a chamada de consulta
	batchID, minLimit, maxLimit, values, productionOrder, testDate, testDescription, sampleResults, finalResult, err := contractInstance.QueryQualityTestBatch(&bind.CallOpts{}, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query contract"})
		return
	}

	// Retornando os dados do lote
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
	})
}
