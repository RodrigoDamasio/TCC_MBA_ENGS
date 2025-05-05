package handler

import (
	"math/big"
	"net/http"

	"project/contract"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
)

// Endereço do contrato e configuração
const contractAddress = "0x370ca606b26d87a894f61395d0eefddb7161a9cb"
const privateKeyHex = "0xf8060dd8211fd80d62023e74aff54d870e13835c630865d6250761394efddc98"

type CreateBatchRequest struct {
	ID              string  `json:"id"`
	MinLimit        int64   `json:"minLimit"`
	MaxLimit        int64   `json:"maxLimit"`
	Values          []int64 `json:"values"`
	ProductionOrder string  `json:"productionOrder"`
	TestDate        string  `json:"testDate"`
	TestDescription string  `json:"testDescription"`
}

// CreateBatch godoc
// @Summary Cria um novo lote de testes de qualidade
// @Description Envia uma transação para criar um novo lote de testes de qualidade no contrato Solidity
// @Tags Batch
// @Accept json
// @Produce json
// @Param batch body CreateBatchRequest true "Dados do lote"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /batch [post]
func CreateBatch(c *gin.Context) {
	var req CreateBatchRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Conectando ao nó Ethereum
	client, err := ethclient.Dial("http://localhost:8550")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to Ethereum client"})
		return
	}

	// Criação do autenticador a partir da chave privada
	privateKey, err := crypto.HexToECDSA(privateKeyHex[2:]) // Remove o prefixo "0x"
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid private key"})
		return
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1337)) // Ajuste a ChainID conforme necessário
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create authenticated transaction"})
		return
	}

	// Instanciando o contrato
	contractInstance, err := contract.NewContract(common.HexToAddress(contractAddress), client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load contract"})
		return
	}

	// Convertendo os valores para big.Int
	values := make([]*big.Int, len(req.Values))
	for i, v := range req.Values {
		values[i] = big.NewInt(v)
	}

	// Enviando a transação
	tx, err := contractInstance.PerformBatchQualityTest(auth, req.ID, big.NewInt(req.MinLimit), big.NewInt(req.MaxLimit), values, req.ProductionOrder, req.TestDate, req.TestDescription)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to submit transaction"})
		return
	}

	// Retornando o hash da transação
	c.JSON(http.StatusOK, gin.H{
		"transactionHash": tx.Hash().Hex(),
	})
}
