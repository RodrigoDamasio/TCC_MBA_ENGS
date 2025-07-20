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

type AWS struct {
	Credentials struct {
		AccessKeyID     string `yaml:"access_key_id"`
		SecretAccessKey string `yaml:"secret_access_key"`
	} `yaml:"credentials"`
	Config struct {
		Region string `yaml:"region"`
		Output string `yaml:"output"`
	} `yaml:"config"`
	DynamoDB struct {
		Endpoint    string `yaml:"endpoint"`
		TablePrefix string `yaml:"table_prefix"`
	} `yaml:"dynamodb"`
	Session struct {
		Token   string `yaml:"token"`
		Profile string `yaml:"profile"`
	} `yaml:"session"`
}

type RPC struct {
	ContractAddress string `yaml:"contract_address"`
	PrivateKey      string `yaml:"private_key"`
	ChainID         int64  `yaml:"chain_id"`
	RPCURL          string `yaml:"rpc_url"`
}
