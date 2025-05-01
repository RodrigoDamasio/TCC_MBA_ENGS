async function main() {
  const [deployer] = await ethers.getSigners();
  console.log("Deploying contracts with the account:", deployer.address);

  const QualityTestContract = await ethers.getContractFactory("QualityTestContract");
  console.log("Contract factory created");

  const qualityTestContract = await QualityTestContract.deploy();
  console.log("Contract deployment transaction sent");

  //await simpleStorage.deployed();
  //console.log("Contract deployed");

  console.log("QualityTestContract deployed to:", QualityTestContract.address);
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });