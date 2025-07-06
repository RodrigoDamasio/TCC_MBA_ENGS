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
