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
import {StatusType, MonType, ActionType, ElementType} from "../codegen/Types.sol";
import {monKey} from "../addressToEntityKey.sol";
import {LibDefaults} from "../libs/LibDefaults.sol";

contract BattleSystem is System {
    function getPlayerFightingMon(bytes32 player, uint8 pos) internal view returns (bytes32) {
        if (pos == 0) {
            return PlayerFirstMon.get(player);
        }
        if (pos == 1) {
            return PlayerSecondMon.get(player);
        }
        if (pos == 2) {
            return PlayerThirdMon.get(player);
        }
        require(false, "invalid pos");
        return bytes32(0);
    }

    function Attack(bytes32 attackingMon, bytes32 attackedMon, uint8 pos) internal returns (int32) {
        (int32 atkDmg,, ElementType atkElement) = LibDefaults.getAttack(MonSpecie.get(attackingMon), pos);
        (, ElementType attackedMonElement) = LibDefaults.getMonSpeedAndType(MonSpecie.get(attackedMon));

        int32 dmg = LibDefaults.getAttackDamage(atkDmg, atkElement, attackedMonElement);
        int32 hp = MonHp.get(attackedMon);

        if (dmg > hp) {
            MonHp.set(attackedMon, 0);
            return 0;
        } else {
            MonHp.set(attackedMon, hp - dmg);
            return hp - dmg;
        }
    }

    function BothAttacks(bytes32 monOne, bytes32 monTwo, uint8 posOne, uint8 posTwo) internal {
        // PlayerOne stats
        int32 playerOneSpeed = 0;
        {
            (int32 speedMonPlayerOne,) = LibDefaults.getMonSpeedAndType(MonSpecie.get(monOne));
            (int32 atkSpeed,,) = LibDefaults.getAttack(MonSpecie.get(monOne), posOne);
            playerOneSpeed = speedMonPlayerOne + atkSpeed;
        }

        int32 playerTwoSpeed = 0;
        {
            (int32 speedMonPlayerTwo,) = LibDefaults.getMonSpeedAndType(MonSpecie.get(monTwo));
            (int32 atkSpeed,,) = LibDefaults.getAttack(MonSpecie.get(monTwo), posTwo);
            playerTwoSpeed = speedMonPlayerTwo + atkSpeed;
        }

        if (playerOneSpeed >= playerTwoSpeed) {
            // Check if the mon is dead after attack
            if (Attack(monOne, monTwo, posOne) > 0) {
                Attack(monTwo, monOne, posTwo);
            }
        } else {
            // Check if the mon is dead after attack
            if (Attack(monTwo, monOne, posTwo) > 0) {
                Attack(monOne, monTwo, posOne);
            }
        }
    }

    function Battle(bytes32 matchID, ActionType playerOneAction, uint8 posOne, ActionType playerTwoAction, uint8 posTwo)
        public
    {
        // TODO: allow only game admin keys to access to this function!
        require(Match.get(matchID), "match is not created");

        // First we swap and then we attack, if both player are attacking we check velocity
        bool p1Executed = false;
        bytes32 p1Mon = PlayerOneCurrentMon.get(matchID);
        bool p2Executed = false;
        bytes32 p2Mon = PlayerTwoCurrentMon.get(matchID);

        if (playerOneAction == ActionType.Swap) {
            p1Executed = true;
            // Do Swap
            p1Mon = getPlayerFightingMon(PlayerOne.get(matchID), posOne);
            PlayerOneCurrentMon.set(matchID, p1Mon);
        }

        if (playerTwoAction == ActionType.Swap) {
            p2Executed = true;
            // Do Swap
            p2Mon = getPlayerFightingMon(PlayerTwo.get(matchID), posTwo);
            PlayerTwoCurrentMon.set(matchID, p2Mon);
        }

        if (!p1Executed) {
            if (!p2Executed) {
                // Check velocity
                BothAttacks(p1Mon, p2Mon, posOne, posTwo);
            } else {
                // Attack with p1
                Attack(p1Mon, p2Mon, posOne);
            }
        } else if (p2Executed == false) {
            // Attack with p2
            Attack(p2Mon, p1Mon, posTwo);
        }

        // TODO: check for gameover
    }
}
