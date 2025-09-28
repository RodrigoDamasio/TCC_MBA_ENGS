# Hyperledger Besu Otimizado

Docker-based deployment of an optimized Hyperledger Besu blockchain. Includes custom genesis and config files.

## Stack
- Hyperledger Besu (official Docker image)
- Custom genesis and config
- Docker

## Deployment Steps

1. **Install Docker**  
   - Required: [Docker](https://docs.docker.com/get-docker/) (and optionally Docker Compose).

2. **Prepare Configuration**  
   - Edit `config/genesis.json`, `besu-config.json`, and `toEncode.json` as needed for your network topology.

3. **Build Docker Image**  
   ```sh
   cd Infrastructure/Hyperledger\ Besu\ Otimizado
   docker build -t besu-node .
   ```

4. **Run the Besu Node**
   ```sh
   docker run -d --name besu-node -p 8545:8545 -p 8546:8546 -p 30303:30303 besu-node
   ```

5. **Access the Node**
   - JSON-RPC: `http://localhost:8545`
   - WS: `ws://localhost:8546`

6. **Check Logs**
   ```sh
   docker logs -f besu-node
   ```

7. **Network Customization**
   - For QBFT, extraData is generated and inserted at build time via Dockerfile scripts.
   - Adjust config for number of validators, bootnodes, etc.

8. **Shut Down**
   ```sh
   docker stop besu-node
   docker rm besu-node
   ```

9. **Persistence**
   - Use Docker volumes to persist blockchain data across restarts.

10. **Security**
   - Expose only necessary ports; secure JSON-RPC in production.

## Notes
- You can extend this setup for multi-node networks using Docker Compose.
