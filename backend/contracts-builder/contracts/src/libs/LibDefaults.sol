// SPDX-License-Identifier: MIT
pragma solidity >=0.8.0;

import {MonType} from "../codegen/Types.sol";

library LibDefaults {
    function getHp(MonType mon) internal pure returns (int32) {
        if (mon == MonType.Renzok) {
            return 100;
        }
        if (mon == MonType.Ramon) {
            return 6;
        }
        if (mon == MonType.Hanchon) {
            return 50;
        }
        return 0;
    }
}
