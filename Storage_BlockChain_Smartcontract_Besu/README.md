# Smart Contracts for Besu (Hardhat Project)

This project uses Hardhat for developing, testing, and deploying Solidity smart contracts on Hyperledger Besu.

## Stack
- Solidity (contracts/)
- Hardhat (Node.js)
- Ethers.js
- Mocha/Chai (testing)

## How to Use (Development)

1. **Install dependencies**
    ```sh
    npm install
    ```

2. **Compile contracts**
    ```sh
    npx hardhat compile
    ```

3. **Run tests**
    ```sh
    npx hardhat test
    ```

4. **Deploy to Besu**
    - Update `hardhat.config.js` with Besu RPC and account.
    - Run:
      ```sh
      npx hardhat run scripts/deploy_besu.js --network besu
      ```

5. **Interact with contracts**
    - Use scripts like `scripts/Assina_Besu.js` (requires ethers.js) to sign and interact.
    - Example:
      ```sh
      node scripts/Assina_Besu.js
      ```

6. **Ignition Deployment**
    - Use Hardhat Ignition for advanced deployments:
      ```sh
      npx hardhat ignition deploy ./ignition/modules/Lock.js
      ```

7. **Configuration**
    - Set up network details and private keys in `hardhat.config.js`.

8. **Troubleshooting**
    - Check Hardhat output and logs for errors.
    - Ensure Besu node is running and accessible.

9. **Production**
    - Use environment variables for secrets and private keys.

## Notes
- All contract code is in `contracts/`.
- Tests are in `test/`.
- Scripts for deployment/interaction are in `scripts/`.
