require("@nomicfoundation/hardhat-toolbox");

module.exports = {
  solidity: "0.8.20",
  networks: {
    ProdGanache: {
      url: "http://127.0.0.1:8546",
      accounts:   ["0x3a5b5ec5bdbca3a6e7e114d3f711ef458be5a6c88dc1d90ced7f487e570281f3"]
    },
    besu: {
      url: "http://127.0.0.1:8550",
      accounts: ["0xf8060dd8211fd80d62023e74aff54d870e13835c630865d6250761394efddc98"]
      
    },
    sepolia: {
      url: "https://sepolia.infura.io/v3/2726b48385b2497a84cdecacb028e6b0",
      //accounts: ["0x0000000000000000000000000000000000000000000000003239313131393934"]
      
    },
    hardhat:{}
  },
};