package main

import (
	entity "project/CCEntitys"
	"testing"
)

// Test para verificar se as estruturas de entidade estão corretas
func TestCreateBatchRequest(t *testing.T) {
	request := entity.CreateBatchRequest{
		ID:              "123",
		MinLimit:        10,
		MaxLimit:        100,
		Values:          []int64{50, 60, 70},
		ProductionOrder: "PO-001",
		TestDate:        "2025-07-20",
		TestDescription: "Test Description",
		ImagesJson:      "{}",
	}

	// Verificar se os campos foram definidos corretamente
	if request.ID != "123" {
		t.Errorf("Expected ID '123', got '%s'", request.ID)
	}

	if request.MinLimit != 10 {
		t.Errorf("Expected MinLimit 10, got %d", request.MinLimit)
	}

	if request.MaxLimit != 100 {
		t.Errorf("Expected MaxLimit 100, got %d", request.MaxLimit)
	}

	if len(request.Values) != 3 {
		t.Errorf("Expected 3 values, got %d", len(request.Values))
	}

	if request.ProductionOrder != "PO-001" {
		t.Errorf("Expected ProductionOrder 'PO-001', got '%s'", request.ProductionOrder)
	}

	if request.TestDate != "2025-07-20" {
		t.Errorf("Expected TestDate '2025-07-20', got '%s'", request.TestDate)
	}
}

func TestTransactionData(t *testing.T) {
	data := entity.TransactionData{
		ID:              123,
		ProductionOrder: "PO-001",
		FinalResult:     "OK",
		Hash:            "0x1234567890abcdef",
		Data_daytime:    "2025-07-20",
	}

	// Verificar se os campos foram definidos corretamente
	if data.ID != 123 {
		t.Errorf("Expected ID 123, got %d", data.ID)
	}

	if data.ProductionOrder != "PO-001" {
		t.Errorf("Expected ProductionOrder 'PO-001', got '%s'", data.ProductionOrder)
	}

	if data.FinalResult != "OK" {
		t.Errorf("Expected FinalResult 'OK', got '%s'", data.FinalResult)
	}

	if data.Hash != "0x1234567890abcdef" {
		t.Errorf("Expected Hash '0x1234567890abcdef', got '%s'", data.Hash)
	}

	if data.Data_daytime != "2025-07-20" {
		t.Errorf("Expected Data_daytime '2025-07-20', got '%s'", data.Data_daytime)
	}
}

func TestTransactionHyperledger_BatchInfo(t *testing.T) {
	info := entity.TransactionHyperledger_BatchInfo{
		ID:              "123",
		ProductionOrder: "PO-001",
		FinalResult:     "OK",
		Data_daytime:    "2025-07-20",
	}

	// Verificar se os campos foram definidos corretamente
	if info.ID != "123" {
		t.Errorf("Expected ID '123', got '%s'", info.ID)
	}

	if info.ProductionOrder != "PO-001" {
		t.Errorf("Expected ProductionOrder 'PO-001', got '%s'", info.ProductionOrder)
	}

	if info.FinalResult != "OK" {
		t.Errorf("Expected FinalResult 'OK', got '%s'", info.FinalResult)
	}

	if info.Data_daytime != "2025-07-20" {
		t.Errorf("Expected Data_daytime '2025-07-20', got '%s'", info.Data_daytime)
	}
}

// Test para validar se os valores são processados corretamente
func TestCreateBatchRequest_Values(t *testing.T) {
	values := []int64{10, 20, 30, 40, 50}

	request := entity.CreateBatchRequest{
		ID:       "test",
		MinLimit: 5,
		MaxLimit: 55,
		Values:   values,
	}

	// Verificar se todos os valores estão dentro dos limites
	allWithinLimits := true
	for _, value := range request.Values {
		if value < request.MinLimit || value > request.MaxLimit {
			allWithinLimits = false
			break
		}
	}

	if !allWithinLimits {
		t.Error("Some values are outside the specified limits")
	}

	// Verificar se o slice foi copiado corretamente
	if len(request.Values) != len(values) {
		t.Errorf("Expected %d values, got %d", len(values), len(request.Values))
	}

	for i, value := range values {
		if request.Values[i] != value {
			t.Errorf("Expected value[%d] = %d, got %d", i, value, request.Values[i])
		}
	}
}

// Test de edge cases
func TestCreateBatchRequest_EdgeCases(t *testing.T) {
	t.Run("Empty values slice", func(t *testing.T) {
		request := entity.CreateBatchRequest{
			ID:     "empty",
			Values: []int64{},
		}

		if len(request.Values) != 0 {
			t.Errorf("Expected empty values slice, got length %d", len(request.Values))
		}
	})

	t.Run("Single value", func(t *testing.T) {
		request := entity.CreateBatchRequest{
			ID:     "single",
			Values: []int64{42},
		}

		if len(request.Values) != 1 {
			t.Errorf("Expected single value, got %d values", len(request.Values))
		}

		if request.Values[0] != 42 {
			t.Errorf("Expected value 42, got %d", request.Values[0])
		}
	})

	t.Run("Large values", func(t *testing.T) {
		largeValue := int64(999999999999)
		request := entity.CreateBatchRequest{
			ID:     "large",
			Values: []int64{largeValue},
		}

		if request.Values[0] != largeValue {
			t.Errorf("Expected large value %d, got %d", largeValue, request.Values[0])
		}
	})
}

// Benchmark test para criação de estruturas
func BenchmarkCreateBatchRequest(b *testing.B) {
	values := []int64{10, 20, 30, 40, 50}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = entity.CreateBatchRequest{
			ID:              "benchmark",
			MinLimit:        1,
			MaxLimit:        100,
			Values:          values,
			ProductionOrder: "PO-BENCH",
			TestDate:        "2025-07-20",
			TestDescription: "Benchmark test",
			ImagesJson:      "{}",
		}
	}
}
