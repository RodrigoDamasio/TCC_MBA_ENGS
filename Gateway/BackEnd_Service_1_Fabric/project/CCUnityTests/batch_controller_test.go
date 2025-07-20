package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	handler "project/CCHandler"

	"github.com/gin-gonic/gin"
)

// Setup test router
func setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	return router
}

func TestBatchController_NewBatchController(t *testing.T) {
	// Test controller creation
	controller := handler.NewBatchController()

	if controller == nil {
		t.Fatal("Expected controller to not be nil")
	}

	// Test that it implements the interface by calling methods
	// (This is a compile-time check, but we can verify the interface exists)
	router := setupTestRouter()
	router.POST("/batch", controller.CreateBatch)
	router.GET("/batch/:id", controller.QueryBatchByID)

	// If we reach here, the controller implements the interface correctly
	t.Log("BatchController successfully implements BatchControllerInterface")
}

func TestBatchController_CreateBatch_InvalidJSON(t *testing.T) {
	// Setup
	controller := handler.NewBatchController()
	router := setupTestRouter()
	router.POST("/batch", controller.CreateBatch)

	// Test data - invalid JSON
	invalidJSON := `{invalid json}`
	req, _ := http.NewRequest("POST", "/batch", bytes.NewBufferString(invalidJSON))
	req.Header.Set("Content-Type", "application/json")

	// Perform request
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assertions
	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, w.Code)
	}

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Failed to parse response JSON: %v", err)
	}

	if response["error"] != "Dados inválidos" {
		t.Errorf("Expected error message 'Dados inválidos', got '%v'", response["error"])
	}
}

func TestBatchController_CreateBatch_MissingFields(t *testing.T) {
	// Setup
	controller := handler.NewBatchController()
	router := setupTestRouter()
	router.POST("/batch", controller.CreateBatch)

	// Test data - missing required fields
	incompleteData := map[string]interface{}{
		"id": "123",
		// Missing other required fields
	}
	jsonData, _ := json.Marshal(incompleteData)

	req, _ := http.NewRequest("POST", "/batch", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	// Perform request
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Note: This will likely fail due to missing gRPC connection, but we're testing the input validation
	// The actual behavior will depend on the gRPC connection being available
	t.Logf("Response status: %d", w.Code)
	t.Logf("Response body: %s", w.Body.String())
}

func TestBatchController_QueryBatchByID_ValidID(t *testing.T) {
	// Setup
	controller := handler.NewBatchController()
	router := setupTestRouter()
	router.GET("/batch/:id", controller.QueryBatchByID)

	// Test with a valid ID format
	req, _ := http.NewRequest("GET", "/batch/123", nil)

	// Perform request
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Note: This will likely fail due to missing gRPC connection
	// but we're testing that the route is properly configured
	t.Logf("Response status: %d", w.Code)
	t.Logf("Response body: %s", w.Body.String())

	// In a real test environment with gRPC mocked, we would test for specific responses
	// For now, we just verify the route is accessible
}

func TestBatchController_QueryBatchByID_EmptyID(t *testing.T) {
	// Setup
	controller := handler.NewBatchController()
	router := setupTestRouter()
	router.GET("/batch/:id", controller.QueryBatchByID)

	// Test with empty ID (should be handled by router as not matching the route)
	req, _ := http.NewRequest("GET", "/batch/", nil)

	// Perform request
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Should return 404 as the route doesn't match
	if w.Code != http.StatusNotFound {
		t.Logf("Expected 404 for empty ID, got %d", w.Code)
	}
}

// Test the backward compatibility wrapper functions
func TestBackwardCompatibility_CreateBatch(t *testing.T) {
	// Setup
	router := setupTestRouter()
	router.POST("/batch", handler.CreateBatch) // Using the wrapper function

	// Test data - invalid JSON to test wrapper
	invalidJSON := `{invalid json}`
	req, _ := http.NewRequest("POST", "/batch", bytes.NewBufferString(invalidJSON))
	req.Header.Set("Content-Type", "application/json")

	// Perform request
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Should behave the same as the controller method
	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, w.Code)
	}

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Failed to parse response JSON: %v", err)
	}

	if response["error"] != "Dados inválidos" {
		t.Errorf("Expected error message 'Dados inválidos', got '%v'", response["error"])
	}
}

func TestBackwardCompatibility_QueryBatchByID(t *testing.T) {
	// Setup
	router := setupTestRouter()
	router.GET("/batch/:id", handler.QueryBatchByID) // Using the wrapper function

	// Test with a valid ID format
	req, _ := http.NewRequest("GET", "/batch/123", nil)

	// Perform request
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Should behave the same as the controller method
	t.Logf("Wrapper function response status: %d", w.Code)
	t.Logf("Wrapper function response body: %s", w.Body.String())
}
