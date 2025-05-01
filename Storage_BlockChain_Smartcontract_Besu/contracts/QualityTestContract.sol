// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract QualityTestContract {
    // Estrutura para representar o resultado de uma amostra individual
    struct SampleResult {
        int256 value;
        string result; // "OK" ou "NOK"
    }

    // Estrutura para representar um lote de teste de qualidade
    struct QualityTestBatch {
        string id;
        int256 minLimit;
        int256 maxLimit;
        int256[] values;
        string productionOrder;
        string testDate;
        string testDescription;
        SampleResult[] sampleResults;
        string finalResult; // "OK" ou "NOK"
    }

    // Mapeamento para armazenar os lotes de teste de qualidade
    mapping(string => QualityTestBatch) private qualityTestBatches;

    // Evento para notificar a criação de um novo lote de teste
    event QualityTestBatchCreated(string id, string finalResult);

    // Função para realizar o teste de qualidade em um lote
    function performBatchQualityTest(
        string memory id,
        int256 minLimit,
        int256 maxLimit,
        int256[] memory values,
        string memory productionOrder,
        string memory testDate,
        string memory testDescription
    ) public {
        // Verifica se o ID já existe
        require(bytes(qualityTestBatches[id].id).length == 0, "Batch with this ID already exists");

        // Criação manual da array de SampleResult
        SampleResult[] memory sampleResults = new SampleResult[](values.length);
        string memory finalResult = "OK";

        // Valida cada valor e determina o resultado
        for (uint256 i = 0; i < values.length; i++) {
            string memory result = "OK";

            if (values[i] < minLimit || values[i] > maxLimit) {
                result = "NOK";
                finalResult = "NOK";
            }

            sampleResults[i] = SampleResult(values[i], result);
        }

        // Armazena o lote de teste no mapeamento
        QualityTestBatch storage batch = qualityTestBatches[id];
        batch.id = id;
        batch.minLimit = minLimit;
        batch.maxLimit = maxLimit;
        batch.values = values;
        batch.productionOrder = productionOrder;
        batch.testDate = testDate;
        batch.testDescription = testDescription;
        batch.finalResult = finalResult;

        // Copiando manualmente os elementos de sampleResults para o storage
        for (uint256 i = 0; i < sampleResults.length; i++) {
            batch.sampleResults.push(sampleResults[i]);
        }

        // Emite o evento de criação do lote
        emit QualityTestBatchCreated(id, finalResult);
    }

    // Função para consultar um lote de teste de qualidade pelo ID
    function queryQualityTestBatch(string memory id) public view returns (
        string memory,
        int256,
        int256,
        int256[] memory,
        string memory,
        string memory,
        string memory,
        SampleResult[] memory,
        string memory
    ) {
        QualityTestBatch storage batch = qualityTestBatches[id];
        require(bytes(batch.id).length != 0, "Test batch with this ID does not exist");

        return (
            batch.id,
            batch.minLimit,
            batch.maxLimit,
            batch.values,
            batch.productionOrder,
            batch.testDate,
            batch.testDescription,
            batch.sampleResults,
            batch.finalResult
        );
    }
}