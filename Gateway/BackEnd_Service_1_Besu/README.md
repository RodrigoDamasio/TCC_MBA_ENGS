# BackEnd_Service_1_Besu

Backend API service to interact with the Hyperledger Besu blockchain, built in Go using the Gin framework. Docker support included.

## Stack
- Go (Golang)
- Gin web framework
- Swagger for API docs (`swaggo`)
- AWS DynamoDB integration
- Docker/Docker Compose

## How to Use (Development)

1. **Install Go and dependencies**
    ```sh
    cd Gateway/BackEnd_Service_1_Besu/project
    go mod tidy
    ```

2. **Configure AWS and Besu**
    - Add your AWS credentials and DynamoDB config to `./CCInfra/DB/aws_config.yaml`.
    - Update Besu RPC settings in `./CCInfra/RPC_besu/`.

3. **Run the API locally**
    ```sh
    go run main.go
    # Or use Main_run.sh for AWS config setup and build/test/run tasks
    ./Main_run.sh build   # Build and run
    ./Main_run.sh test    # Run tests
    ./Main_run.sh         # Will default to go run main.go
    ```

4. **Access API**
    - Default port: `8082`
    - Health check: [http://localhost:8082/health](http://localhost:8082/health)
    - Swagger docs: [http://localhost:8082/swagger/index.html](http://localhost:8082/swagger/index.html)

5. **Build and run with Docker**
    ```sh
    docker build -t besu-backend .
    docker run --env-file .env -p 8082:8082 besu-backend
    # Or use docker-compose if provided
    ```

6. **Connect to Besu**
    - Ensure Besu node is running and accessible at the configured RPC/WS endpoint.

7. **API Endpoints**
    - `/batch` (POST): Create a batch quality test (calls Solidity contract)
    - `/batch/:id` (GET): Query batch by ID
    - `/batch/tx/:txHash` (GET): Query batch by transaction hash

8. **Unit Testing**
    ```sh
    go test ./...
    ```

9. **Logs & Troubleshooting**
    - Check `logs/` directory or Docker logs.
    - All errors and debug info are printed to stdout/stderr.

10. **Production**
    - Set `GIN_MODE=release`
    - Use Docker for consistent deployment.

## Notes
- Update `go.mod` and `go.sum` if you add dependencies.
- Ensure the Besu node is synced and accessible to avoid RPC errors.
