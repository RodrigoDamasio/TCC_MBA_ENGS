const fs = require("fs");
const path = require("path");

async function extractABI() {
    const artifactsPath = path.join(__dirname, "./QualityTestContract.json");
    const artifact = JSON.parse(fs.readFileSync(artifactsPath, "utf8"));

    // Extrai o ABI
    const abi = artifact.abi;

    // Salva o ABI em um arquivo separado, se necessário
    fs.writeFileSync("QualityTestContract.abi", JSON.stringify(abi, null, 2));

    console.log("ABI extraído com sucesso e salvo em QualityTestContract.abi");
}

extractABI();