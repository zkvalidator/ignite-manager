pragma solidity ^0.8.0;

contract Voting {
  bytes32 public root;
  mapping(address => bool) public voted;
  uint256 public yesVotes;
  uint256 public noVotes;

  event VoteRegistered(address indexed user, bool indexed vote);

  constructor(bytes32 _root) {
    root = _root;
  }

  function vote(
    uint256[2] calldata a,
    uint256[2][2] calldata b,
    uint256[2] calldata c,
    uint256[8] calldata input
  ) external {
    require(!voted[msg.sender], "Sender already voted");
    
    bool isValid = verifyProof(a, b, c, input);
    require(isValid, "Invalid proof");

    bool userVote = input[7] == 1;
    if (userVote) {
      yesVotes += 1;
    } else {
      noVotes += 1;
    }

    voted[msg.sender] = true;
    emit VoteRegistered(msg.sender, userVote);
  }

  function verifyProof(
    uint256[2] memory a,
    uint256[2][2] memory b,
    uint256[2] memory c,
    uint256[8] memory input
  ) public view returns (bool) {
    uint256[2] memory vk_alpha = [uint256(0x09), uint256(0x0c)];
    // Put your own Verifying Key (as generated below) here
    uint256[] memory vk_gamma_abc = [...] // Don't modify this comment, replace [...] with the generated values
    
    return snarkjs.Verify(vk_alpha, vk_gamma_abc, a, b, c, input);
  }
}
