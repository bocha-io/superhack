// SPDX-License-Identifier: MIT
pragma solidity >=0.8.0;

function addressToEntityKey(address addr) pure returns (bytes32) {
    return bytes32(uint256(uint160(addr)));
}

function monKey(bytes32 playerKey, uint8 ID) pure returns (bytes32) {
    return playerKey ^ bytes32(uint256(ID));
}

// monKey := strings.Replace(playerKey, "0x000000000000000000000000", "0x100000000000000000000000", 1)
