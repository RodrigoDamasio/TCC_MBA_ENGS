package handler_test

import (
	"testing"

	handler "project/CCHandler"
)

func TestExtractBatchInfo_Success(t *testing.T) {
	// Test data - valid JSON
	jsonStr := `{
		"id": "123",
		"production_order": "PO-2024-001",
		"final_result": "APPROVED",
		"test_Date": "2024-07-20T10:30:00Z"
	}`
	transactionHash := "tx_hash_123456"

	// Execute function
	result, err := handler.ExtractBatchInfo(jsonStr, transactionHash)

	// Assertions
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if result == nil {
		t.Fatal("Expected result to not be nil")
	}

	// Check values
	if result.ID != 123 {
		t.Errorf("Expected ID to be 123, got %d", result.ID)
	}

	if result.ProductionOrder != "PO-2024-001" {
		t.Errorf("Expected ProductionOrder to be 'PO-2024-001', got '%s'", result.ProductionOrder)
	}

	if result.FinalResult != "APPROVED" {
		t.Errorf("Expected FinalResult to be 'APPROVED', got '%s'", result.FinalResult)
	}

	if result.Hash != transactionHash {
		t.Errorf("Expected Hash to be '%s', got '%s'", transactionHash, result.Hash)
	}

	if result.Data_daytime != "2024-07-20T10:30:00Z" {
		t.Errorf("Expected Data_daytime to be '2024-07-20T10:30:00Z', got '%s'", result.Data_daytime)
	}
}

func TestExtractBatchInfo_InvalidJSON(t *testing.T) {
	// Test data - invalid JSON
	jsonStr := `{invalid json}`
	transactionHash := "tx_hash_123456"

	// Execute function
	result, err := handler.ExtractBatchInfo(jsonStr, transactionHash)

	// Assertions
	if err == nil {
		t.Error("Expected error for invalid JSON, got nil")
	}

	if result != nil {
		t.Error("Expected result to be nil for invalid JSON")
	}
}

func TestExtractBatchInfo_InvalidID(t *testing.T) {
	// Test data - ID is not a valid integer
	jsonStr := `{
		"id": "not_a_number",
		"production_order": "PO-2024-001",
		"final_result": "APPROVED",
		"test_Date": "2024-07-20T10:30:00Z"
	}`
	transactionHash := "tx_hash_123456"

	// Execute function
	result, err := handler.ExtractBatchInfo(jsonStr, transactionHash)

	// Assertions
	if err == nil {
		t.Error("Expected error for invalid ID, got nil")
	}

	if result != nil {
		t.Error("Expected result to be nil for invalid ID")
	}
}

func TestExtractBatchInfo_EmptyJSON(t *testing.T) {
	// Test data - empty JSON
	jsonStr := `{}`
	transactionHash := "tx_hash_123456"

	// Execute function
	result, err := handler.ExtractBatchInfo(jsonStr, transactionHash)

	// Assertions - should handle empty JSON gracefully
	if err == nil {
		t.Error("Expected error for empty ID field, got nil")
	}

	if result != nil {
		t.Error("Expected result to be nil for empty JSON")
	}
}
