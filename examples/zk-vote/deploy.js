const Web3 = require("web3");
const { abi, bytecode } = require("./ContractVoting.json");

async function main() {
  const web3 = new Web3("http://localhost:8545");
  const accounts = await web3.eth.getAccounts();
  const root = "0x..."; // Replace with a calculated Merkle Tree root from before

  // Deploy contract
  const contract = await new web3.eth.Contract(abi)
    .deploy({ data: bytecode, arguments: [root] })
    .send({ from: accounts[0], gas: 4500000 });

  console.log("Contract deployed at address:", contract.options.address);

  const proof = [
    "0x...", "0x...", // Replace with the valid proof from zkp.js
  ];
  const publicSignals = [
    "0x...", // Replace with the valid public signals from zkp.js
  ];

  // Perform a sample vote
  await contract.methods
    .vote(proof, publicSignals)
    .send({ from: accounts[0], gas: 4500000 });
}

main().catch(console.error);
