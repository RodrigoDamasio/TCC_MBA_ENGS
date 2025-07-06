package entitys

type TransactionHyperledger_BatchInfo struct {
	ID              string `json:"id"`
	ProductionOrder string `json:"production_order"`
	FinalResult     string `json:"final_result"`
	Data_daytime    string `json:"test_Date"`
}

type TransactionData struct {
	ID              int
	ProductionOrder string
	FinalResult     string
	Hash            string
	Data_daytime    string
}

type CreateBatchRequest struct {
	ID              string  `json:"id"`
	MinLimit        int64   `json:"minLimit"`
	MaxLimit        int64   `json:"maxLimit"`
	Values          []int64 `json:"values"`
	ProductionOrder string  `json:"productionOrder"`
	TestDate        string  `json:"testDate"`
	TestDescription string  `json:"testDescription"`
	ImagesJson      string  `json:"imagesJson"`
}

type Transactioninfo struct {
	txHash1 string `json:"txHash"`
	Status  string `json:"status"`
	Logs    string `json:"logs"`
	/*"txHash":  receipt.TxHash.Hex(),
	  "status":  receipt.Status,
	  "logs":    receipt.Logs,
	  "block":   receipt.BlockNumber.String(),
	  "from":    receipt.From.Hex(),
	  "to":      receipt.To.Hex(),
	  "gasUsed": receipt.GasUsed,*/

}
