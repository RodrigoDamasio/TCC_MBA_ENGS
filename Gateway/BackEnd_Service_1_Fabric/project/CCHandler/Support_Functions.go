package handler

import (
	"encoding/json"
	entity "project/CCEntitys"
	"strconv"
)

// Retorna TransactionData enviado ao Hyperledger pronto para DynamoDB
func ExtractBatchInfo(jsonStr string, transactionHash string) (*entity.TransactionData, error) {

	var batchInfo entity.TransactionHyperledger_BatchInfo
	if err := json.Unmarshal([]byte(jsonStr), &batchInfo); err != nil {
		return nil, err
	}

	idDinamo, err := strconv.Atoi(batchInfo.ID)
	if err != nil {
		return nil, err
	}

	transactionData := &entity.TransactionData{
		ID:              idDinamo,
		ProductionOrder: batchInfo.ProductionOrder,
		FinalResult:     batchInfo.FinalResult,
		Hash:            transactionHash,
		Data_daytime:    batchInfo.Data_daytime,
	}

	return transactionData, nil
}
