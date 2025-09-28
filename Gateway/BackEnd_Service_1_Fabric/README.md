# BackEnd_Service_1_Fabric

Backend API for Hyperledger Fabric, written in Go using Gin. Integrates with AWS DynamoDB and Fabric Gateway SDK. Docker-ready.

## Stack
- Go (Golang)
- Gin web framework
- Swagger (`swaggo`)
- AWS DynamoDB
- Hyperledger Fabric Gateway SDK
- Docker

## How to Use (Development)

1. **Install Go and dependencies**
    ```sh
    cd Gateway/BackEnd_Service_1_Fabric/project
    go mod tidy
    ```

2. **Configure AWS and Fabric**
    - Add AWS credentials/config to `./CCInfra/DB/aws_config.yaml`.
    - Add Fabric network config to `./CCInfra/Grpc_Hyperledger/fabric_config.yaml`.
    - Place certificates and organizations in `./organizations/`.

3. **Run API locally**
    ```sh
    go run main.go
    # Or use Main_run.sh for AWS/Fabric setup and build/test/run
    ./Main_run.sh build   # Build and run
    ./Main_run.sh test    # Run tests
    ./Main_run.sh         # Defaults to go run main.go
    ```

4. **Access API**
    - Default port: `8081`
    - Health check: [http://localhost:8081/health](http://localhost:8081/health)
    - Swagger docs: [http://localhost:8081/swagger/index.html](http://localhost:8081/swagger/index.html)

5. **Build and run with Docker**
    ```sh
    docker build -t fabric-backend .
    docker run --env-file .env -p 8081:8081 fabric-backend
    # Or use docker-compose if provided
    ```

6. **Connect to Fabric**
    - Make sure Fabric network is up and chaincode deployed.
    - The backend uses Fabric Gateway SDK to submit/query transactions.

7. **API Endpoints**
    - `/batch` (POST): Create a quality test batch (calls Fabric chaincode)
    - `/batch/:id` (GET): Query batch by ID

8. **Unit Testing**
    ```sh
    go test ./...
    ```

9. **Logs & Troubleshooting**
    - Check logs in the Docker container or stdout.
    - Use verbose output for debugging.

10. **Production**
    - Set `GIN_MODE=release`
    - Use Docker for consistent deployment and environment variables for secrets.

## Notes
- Update Go dependencies via `go mod tidy` as needed.
- Ensure all Fabric certificates and connections are valid.
