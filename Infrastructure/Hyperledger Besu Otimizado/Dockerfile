# Use the official Hyperledger Besu image
FROM hyperledger/besu:latest

# Set the working directory
WORKDIR /opt/besu

# Copy the configuration files into the container
COPY ./config/genesis.json /opt/besu/config/genesis.json
COPY ./config/besu-config.json /opt/besu/config/besu-config.json
COPY ./config/toEncode.json /opt/besu/config/toEncode.json

# Expose ports for p2p and RPC communication
EXPOSE 30303 8545 8546

USER root
# Generate extraData using the RLP encode command during build
RUN besu rlp encode --from=/opt/besu/config/toEncode.json --type=QBFT_EXTRA_DATA > /opt/besu/config/extraData.txt

# Modify the genesis.json file to include the generated extraData
RUN EXTRA_DATA=$(cat /opt/besu/config/extraData.txt) && \
    sed -i 's/"extraData": ""/"extraData": "'$EXTRA_DATA'"/' /opt/besu/config/genesis.json

# Instala curl e net-tools (para ifconfig)
RUN apt-get update && \
    apt-get install -y curl net-tools && \
    apt-get install -y iputils-ping &&\
    apt-get install -y telnet &&\
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

# Define the entrypoint to run besu with the proper configurations
ENTRYPOINT ["besu"]

