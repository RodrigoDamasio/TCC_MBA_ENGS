package main

import (
	"fmt"
	"testing"
)

// Test básico para verificar se o ambiente de teste está funcionando
func TestBasicFunctionality(t *testing.T) {
	// Test simple arithmetic
	result := 2 + 2
	expected := 4

	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

// Test para verificar string operations
func TestStringOperations(t *testing.T) {
	input := "Hello World"
	expected := "HELLO WORLD"

	// Simple string conversion
	result := fmt.Sprintf("%s", input)
	if len(result) != len(input) {
		t.Errorf("String length mismatch. Expected %d, got %d", len(input), len(result))
	}
}

// Test para verificar slice operations
func TestSliceOperations(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}

	// Test slice length
	if len(numbers) != 5 {
		t.Errorf("Expected slice length 5, got %d", len(numbers))
	}

	// Test slice access
	if numbers[0] != 1 {
		t.Errorf("Expected first element to be 1, got %d", numbers[0])
	}

	if numbers[len(numbers)-1] != 5 {
		t.Errorf("Expected last element to be 5, got %d", numbers[len(numbers)-1])
	}
}

// Test para verificar map operations
func TestMapOperations(t *testing.T) {
	testMap := make(map[string]int)
	testMap["key1"] = 100
	testMap["key2"] = 200

	// Test map access
	if testMap["key1"] != 100 {
		t.Errorf("Expected testMap[key1] = 100, got %d", testMap["key1"])
	}

	// Test map length
	if len(testMap) != 2 {
		t.Errorf("Expected map length 2, got %d", len(testMap))
	}

	// Test key existence
	_, exists := testMap["key3"]
	if exists {
		t.Error("Expected key3 to not exist in map")
	}
}

// Test para verificar interface simples
type TestInterface interface {
	GetValue() int
}

type TestStruct struct {
	value int
}

func (ts TestStruct) GetValue() int {
	return ts.value
}

func TestInterfaceImplementation(t *testing.T) {
	ts := TestStruct{value: 42}

	var ti TestInterface = ts

	result := ti.GetValue()
	if result != 42 {
		t.Errorf("Expected GetValue() to return 42, got %d", result)
	}
}

// Test para verificar error handling
func TestErrorHandling(t *testing.T) {
	// Function that returns an error
	testFunc := func(shouldError bool) error {
		if shouldError {
			return fmt.Errorf("test error")
		}
		return nil
	}

	// Test no error case
	err := testFunc(false)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Test error case
	err = testFunc(true)
	if err == nil {
		t.Error("Expected error, got nil")
	}
}

// Benchmark test simples
func BenchmarkStringConcatenation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("test_%d", i)
	}
}

// Test table-driven
func TestTableDriven(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{"positive", 5, 10},
		{"zero", 0, 0},
		{"negative", -3, -6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.input * 2
			if result != tt.expected {
				t.Errorf("Expected %d * 2 = %d, got %d", tt.input, tt.expected, result)
			}
		})
	}
}
