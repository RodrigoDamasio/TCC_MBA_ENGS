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

type Fabric struct {
	MSPID         string `yaml:"msp_id"`
	CryptoPath    string `yaml:"crypto_path"`
	PeerEndpoint  string `yaml:"peer_endpoint"`
	GatewayPeer   string `yaml:"gateway_peer"`
	ChannelName   string `yaml:"channel_name"`
	ChaincodeName string `yaml:"chaincode_name"`
	Paths         struct {
		CertPath    string `yaml:"cert_path"`
		KeyPath     string `yaml:"key_path"`
		TLSCertPath string `yaml:"tls_cert_path"`
	} `yaml:"paths"`
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
