const { ethers } = require("ethers");

// Configurações da rede e credenciais
const providerUrl = "http://localhost:8551";
const privateKey = "0xf8060dd8211fd80d62023e74aff54d870e13835c630865d6250761394efddc98";
const contractAddress = "0x16fb14f423821b20c5da2b14a6b2cfbb061b71cf";

async function interactWithContract() {
    const provider = new ethers.JsonRpcProvider(providerUrl);
    const wallet = new ethers.Wallet(privateKey, provider);

    // ABI ajustada para o contrato NFT
    const contractAbi = [
        "function performBatchQualityTest(string id, int256 minLimit, int256 maxLimit, int256[] values, string productionOrder, string testDate, string testDescription, string imagesJson) public",
        "function queryQualityTestBatch(string id) public view returns (string, int256, int256, int256[], string, string, string, tuple(int256 value, string result)[], string, string)"
    ];

    const contract = new ethers.Contract(contractAddress, contractAbi, wallet);

    const batchId = "batch003";
    const minLimit = 10;
    const maxLimit = 100;
    const values = [12, 50, 120, 85];
    const productionOrder = "PO12345";
    const testDate = "2025-05-01";
    const testDescription = "Teste de qualidade de exemplo";

    // JSON formatado padrão tokenURI
    const imagesJson = JSON.stringify({
        name: `Batch NFT #${batchId}`,
        description: testDescription,
        image: "https://tcc-contract-images-besu.s3.us-east-1.amazonaws.com/block_Image.jpeg",
        attributes: [
            { trait_type: "Ordem de Produção", value: productionOrder },
            { trait_type: "Limite Mínimo", value: minLimit },
            { trait_type: "Limite Máximo", value: maxLimit },
            { trait_type: "Data do Teste", value: testDate },
            { trait_type: "Valores", value: values.join(", ") }
        ]
    });

    try {
        console.log(`Executando teste de qualidade no contrato ${contractAddress}`);
        const txResponse = await contract.performBatchQualityTest(
            batchId,
            minLimit,
            maxLimit,
            values,
            productionOrder,
            testDate,
            testDescription,
            imagesJson // <-- JSON padrão tokenURI
        );

        console.log(`Transação enviada: ${txResponse.hash}`);
        await txResponse.wait();
        console.log(`Transação confirmada!`);

        const batchData = await contract.queryQualityTestBatch(batchId);

        console.log(`Dados do lote de teste recuperados:`);
        console.log(`ID: ${batchData[0]}`);
        console.log(`Min Limit: ${batchData[1].toString()}`);
        console.log(`Max Limit: ${batchData[2].toString()}`);
        console.log(`Valores: ${batchData[3].map((val) => val.toString())}`);
        console.log(`Ordem de Produção: ${batchData[4]}`);
        console.log(`Data do Teste: ${batchData[5]}`);
        console.log(`Descrição do Teste: ${batchData[6]}`);
        
        const sampleResults = batchData[7].map((sample) => ({
            value: sample.value.toString(),
            result: sample.result
        }));
        console.log(`Resultados das Amostras: ${JSON.stringify(sampleResults)}`);
        console.log(`Resultado Final: ${batchData[8]}`);
        console.log(`tokenURI/Imagens (JSON): ${batchData[9]}`);
    } catch (error) {
        console.error(`Erro ao interagir com o contrato: ${error.message}`);
    }
}

interactWithContract();