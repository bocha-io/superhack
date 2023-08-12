// SPDX-License-Identifier: MIT
pragma solidity >=0.8.0;

// Core
import {System} from "@latticexyz/world/src/System.sol";

// Battle
import {Match} from "../codegen/tables/Match.sol";
import {Status} from "../codegen/tables/Status.sol";
import {MatchResult} from "../codegen/tables/MatchResult.sol";
import {PlayerOne} from "../codegen/tables/PlayerOne.sol";
import {PlayerTwo} from "../codegen/tables/PlayerTwo.sol";

import {MonSpecie} from "../codegen/tables/MonSpecie.sol";
import {MonHp} from "../codegen/tables/MonHp.sol";

import {PlayerOneCurrentMon} from "../codegen/tables/PlayerOneCurrentMon.sol";
import {PlayerTwoCurrentMon} from "../codegen/tables/PlayerTwoCurrentMon.sol";

import {PlayerFirstMon} from "../codegen/tables/PlayerFirstMon.sol";
import {PlayerSecondMon} from "../codegen/tables/PlayerSecondMon.sol";
import {PlayerThirdMon} from "../codegen/tables/PlayerThirdMon.sol";

// Libs
import {ActionType, ElementType, StatusType} from "../codegen/Types.sol";
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

    function attack(bytes32 attackingMon, bytes32 attackedMon, uint8 pos) internal returns (int32) {
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

    function fight(bytes32 matchID, bytes32 monOne, bytes32 monTwo, uint8 posOne, uint8 posTwo)
        internal
        returns (bool)
    {
        // PlayerOne stats
        int32 playerOneSpeed = 0;
        {
            (int32 speedMonPlayerOne,) = LibDefaults.getMonSpeedAndType(MonSpecie.get(monOne));
            (, int32 atkSpeedMonOne,) = LibDefaults.getAttack(MonSpecie.get(monOne), posOne);
            playerOneSpeed = speedMonPlayerOne + atkSpeedMonOne;
        }

        int32 playerTwoSpeed = 0;
        {
            (int32 speedMonPlayerTwo,) = LibDefaults.getMonSpeedAndType(MonSpecie.get(monTwo));
            (, int32 atkSpeedMonTwo,) = LibDefaults.getAttack(MonSpecie.get(monTwo), posTwo);
            playerTwoSpeed = speedMonPlayerTwo + atkSpeedMonTwo;
        }

        if (playerOneSpeed >= playerTwoSpeed) {
            // Player one goes first
            if (attack(monOne, monTwo, posOne) == 0) {
                // Check if it was the last mon
                if (checkIfGameHasEnded(PlayerTwo.get(matchID), monTwo)) {
                    // Surrender game for player two
                    endGame(matchID, PlayerOne.get(matchID), PlayerTwo.get(matchID));
                    return true;
                }
            } else {
                // The player two mon is alive
                if (attack(monTwo, monOne, posTwo) == 0) {
                    // Check if it was the last mon
                    if (checkIfGameHasEnded(PlayerOne.get(matchID), monOne)) {
                        // Surrender game for player one
                        endGame(matchID, PlayerTwo.get(matchID), PlayerOne.get(matchID));
                        return true;
                    }
                }
            }
        } else {
            // Player two goes first
            if (attack(monTwo, monOne, posTwo) == 0) {
                // Check if it was the last mon
                if (checkIfGameHasEnded(PlayerOne.get(matchID), monOne)) {
                    // Surrender game for player one
                    endGame(matchID, PlayerTwo.get(matchID), PlayerOne.get(matchID));
                    return true;
                }
            } else {
                // Player one mon is alive
                if (attack(monOne, monTwo, posOne) == 0) {
                    // Check if it was the last mon
                    if (checkIfGameHasEnded(PlayerTwo.get(matchID), monTwo)) {
                        // Surrender game for player two
                        endGame(matchID, PlayerOne.get(matchID), PlayerTwo.get(matchID));
                        return true;
                    }
                }
            }
        }

        // Game continues
        return false;
    }

    // We need pass the mon that was dead to be supported by the golang predictions.
    // Currently there is no support to read stores updated in the same transaction
    function checkIfGameHasEnded(bytes32 playerID, bytes32 monDead) internal view returns (bool) {
        int32 res = 0;
        {
            bytes32 firstMon = PlayerFirstMon.get(playerID);
            if (firstMon != monDead) {
                res = res + MonHp.get(firstMon);
            }
        }

        {
            bytes32 secondMon = PlayerSecondMon.get(playerID);
            if (secondMon != monDead) {
                res = res + MonHp.get(secondMon);
            }
        }

        {
            bytes32 thirdMon = PlayerThirdMon.get(playerID);
            if (thirdMon != monDead) {
                res = res + MonHp.get(thirdMon);
            }
        }

        return res == 0;
    }

    function endGame(bytes32 matchID, bytes32 winner, bytes32 loser) internal {
        Match.deleteRecord(matchID);
        PlayerOne.deleteRecord(matchID);
        PlayerTwo.deleteRecord(matchID);
        MatchResult.set(matchID, winner, loser);
        Status.set(winner, StatusType.Walking);
        Status.set(loser, StatusType.Walking);
    }

    address immutable adminAddressBattle = 0x28D6D4078DAA1D192e3854D7BAfF51AE337f4635;

    function Battle(bytes32 matchID, ActionType playerOneAction, uint8 posOne, ActionType playerTwoAction, uint8 posTwo)
        public
    {
        require(_msgSender() == adminAddressBattle, "only admin wallet can create match");

        require(Match.get(matchID), "match is not created");

        // First we swap and then we attack, if both player are attacking we check velocity
        bool p1Executed = false;
        bytes32 p1Mon = PlayerOneCurrentMon.get(matchID);
        bool p2Executed = false;
        bytes32 p2Mon = PlayerTwoCurrentMon.get(matchID);

        // Handle surrender cases
        if (playerOneAction == ActionType.Surrender) {
            endGame(matchID, PlayerTwo.get(matchID), PlayerOne.get(matchID));
            return;
        }

        if (playerTwoAction == ActionType.Surrender) {
            endGame(matchID, PlayerOne.get(matchID), PlayerTwo.get(matchID));
            return;
        }

        // Handle Swaps
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

        // If the user want to send an action with a dead mon, assume surrender
        if (MonHp.get(p1Mon) == 0) {
            // Surrender player one
            endGame(matchID, PlayerTwo.get(matchID), PlayerOne.get(matchID));
            return;
        } else if (MonHp.get(p2Mon) == 0) {
            // Surrender player two
            endGame(matchID, PlayerOne.get(matchID), PlayerTwo.get(matchID));
            return;
        }

        // Handle Attacks
        if (!p1Executed) {
            if (!p2Executed) {
                // fight returns true if the game ended
                if (fight(matchID, p1Mon, p2Mon, posOne, posTwo)) {
                    return;
                }
            } else {
                // Attack with p1
                if (attack(p1Mon, p2Mon, posOne) == 0) {
                    // Check if player two lost the game
                    if (checkIfGameHasEnded(PlayerTwo.get(matchID), p2Mon)) {
                        // Surrender game for player two
                        endGame(matchID, PlayerOne.get(matchID), PlayerTwo.get(matchID));
                        return;
                    }
                }
            }
        } else if (p2Executed == false) {
            // Attack with p2
            if (attack(p2Mon, p1Mon, posTwo) == 0) {
                // Check if player one lost the game
                if (checkIfGameHasEnded(PlayerOne.get(matchID), p1Mon)) {
                    // Surrender game for player one
                    endGame(matchID, PlayerTwo.get(matchID), PlayerOne.get(matchID));
                    return;
                }
            }
        }
    }
}
