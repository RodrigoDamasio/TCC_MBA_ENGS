# Chaincode Quality Traceability (Fabric)

Go chaincode for tracking and validating quality test batches on Hyperledger Fabric.

## Stack
- Go (Golang) chaincode
- Hyperledger Fabric Contract API

## How to Use

1. **Install prerequisites**
    - Go 1.18+ (for chaincode build)
    - Fabric binaries and a running Fabric network

2. **Build chaincode**
    ```sh
    go mod tidy
    go build
    ```

3. **Package chaincode**
    ```sh
    peer lifecycle chaincode package qualitytest.tar.gz --path . --lang golang --label qualitytest_1
    ```

4. **Install and approve chaincode**
    - Install on peers:
      ```sh
      peer lifecycle chaincode install qualitytest.tar.gz
      ```
    - Approve and commit using Fabric CLI as per your network's process.

5. **Invoke chaincode**
    - Use Fabric Gateway SDK or `peer chaincode invoke` to call:
      - `PerformBatchQualityTest` method to store new batches
      - `QueryQualityTestBatch` to retrieve batch data

6. **Test chaincode**
    - Write and run unit tests in Go.
    - Deploy on a Fabric test-network for integration tests.

7. **Logs & Troubleshooting**
    - Check peer logs for chaincode errors.
    - Use `peer chaincode query` for debugging.

## Notes
- The contract stores NFT-like JSON objects for each batch.
- Update `main.go` for new features or bugfixes.
