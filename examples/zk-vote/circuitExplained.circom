pragma circom 2.0.0;

// Circom is a domain-specific language for defining circuits used in
// zero-knowledge proofs (ZKP). ZKPs enable a prover to convince a verifier
// about the validity of a statement without revealing any additional
// information.

// In this example, we will create a circuit for verifying the authenticity of
// a Merkle proof, which is a proof that a certain data (or its hash) is
// included in a Merkle tree. 

// This program consists of several "templates" that define custom components
// for the circuits, performing various operations such as hashing, muxing,
// and comparing values. Finally, there is a "main" template that combines
// these components to verify the Merkle proof.

// The input elements for this circuit include the Merkle tree root hash,
// the leaf data (or its hash), sibling nodes in the proof path,
// and the indices indicating left or right child for each level of the tree.
// The output of the circuit is a "valid" signal, which will be 1 if the
// Merkle proof is valid, otherwise 0.

// First, we include the necessary circuits from the circomlib library,
// which contains several prebuilt circuits for common operations in ZKPs.
include "./node_modules/circomlib/circuits/pedersen.circom";
include "./node_modules/circomlib/circuits/comparators.circom";

// The MerkleProofHash template will be used for computing the hash of two
// Merkle tree nodes using the Pedersen hash function (a cryptographic hash
// function commonly used in zk-SNARKs circuits). This template expects the
// right and left child nodes' values as inputs and outputs the resulting hash.
template MerkelProofHash() {
    // Inputs for the right and left child nodes.
    signal input inR;
    signal input inL;

    // Output signal for the resulting hash.
    signal output out;

    // Instantiate the Pedersen hash function for 128 bits (two 64-bit inputs).
    component h = Pedersen(128);
    h.in[0] <== inR;
    h.in[1] <== inL;

    // Set unused inputs for the hash function to zero.
    for (var i = 2; i < 128; i++) {
        h.in[i] <== 0;
    }

    // Assign the output signal to the hash of the two child nodes.
    out <== h.out[0];
}

// The ConditionalMux template implements a conditional multiplexer (mux)
// operation. A mux selects one value from several inputs based on a control
// signal. In this case, the mux operator serves as an if-then-else chooser:
// for a control signal c, it selects the "true" input value t when c=1,
// and the "false" input value f when c=0.
template ConditionalMux() {
    // The control input c serves as a boolean flag: 1 for true, 0 for false.
    signal input c;

    // "True" and "false" input values.
    signal input t;
    signal input f;
    
    // The selected input value based on the control signal.
    signal output o;

    // Compute the output value based on the control signal.
    // When c is 1 (true), the output o = (1 * (t - f)) + f = t.
    // When c is 0 (false), the output o = (0 * (t - f)) + f = f.
    o <== (c * (t - f)) + f;
}

// The MerkleProof template is the core component for verifying a Merkle proof,
// combining the hasher components and the multiplexer components for the
// proof path. The idea is to hash the elements along the proof path, calculate
// the root hash, and check if that matches the expected root hash. If the
// hashes match, the leaf data is considered a valid part of the tree, and the
// "valid" output signal is asserted (set to 1).
template MerkleProof() {
    // Inputs for the Merkle tree root hash, leaf data (or hash),
    // the sibling nodes in the proof path, and the indices indicating
    // left or right child for each level of the tree.
    signal input root;
    signal input leaf;
    signal input pathElements[4];
    signal input pathIndices[4];

    // Instantiate hasher components and mux to process the Merkle proof.
    component hasher[4];
    component mux[4];

    // Temporary signals used in the hashing process.
    signal temp[4];
    temp[0] <== leaf;

    // Iterate through each level of the tree, hashing the nodes and updating
    // the mux control signal as needed.
    for (var i = 0; i < 4; i++) {
        hasher[i] = MerkelProofHash();
        mux[i] = ConditionalMux();

        // Set the mux control signal to the pathIndices value at the current level.
        mux[i].c <== pathIndices[i];
        mux[i].t <== temp[i];
        mux[i].f <== pathElements[i];

        // Choose the appropriate hash value based on the pathIndices value for this level.
        hasher[i].inL <== mux[i].o;
        hasher[i].inR <== (temp[i] - mux[i].o) + pathElements[i];

        // Update the temp signal for the next level, if not at the last level.
        if (i < 3) {
            temp[i + 1] <== hasher[i].out;
        }
    }

    // Instantiate an IsEqual component to compare the calculated root hash
    // with the input root hash.
    component isEqual = IsEqual();
    isEqual.in[0] <== hasher[3].out;
    isEqual.in[1] <== root;

    // The output signal named "valid" will be 1 if the root hashes match,
    // and 0 otherwise.
    signal output valid;
    valid <== isEqual.out;
}

// The Main template combines the components and wires them together as a
// top-level circuit for verifying the Merkle proof. The main component
// instantiates a MerkleProof component, connecting the main component's
// inputs and outputs to the appropriate signals.
template Main() {
    // Input signals for the Merkle tree root hash, leaf data (or its hash),
    // the sibling nodes in the proof path, and the indices indicating
    // left or right child for each level of the tree.
    signal input root;
    signal input leaf;
    signal input pathElements[4];
    signal input pathIndices[4];

    // Instantiate a MerkleProof component to perform the proof validation.
    component merkleProof = MerkleProof();
    merkleProof.root <== root;
    merkleProof.leaf <== leaf;
    merkleProof.pathElements <== pathElements;
    merkleProof.pathIndices <== pathIndices;

    // Propagate the "valid" signal from the Merkle proof to the main component's output.
    signal output valid;
    valid <== merkleProof.valid;
}

// Finally, instantiate the main Merkle proof verification component to use in
// a zero-knowledge proof (ZKP) circuit. This main component acquires the proper
// input signals, processes them using the MerkleProof template, and generates
// a "valid" output signal, indicating if the Merkle proof is considered valid.
component main = Main();
