const fs = require('fs');
const { Circuit } = require('snarkjs');
const { MerkleTree } = require('merkletreejs');
const keccak256 = require('keccak256');

async function testCircuit() {
  const circuitDef = JSON.parse(fs.readFileSync('circuit.json'));
  const circuit = new Circuit(circuitDef);

  // Sample Merkle tree data
  const treeDepth = 4;
  const eligibleVoters = [
    '0x742d35Cc6634C0532925a3b844Bc454e4438f44e',
    '0x742d35Cc6634C0532925a9b844Bc454e4438f44d',
    '0x742d35Cc6634C0532925a2b844Bc454e4438f44c',
    '0x742d35Cc6634C0532925a1b844Bc454e4438f44b'
  ];
  const leaves = eligibleVoters.map(voter => keccak256(voter));
  const tree = new MerkleTree(leaves, keccak256, {hashLeaves: true, sort: true});

  // Choose a leaf to test
  const leafIndex = 0;
  const leafValue = leaves[leafIndex];
  const leaf = keccak256(leafValue);

  // Calculate Merkle proof for the leaf
  const proof = tree.getProof(leaf, leafIndex);

  // Construct the circuit input from proof data
  const input = {
    root: tree.getRoot().toString('hex'),
    leaf: leaf.toString('hex'),
    path: proof.map(p => p.data.toString('hex')),
    pathIndex: proof.map(p => p.index)
  };

  // Compute witness on the circuit
  const witness = circuit.calculateWitness(input);

  // Check if the output signal is valid
  if (circuit.getSignal(witness, 'main.valid') == 1) {
    console.log('TEST PASSED: Valid Merkle proof');
  } else {
    console.log('TEST FAILED: Invalid Merkle proof');
  }
}

testCircuit()
  .catch((err) => console.error('Circuit test error:', err));
