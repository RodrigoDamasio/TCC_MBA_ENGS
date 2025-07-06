package handler

import (
	"encoding/json"
	entity "project/CCEntitys"
)

// ExtractBatchInfo recebe uma string JSON e retorna a struct preenchida com os campos relevantes
func ExtractBatchInfo(jsonStr string) (*entity.TransactionHyperledger_BatchInfo, error) {

	var info entity.TransactionHyperledger_BatchInfo
	if err := json.Unmarshal([]byte(jsonStr), &info); err != nil {
		return nil, err
	}

	return &info, nil

}
