package handler_test

import (
	"encoding/json"
	"testing"

	entity "project/CCEntitys"
)

func TestTransactionHyperledger_BatchInfo_JSONMarshaling(t *testing.T) {
	// Test data
	batchInfo := entity.TransactionHyperledger_BatchInfo{
		ID:              "123",
		ProductionOrder: "PO-2024-001",
		FinalResult:     "APPROVED",
		Data_daytime:    "2024-07-20T10:30:00Z",
	}

	// Test JSON marshaling
	jsonData, err := json.Marshal(batchInfo)
	if err != nil {
		t.Errorf("Failed to marshal to JSON: %v", err)
	}

	// Test JSON unmarshaling
	var unmarshaled entity.TransactionHyperledger_BatchInfo
	err = json.Unmarshal(jsonData, &unmarshaled)
	if err != nil {
		t.Errorf("Failed to unmarshal from JSON: %v", err)
	}

	// Verify values
	if unmarshaled.ID != batchInfo.ID {
		t.Errorf("Expected ID '%s', got '%s'", batchInfo.ID, unmarshaled.ID)
	}

	if unmarshaled.ProductionOrder != batchInfo.ProductionOrder {
		t.Errorf("Expected ProductionOrder '%s', got '%s'", batchInfo.ProductionOrder, unmarshaled.ProductionOrder)
	}

	if unmarshaled.FinalResult != batchInfo.FinalResult {
		t.Errorf("Expected FinalResult '%s', got '%s'", batchInfo.FinalResult, unmarshaled.FinalResult)
	}

	if unmarshaled.Data_daytime != batchInfo.Data_daytime {
		t.Errorf("Expected Data_daytime '%s', got '%s'", batchInfo.Data_daytime, unmarshaled.Data_daytime)
	}
}

func TestTransactionHyperledger_BatchInfo_JSONTags(t *testing.T) {
	// Test that JSON tags work correctly
	batchInfo := entity.TransactionHyperledger_BatchInfo{
		ID:              "456",
		ProductionOrder: "PO-2024-002",
		FinalResult:     "REJECTED",
		Data_daytime:    "2024-07-20T15:45:00Z",
	}

	jsonData, err := json.Marshal(batchInfo)
	if err != nil {
		t.Errorf("Failed to marshal to JSON: %v", err)
	}

	// Check that JSON contains the expected field names (based on JSON tags)
	jsonStr := string(jsonData)

	expectedFields := []string{
		`"id":"456"`,
		`"production_order":"PO-2024-002"`,
		`"final_result":"REJECTED"`,
		`"test_Date":"2024-07-20T15:45:00Z"`,
	}

	for _, field := range expectedFields {
		if !contains(jsonStr, field) {
			t.Errorf("Expected JSON to contain '%s', but got: %s", field, jsonStr)
		}
	}
}

func TestTransactionData_Structure(t *testing.T) {
	// Test TransactionData structure
	transactionData := entity.TransactionData{
		ID:              789,
		ProductionOrder: "PO-2024-003",
		FinalResult:     "APPROVED",
		Hash:            "hash_123456789",
		Data_daytime:    "2024-07-20T16:00:00Z",
	}

	// Verify all fields can be set and retrieved
	if transactionData.ID != 789 {
		t.Errorf("Expected ID 789, got %d", transactionData.ID)
	}

	if transactionData.ProductionOrder != "PO-2024-003" {
		t.Errorf("Expected ProductionOrder 'PO-2024-003', got '%s'", transactionData.ProductionOrder)
	}

	if transactionData.FinalResult != "APPROVED" {
		t.Errorf("Expected FinalResult 'APPROVED', got '%s'", transactionData.FinalResult)
	}

	if transactionData.Hash != "hash_123456789" {
		t.Errorf("Expected Hash 'hash_123456789', got '%s'", transactionData.Hash)
	}

	if transactionData.Data_daytime != "2024-07-20T16:00:00Z" {
		t.Errorf("Expected Data_daytime '2024-07-20T16:00:00Z', got '%s'", transactionData.Data_daytime)
	}
}

func TestFabric_Structure(t *testing.T) {
	// Test Fabric configuration structure
	fabricConfig := entity.Fabric{
		MSPID:         "Org1MSP",
		CryptoPath:    "/path/to/crypto",
		PeerEndpoint:  "localhost:7051",
		GatewayPeer:   "peer0.org1.example.com",
		ChannelName:   "mychannel",
		ChaincodeName: "mychaincode",
	}

	fabricConfig.Paths.CertPath = "/path/to/cert"
	fabricConfig.Paths.KeyPath = "/path/to/key"
	fabricConfig.Paths.TLSCertPath = "/path/to/tls"

	// Verify all fields
	if fabricConfig.MSPID != "Org1MSP" {
		t.Errorf("Expected MSPID 'Org1MSP', got '%s'", fabricConfig.MSPID)
	}

	if fabricConfig.CryptoPath != "/path/to/crypto" {
		t.Errorf("Expected CryptoPath '/path/to/crypto', got '%s'", fabricConfig.CryptoPath)
	}

	if fabricConfig.PeerEndpoint != "localhost:7051" {
		t.Errorf("Expected PeerEndpoint 'localhost:7051', got '%s'", fabricConfig.PeerEndpoint)
	}

	if fabricConfig.GatewayPeer != "peer0.org1.example.com" {
		t.Errorf("Expected GatewayPeer 'peer0.org1.example.com', got '%s'", fabricConfig.GatewayPeer)
	}

	if fabricConfig.ChannelName != "mychannel" {
		t.Errorf("Expected ChannelName 'mychannel', got '%s'", fabricConfig.ChannelName)
	}

	if fabricConfig.ChaincodeName != "mychaincode" {
		t.Errorf("Expected ChaincodeName 'mychaincode', got '%s'", fabricConfig.ChaincodeName)
	}

	// Test nested Paths structure
	if fabricConfig.Paths.CertPath != "/path/to/cert" {
		t.Errorf("Expected CertPath '/path/to/cert', got '%s'", fabricConfig.Paths.CertPath)
	}

	if fabricConfig.Paths.KeyPath != "/path/to/key" {
		t.Errorf("Expected KeyPath '/path/to/key', got '%s'", fabricConfig.Paths.KeyPath)
	}

	if fabricConfig.Paths.TLSCertPath != "/path/to/tls" {
		t.Errorf("Expected TLSCertPath '/path/to/tls', got '%s'", fabricConfig.Paths.TLSCertPath)
	}
}

// Helper function to check if a string contains a substring
func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(substr) == 0 ||
		(len(s) > len(substr) && (s[:len(substr)] == substr ||
			s[len(s)-len(substr):] == substr ||
			findInString(s, substr))))
}

func findInString(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
