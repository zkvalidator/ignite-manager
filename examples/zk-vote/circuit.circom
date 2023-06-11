pragma circom 2.0.0;

include "./node_modules/circomlib/circuits/pedersen.circom";
include "./node_modules/circomlib/circuits/comparators.circom";

template MerkelProofHash() {
    signal input inR;
    signal input inL;
    signal output out;

    component h = Pedersen(128);
    h.in[0] <== inR;
    h.in[1] <== inL;
    for (var i = 2; i < 128; i++) {
        h.in[i] <== 0;
    }
    out <== h.out[0];
}

template ConditionalMux() {
    signal input c;
    signal input t;
    signal input f;
    signal output o;
    
    o <== (c * (t - f)) + f;
}

template MerkleProof() {
    signal input root;
    signal input leaf;
    signal input pathElements[4];
    signal input pathIndices[4];

    component hasher[4];
    component mux[4];
    
    signal temp[4];
    temp[0] <== leaf;

    for (var i = 0; i < 4; i++) {
        hasher[i] = MerkelProofHash();
        mux[i] = ConditionalMux();

        mux[i].c <== pathIndices[i];
        mux[i].t <== temp[i];
        mux[i].f <== pathElements[i];
        hasher[i].inL <== mux[i].o;
        hasher[i].inR <== (temp[i] - mux[i].o) + pathElements[i];

        if (i < 3) {
            temp[i + 1] <== hasher[i].out;
        }
    }

    component isEqual = IsEqual();
    isEqual.in[0] <== hasher[3].out;
    isEqual.in[1] <== root;

    signal output valid;
    valid <== isEqual.out;
}

template Main() {
    signal input root;
    signal input leaf;
    signal input pathElements[4];
    signal input pathIndices[4];

    component merkleProof = MerkleProof();
    merkleProof.root <== root;
    merkleProof.leaf <== leaf;
    merkleProof.pathElements <== pathElements;
    merkleProof.pathIndices <== pathIndices;

    signal output valid;
    valid <== merkleProof.valid;
}

component main = Main();
