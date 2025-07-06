// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/utils/Counters.sol";

contract QualityTestContract is ERC721 {
    using Counters for Counters.Counter;
    Counters.Counter private _tokenIds;

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
        string imagesJson; // Campo opcional para imagens em JSON
    }

    // Mapeamento do ID do token para os dados do batch
    mapping(uint256 => QualityTestBatch) private qualityTestBatches;
    // Mapeamento do id string para o tokenId ERC721 (para compatibilidade)
    mapping(string => uint256) private batchIdToTokenId;

    // Evento para notificar a criação de um novo lote de teste
    event QualityTestBatchCreated(string id, uint256 tokenId, string finalResult);

    constructor() ERC721("QualityTestBatchNFT", "QTB") {}

    // Função para realizar o teste de qualidade em um lote (mint NFT)
    function performBatchQualityTest(
        string memory id,
        int256 minLimit,
        int256 maxLimit,
        int256[] memory values,
        string memory productionOrder,
        string memory testDate,
        string memory testDescription,
        string memory imagesJson // Novo campo opcional
    ) public {
        require(batchIdToTokenId[id] == 0, "Batch with this ID already exists");

        SampleResult[] memory sampleResults = new SampleResult[](values.length);
        string memory finalResult = "OK";

        for (uint256 i = 0; i < values.length; i++) {
            string memory result = "OK";
            if (values[i] < minLimit || values[i] > maxLimit) {
                result = "NOK";
                finalResult = "NOK";
            }
            sampleResults[i] = SampleResult(values[i], result);
        }

        _tokenIds.increment();
        uint256 newTokenId = _tokenIds.current();

        // Mint do NFT para o remetente
        _mint(msg.sender, newTokenId);

        // Armazena o batch no mapeamento
        QualityTestBatch storage batch = qualityTestBatches[newTokenId];
        batch.id = id;
        batch.minLimit = minLimit;
        batch.maxLimit = maxLimit;
        batch.values = values;
        batch.productionOrder = productionOrder;
        batch.testDate = testDate;
        batch.testDescription = testDescription;
        batch.finalResult = finalResult;
        batch.imagesJson = imagesJson;

        for (uint256 i = 0; i < sampleResults.length; i++) {
            batch.sampleResults.push(sampleResults[i]);
        }

        batchIdToTokenId[id] = newTokenId; // Relaciona o id string ao tokenId

        emit QualityTestBatchCreated(id, newTokenId, finalResult);
    }

    // Função para consultar um lote de teste de qualidade pelo ID string
    function queryQualityTestBatch(string memory id) public view returns (
        string memory,
        int256,
        int256,
        int256[] memory,
        string memory,
        string memory,
        string memory,
        SampleResult[] memory,
        string memory,
        string memory // imagesJson
    ) {
        require(batchIdToTokenId[id] != 0, "Test batch with this ID does not exist");
        uint256 tokenId = batchIdToTokenId[id];
        QualityTestBatch storage batch = qualityTestBatches[tokenId];

        return (
            batch.id,
            batch.minLimit,
            batch.maxLimit,
            batch.values,
            batch.productionOrder,
            batch.testDate,
            batch.testDescription,
            batch.sampleResults,
            batch.finalResult,
            batch.imagesJson
        );
    }

    // Consulta por tokenId (caso queira usar a lógica padrão ERC721)
    function queryQualityTestBatchByToken(uint256 tokenId) public view returns (
        string memory,
        int256,
        int256,
        int256[] memory,
        string memory,
        string memory,
        string memory,
        SampleResult[] memory,
        string memory,
        string memory
    ) {
         _requireOwned(tokenId); // lança exceção se não existir
        QualityTestBatch storage batch = qualityTestBatches[tokenId];

        return (
            batch.id,
            batch.minLimit,
            batch.maxLimit,
            batch.values,
            batch.productionOrder,
            batch.testDate,
            batch.testDescription,
            batch.sampleResults,
            batch.finalResult,
            batch.imagesJson
        );
    }
}