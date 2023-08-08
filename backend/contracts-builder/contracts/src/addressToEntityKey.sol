// SPDX-License-Identifier: MIT
pragma solidity >=0.8.0;

function addressToEntityKey(address addr) pure returns (bytes32) {
    return bytes32(uint256(uint160(addr)));
}

function monKey(bytes32 playerKey, uint8 ID) pure returns (bytes32) {
    if (ID == 0) {
        return playerKey ^ bytes32("1");
    }
    if (ID == 1) {
        return playerKey ^ bytes32("2");
    }
    if (ID == 2) {
        return playerKey ^ bytes32("3");
    }
    return bytes32(playerKey);
}
