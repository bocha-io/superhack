// SPDX-License-Identifier: MIT
pragma solidity >=0.8.0;

// Core
import {System} from "@latticexyz/world/src/System.sol";

// Player options
import {Player} from "../codegen/tables/Player.sol";
import {Status} from "../codegen/tables/Status.sol";

// Battle
import {Match} from "../codegen/tables/Match.sol";
import {PlayerOne} from "../codegen/tables/PlayerOne.sol";
import {PlayerTwo} from "../codegen/tables/PlayerTwo.sol";

import {Mon} from "../codegen/tables/Mon.sol";
import {MonSpecie} from "../codegen/tables/MonSpecie.sol";
import {MonHp} from "../codegen/tables/MonHp.sol";

import {PlayerOneCurrentMon} from "../codegen/tables/PlayerOneCurrentMon.sol";
import {PlayerTwoCurrentMon} from "../codegen/tables/PlayerTwoCurrentMon.sol";

import {PlayerFirstMon} from "../codegen/tables/PlayerFirstMon.sol";
import {PlayerSecondMon} from "../codegen/tables/PlayerSecondMon.sol";
import {PlayerThirdMon} from "../codegen/tables/PlayerThirdMon.sol";

// Inventory
import {InventoryFirstMon} from "../codegen/tables/InventoryFirstMon.sol";
import {InventorySecondMon} from "../codegen/tables/InventorySecondMon.sol";
import {InventoryThirdMon} from "../codegen/tables/InventoryThirdMon.sol";

// Libs
import {StatusType, MonType} from "../codegen/Types.sol";
import {monKey} from "../addressToEntityKey.sol";
import {LibDefaults} from "../libs/LibDefaults.sol";

contract CreateMatchSystem is System {
    function getMonFromInventory(bytes32 player, uint8 pos) internal view returns (MonType) {
        if (pos == 0) {
            return InventoryFirstMon.get(player);
        }
        if (pos == 1) {
            return InventorySecondMon.get(player);
        }
        if (pos == 2) {
            return InventoryThirdMon.get(player);
        }
        require(false, "invalid pos");
        return MonType.Ramon;
    }

    function setPlayerCurrentMon(bytes32 gameKey, uint8 player, bytes32 monID) internal {
        if (player == 0) {
            PlayerOneCurrentMon.set(gameKey, monID);
            return;
        }
        if (player == 1) {
            PlayerTwoCurrentMon.set(gameKey, monID);
            return;
        }
        require(false, "invalid player");
    }

    function setGameMon(bytes32 player, uint8 pos, bytes32 monID) internal {
        if (pos == 0) {
            PlayerFirstMon.set(player, monID);
            return;
        }
        if (pos == 1) {
            PlayerSecondMon.set(player, monID);
            return;
        }
        if (pos == 2) {
            PlayerThirdMon.set(player, monID);
            return;
        }
        require(false, "invalid pos");
    }

    address immutable adminAddressCreateMatch = 0x28D6D4078DAA1D192e3854D7BAfF51AE337f4635;

    function CreateMatch(bytes32 playerA, bytes32 playerB) public {
        require(_msgSender() == adminAddressCreateMatch, "only admin wallet can create match");

        // Make sure that the player is registered
        require(Player.get(playerA) == true, "player a is not registered");
        require(Player.get(playerB) == true, "player b is not registered");

        // The user can play one game at the time
        require(Status.get(playerA) == StatusType.Walking, "player a is already in a match");
        require(Status.get(playerB) == StatusType.Walking, "player b is already in a match");

        // Game key will be the same as playerA
        bytes32 gameKey = playerA;

        // CreateMatch
        Match.set(gameKey, true);
        PlayerOne.set(gameKey, playerA);
        PlayerTwo.set(gameKey, playerB);

        Status.set(playerA, StatusType.Fighting);
        Status.set(playerB, StatusType.Fighting);

        // Create Mons
        for (uint8 player = 0; player < 2; player++) {
            bytes32 playerKey = playerA;
            if (player == 1) {
                playerKey = playerB;
            }
            for (uint8 i = 0; i < 3; i++) {
                bytes32 key = monKey(playerKey, i);
                Mon.set(key, true);
                MonSpecie.set(key, getMonFromInventory(playerKey, i));
                MonHp.set(key, LibDefaults.getHp(getMonFromInventory(playerKey, i)));
                setGameMon(playerKey, i, key);
                if (i == 0) {
                    setPlayerCurrentMon(gameKey, player, key);
                }
            }
        }
    }
}
