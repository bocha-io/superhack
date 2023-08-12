// SPDX-License-Identifier: MIT
pragma solidity >=0.8.0;

import {System} from "@latticexyz/world/src/System.sol";
import {Player} from "../codegen/tables/Player.sol";

import {InventoryFirstMon} from "../codegen/tables/InventoryFirstMon.sol";
import {InventorySecondMon} from "../codegen/tables/InventorySecondMon.sol";
import {InventoryThirdMon} from "../codegen/tables/InventoryThirdMon.sol";

import {MonType} from "../codegen/Types.sol";
import {addressToEntityKey} from "../addressToEntityKey.sol";

contract SwapMonSystem is System {
    function SwapMon(MonType mon, uint8 pos) public {
        // Add 0 padding to the address
        bytes32 senderKey = addressToEntityKey(_msgSender());

        require(Player.get(senderKey), "user is not registered");
        require(pos == 0 || pos == 1 || pos == 2, "invalid pos");

        // Init the inventary
        if (pos == 0) {
            InventoryFirstMon.set(senderKey, mon);
        } else if (pos == 1) {
            InventorySecondMon.set(senderKey, mon);
        } else if (pos == 2) {
            InventoryThirdMon.set(senderKey, mon);
        }
    }
}
