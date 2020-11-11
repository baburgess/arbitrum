// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.5.17;

contract Node {
    bytes32 public assertionHash;
    uint256 public prev;
    uint256 public proposedBlock;
    uint256 public deadlineBlock;
    uint256 public stakerCount;
    mapping(address => bool) public stakers;

    address rollup;

    modifier onlyRollup {
        require(msg.sender == rollup, "ROLLUP_ONLY");
        _;
    }

    constructor(
        bytes32 _assertionHash,
        uint256 _prev,
        uint256 _proposedBlock,
        uint256 _deadlineBlock,
        uint256 _stakerCount
    ) public {
        assertionHash = _assertionHash;
        prev = _prev;
        proposedBlock = _proposedBlock;
        deadlineBlock = _deadlineBlock;
        stakerCount = _stakerCount;
    }

    function confirmValid(uint256 totalStakerCount) external onlyRollup {
        // Verify the block's deadline has passed
        require(deadlineBlock <= block.number);

        // All non-zombie stakers are staked on this node, and no zombie stakers are staked here
        require(stakerCount == totalStakerCount);

        // There is at least one non-zombie staker
        require(totalStakerCount > 0);

        selfdestruct(msg.sender);
    }

    function confirmInvalid() external onlyRollup {
        // Verify the block's deadline has passed
        require(deadlineBlock <= block.number);

        // Verify that no staker is staked on this node
        require(stakerCount == 0);

        selfdestruct(msg.sender);
    }

    function confirmOutOfOrder(uint256 latestConfirmed) external onlyRollup {
        require(prev != latestConfirmed);

        selfdestruct(msg.sender);
    }

    function addStaker(address staker) external onlyRollup {
        require(!stakers[staker], "ALREADY_STAKED");
        stakers[staker] = true;
        stakerCount++;
    }

    function removeStaker(address staker) external onlyRollup {
        require(stakers[staker], "NOT_STAKED");
        stakers[staker] = false;
        stakerCount--;
    }
}