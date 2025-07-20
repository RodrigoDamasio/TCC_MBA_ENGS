package handler

import (
	"log"
	entity "project/CCEntitys"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/core/types"
)

// ExtractBatchInfo ajusta dados da trnasação blockchain para serem salvos no dinamo

func ExtractBatchInfo(receipt *types.Receipt, tx *types.Transaction, Besu_Requisition_data entity.CreateBatchRequest) (entity.TransactionData, error) {
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

	//Busca resultado do log

	ID_Besu, err := strconv.Atoi(Besu_Requisition_data.ID)
	if err != nil {
		return entity.TransactionData{}, err
	}

	// PREPARAÇÃO DOS DADOS PARA DYNAMODB
	DinamoPopulate := entity.TransactionData{
		ID:              ID_Besu,
		ProductionOrder: Besu_Requisition_data.ProductionOrder,
		FinalResult:     statusLog,
		Hash:            tx.Hash().Hex(),
		Data_daytime:    Besu_Requisition_data.TestDate,
	}

	return DinamoPopulate, nil
}

func stringContains(s, substr string) bool {
	return strings.Contains(strings.ToUpper(s), strings.ToUpper(substr))
}
