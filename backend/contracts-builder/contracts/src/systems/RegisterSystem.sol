// SPDX-License-Identifier: MIT
pragma solidity >=0.8.0;

import {System} from "@latticexyz/world/src/System.sol";
import {Player} from "../codegen/tables/Player.sol";

import {InventoryFirstMon} from "../codegen/tables/InventoryFirstMon.sol";
import {InventorySecondMon} from "../codegen/tables/InventorySecondMon.sol";
import {InventoryThirdMon} from "../codegen/tables/InventoryThirdMon.sol";

import {Position} from "../codegen/tables/Position.sol";
import {Status} from "../codegen/tables/Status.sol";

import {StatusType, MonType} from "../codegen/Types.sol";
import {addressToEntityKey} from "../addressToEntityKey.sol";

contract RegisterSystem is System {
    function register() public {
        // Add 0 padding to the address
        bytes32 senderKey = addressToEntityKey(_msgSender());

        // Make sure that we are not already registered
        require(Player.get(senderKey) == false, "wallet already registered");
        Player.set(senderKey, true);

        // Init the inventary
        InventoryFirstMon.set(senderKey, MonType.Hanchon);
        InventorySecondMon.set(senderKey, MonType.Renzok);
        InventoryThirdMon.set(senderKey, MonType.Ramon);

        // Set initial position
        Position.set(senderKey, 0, 0);

        // Set player status
        Status.set(senderKey, StatusType.Walking);
    }
}