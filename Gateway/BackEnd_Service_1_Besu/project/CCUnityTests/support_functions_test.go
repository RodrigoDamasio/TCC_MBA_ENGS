package main

import (
	"math/big"
	"testing"

	entity "project/CCEntitys"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// Test para a função stringContains
func TestStringContains(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		substr   string
		expected bool
	}{
		{
			name:     "Contains substring case insensitive",
			s:        "Hello World",
			substr:   "world",
			expected: true,
		},
		{
			name:     "Contains substring exact case",
			s:        "Hello World",
			substr:   "World",
			expected: true,
		},
		{
			name:     "Does not contain substring",
			s:        "Hello World",
			substr:   "xyz",
			expected: false,
		},
		{
			name:     "Empty substring",
			s:        "Hello World",
			substr:   "",
			expected: true,
		},
		{
			name:     "Empty string",
			s:        "",
			substr:   "hello",
			expected: false,
		},
		{
			name:     "Both empty",
			s:        "",
			substr:   "",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := stringContains(tt.s, tt.substr)
			if result != tt.expected {
				t.Errorf("stringContains(%q, %q) = %v, expected %v", tt.s, tt.substr, result, tt.expected)
			}
		})
	}
}

// Test para ExtractBatchInfo
func TestExtractBatchInfo(t *testing.T) {
	// Criando dados de teste
	testRequest := entity.CreateBatchRequest{
		ID:              "123",
		MinLimit:        10,
		MaxLimit:        100,
		Values:          []int64{50, 60, 70},
		ProductionOrder: "PO-001",
		TestDate:        "2025-07-20",
		TestDescription: "Test Description",
		ImagesJson:      "{}",
	}

	// Simulando uma transação
	tx := types.NewTransaction(
		0,
		common.HexToAddress("0x0000000000000000000000000000000000000000"),
		big.NewInt(0),
		21000,
		big.NewInt(1000000000),
		nil,
	)

	// Simulando um recibo com logs
	receipt := &types.Receipt{
		Status: types.ReceiptStatusSuccessful,
		Logs: []*types.Log{
			{
				Data: []byte("dummy"),
			},
			{
				Data: []byte("OK"), // Log que indica sucesso
			},
		},
		TxHash: tx.Hash(),
	}

	// Testando a função
	result, err := ExtractBatchInfo(receipt, tx, testRequest)

	// Verificações
	if err != nil {
		t.Fatalf("ExtractBatchInfo returned error: %v", err)
	}

	if result.ID != 123 {
		t.Errorf("Expected ID 123, got %d", result.ID)
	}

	if result.ProductionOrder != "PO-001" {
		t.Errorf("Expected ProductionOrder 'PO-001', got %s", result.ProductionOrder)
	}

	if result.FinalResult != "OK" {
		t.Errorf("Expected FinalResult 'OK', got %s", result.FinalResult)
	}

	if result.Data_daytime != "2025-07-20" {
		t.Errorf("Expected Data_daytime '2025-07-20', got %s", result.Data_daytime)
	}

	if result.Hash != tx.Hash().Hex() {
		t.Errorf("Expected Hash %s, got %s", tx.Hash().Hex(), result.Hash)
	}
}

// Test para ExtractBatchInfo com log NOK
func TestExtractBatchInfo_NOK(t *testing.T) {
	testRequest := entity.CreateBatchRequest{
		ID:              "456",
		ProductionOrder: "PO-002",
		TestDate:        "2025-07-21",
	}

	tx := types.NewTransaction(
		1,
		common.HexToAddress("0x0000000000000000000000000000000000000001"),
		big.NewInt(0),
		21000,
		big.NewInt(1000000000),
		nil,
	)

	// Recibo com log indicando falha
	receipt := &types.Receipt{
		Status: types.ReceiptStatusSuccessful,
		Logs: []*types.Log{
			{
				Data: []byte("dummy"),
			},
			{
				Data: []byte("NOK"), // Log que indica falha
			},
		},
		TxHash: tx.Hash(),
	}

	result, err := ExtractBatchInfo(receipt, tx, testRequest)

	if err != nil {
		t.Fatalf("ExtractBatchInfo returned error: %v", err)
	}

	if result.FinalResult != "NOK" {
		t.Errorf("Expected FinalResult 'NOK', got %s", result.FinalResult)
	}
}

// Test para ExtractBatchInfo com log vazio
func TestExtractBatchInfo_EmptyLog(t *testing.T) {
	testRequest := entity.CreateBatchRequest{
		ID:              "789",
		ProductionOrder: "PO-003",
		TestDate:        "2025-07-22",
	}

	tx := types.NewTransaction(
		2,
		common.HexToAddress("0x0000000000000000000000000000000000000002"),
		big.NewInt(0),
		21000,
		big.NewInt(1000000000),
		nil,
	)

	// Recibo com log vazio
	receipt := &types.Receipt{
		Status: types.ReceiptStatusSuccessful,
		Logs: []*types.Log{
			{
				Data: []byte("dummy"),
			},
			{
				Data: []byte(""), // Log vazio
			},
		},
		TxHash: tx.Hash(),
	}

	result, err := ExtractBatchInfo(receipt, tx, testRequest)

	if err != nil {
		t.Fatalf("ExtractBatchInfo returned error: %v", err)
	}

	if result.FinalResult != "Log vazio" {
		t.Errorf("Expected FinalResult 'Log vazio', got %s", result.FinalResult)
	}
}

// Test para ExtractBatchInfo com ID inválido
func TestExtractBatchInfo_InvalidID(t *testing.T) {
	testRequest := entity.CreateBatchRequest{
		ID:              "invalid_id", // ID não numérico
		ProductionOrder: "PO-004",
		TestDate:        "2025-07-23",
	}

	tx := types.NewTransaction(
		3,
		common.HexToAddress("0x0000000000000000000000000000000000000003"),
		big.NewInt(0),
		21000,
		big.NewInt(1000000000),
		nil,
	)

	receipt := &types.Receipt{
		Status: types.ReceiptStatusSuccessful,
		Logs: []*types.Log{
			{
				Data: []byte("dummy"),
			},
			{
				Data: []byte("OK"),
			},
		},
		TxHash: tx.Hash(),
	}

	_, err := ExtractBatchInfo(receipt, tx, testRequest)

	if err == nil {
		t.Error("Expected error for invalid ID, but got nil")
	}
}

// Test do BatchController constructor
func TestNewBatchController(t *testing.T) {
	controller := NewBatchController()

	if controller == nil {
		t.Error("NewBatchController returned nil")
	}

	// Verificar se implementa a interface
	var _ BatchControllerInterface = controller
}
