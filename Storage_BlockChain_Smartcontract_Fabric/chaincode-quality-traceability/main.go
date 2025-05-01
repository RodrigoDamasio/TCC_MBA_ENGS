package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/hyperledger/fabric-contract-api-go/v2/contractapi"
)

// QualityTestContract defines the smart contract structure.
type QualityTestContract struct {
	contractapi.Contract
}

// SampleResult represents the result of a single value validation.
type SampleResult struct {
	Value  int    `json:"value"`
	Result string `json:"result"` // "OK" or "NOK"
}

// QualityTestBatch represents the data structure for batch quality inspection.
type QualityTestBatch struct {
	ID              string         `json:"id"`
	MinLimit        int            `json:"min_limit"`
	MaxLimit        int            `json:"max_limit"`
	Values          []int          `json:"values"`
	ProductionOrder string         `json:"production_order"`
	TestDate        string         `json:"test_date"`
	TestDescription string         `json:"test_description"`
	SampleResults   []SampleResult `json:"sample_results"`
	FinalResult     string         `json:"final_result"` // "OK" or "NOK"
}

// PerformBatchQualityTest validates multiple values against the limits and stores the result in the ledger.
func (q *QualityTestContract) PerformBatchQualityTest(ctx contractapi.TransactionContextInterface, id string, minLimit int, maxLimit int, values []int, productionOrder string, testDate string, testDescription string) (*QualityTestBatch, error) {
	// Check if the ID already exists
	existingBatch, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("failed to check for existing batch: %v", err)
	}
	if existingBatch != nil {
		return nil, fmt.Errorf("batch with ID %s already exists", id)
	}

	var sampleResults []SampleResult
	finalResult := "OK"

	// Validate each value
	for _, value := range values {
		result := "OK"

		if value < minLimit || value > maxLimit {

			result = "NOK"
			finalResult = "NOK"

		}

		sampleResults = append(sampleResults, SampleResult{
			Value:  value,
			Result: result,
		})

	}

	// Create the QualityTestBatch object
	testBatch := QualityTestBatch{
		ID:              id,
		MinLimit:        minLimit,
		MaxLimit:        maxLimit,
		Values:          values,
		ProductionOrder: productionOrder,
		TestDate:        testDate,
		TestDescription: testDescription,
		SampleResults:   sampleResults,
		FinalResult:     finalResult,
	}

	// Serialize the object to JSON
	testBatchJSON, err := json.Marshal(testBatch)
	if err != nil {
		return nil, fmt.Errorf("failed to serialize test batch data: %v", err)
	}

	// Store the data in the ledger
	err = ctx.GetStub().PutState(id, testBatchJSON)
	if err != nil {
		return nil, fmt.Errorf("failed to store test batch data in the ledger: %v", err)
	}

	return &testBatch, nil
}

// QueryQualityTestBatch retrieves a quality test batch record from the ledger by its ID.
func (q *QualityTestContract) QueryQualityTestBatch(ctx contractapi.TransactionContextInterface, id string) (*QualityTestBatch, error) {
	// Retrieve the data from the ledger
	testBatchJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("failed to read test batch data from the ledger: %v", err)
	}
	if testBatchJSON == nil {
		return nil, fmt.Errorf("test batch with ID %s does not exist", id)
	}

	// Deserialize the JSON data
	var testBatch QualityTestBatch
	err = json.Unmarshal(testBatchJSON, &testBatch)
	if err != nil {
		return nil, fmt.Errorf("failed to deserialize test batch data: %v", err)
	}

	return &testBatch, nil
}

func main() {
	chaincode, err := contractapi.NewChaincode(new(QualityTestContract))
	if err != nil {
		log.Panicf("Error creating chaincode: %v", err)
	}

	if err := chaincode.Start(); err != nil {
		log.Panicf("Error starting chaincode: %v", err)
	}
}
