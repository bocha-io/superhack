// SPDX-License-Identifier: MIT
pragma solidity >=0.8.0;

// Core
import {System} from "@latticexyz/world/src/System.sol";

import {Player} from "../codegen/tables/Player.sol";
import {Position} from "../codegen/tables/Position.sol";

import {addressToEntityKey} from "../addressToEntityKey.sol";

contract MoveSystem is System {
    function Move(int32 newX, int32 newY) public {
        bytes32 senderKey = addressToEntityKey(_msgSender());

        // Make sure that the player is registered
        require(Player.get(senderKey) == true, "wallet is not registered");

        // Check that the X and Y are valid values
        require(newX >= 0 && newY >= 0, "invalid X and Y");
        Position.set(senderKey, newX, newY);
    }
}
