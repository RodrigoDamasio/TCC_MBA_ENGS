const { ethers } = require("ethers");

// Configurações da rede e credenciais
const providerUrl = "http://localhost:8551"; // ou outra URL de provedor
const privateKey = "0xf8060dd8211fd80d62023e74aff54d870e13835c630865d6250761394efddc98"; // Chave privada da conta de origem
const contractAddress = "0x370ca606b26d87a894f61395d0eefddb7161a9cb"; // Endereço do contrato QualityTestContract

async function interactWithContract() {
    // Cria um provedor para interagir com a blockchain
    const provider = new ethers.JsonRpcProvider(providerUrl);

    // Cria uma carteira a partir da chave privada
    const wallet = new ethers.Wallet(privateKey, provider);

    // Interface do contrato QualityTestContract
    const contractAbi = [
        // ABI mínima necessária para as funções
        "function performBatchQualityTest(string id, int256 minLimit, int256 maxLimit, int256[] values, string productionOrder, string testDate, string testDescription) public",
        "function queryQualityTestBatch(string id) public view returns (string, int256, int256, int256[], string, string, string, tuple(int256 value, string result)[], string)"
    ];

    // Cria uma instância do contrato
    const contract = new ethers.Contract(contractAddress, contractAbi, wallet);

    // Dados do lote de teste
    const batchId = "batch003";
    const minLimit = 10;
    const maxLimit = 100;
    const values = [12, 50, 120, 85];
    const productionOrder = "PO12345";
    const testDate = "2025-05-01";
    const testDescription = "Teste de qualidade de exemplo";

    try {
        // Chama a função performBatchQualityTest do contrato
        console.log(`Executando teste de qualidade no contrato ${contractAddress}`);
        const txResponse = await contract.performBatchQualityTest(
            batchId,
            minLimit,
            maxLimit,
            values,
            productionOrder,
            testDate,
            testDescription
        );

        // Espera a confirmação da transação
        console.log(`Transação enviada: ${txResponse.hash}`);
        await txResponse.wait();
        console.log(`Transação confirmada!`);

        // Chama a função queryQualityTestBatch para consultar o lote
        const batchData = await contract.queryQualityTestBatch(batchId);

        // Converte valores BigInt para string antes de exibi-los
        console.log(`Dados do lote de teste recuperados:`);
        console.log(`ID: ${batchData[0]}`);
        console.log(`Min Limit: ${batchData[1].toString()}`); // Converte BigInt para string
        console.log(`Max Limit: ${batchData[2].toString()}`); // Converte BigInt para string
        console.log(`Valores: ${batchData[3].map((val) => val.toString())}`); // Converte array de BigInt para array de strings
        console.log(`Ordem de Produção: ${batchData[4]}`);
        console.log(`Data do Teste: ${batchData[5]}`);
        console.log(`Descrição do Teste: ${batchData[6]}`);
        
        // Converte os resultados das amostras (SampleResult)
        const sampleResults = batchData[7].map((sample) => ({
            value: sample.value.toString(), // Converte BigInt para string
            result: sample.result
        }));
        console.log(`Resultados das Amostras: ${JSON.stringify(sampleResults)}`);
        console.log(`Resultado Final: ${batchData[8]}`);
    } catch (error) {
        console.error(`Erro ao interagir com o contrato: ${error.message}`);
    }
}

// Executa a função
interactWithContract();