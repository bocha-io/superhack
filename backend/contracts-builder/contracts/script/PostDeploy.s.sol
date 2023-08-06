// SPDX-License-Identifier: MIT
pragma solidity >=0.8.0;

import {Script} from "forge-std/Script.sol";
import {IWorld} from "../src/codegen/world/IWorld.sol";

contract PostDeploy is Script {
    function run(address worldAddress) external {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        IWorld world = IWorld(worldAddress);
        vm.startBroadcast(deployerPrivateKey);
        vm.stopBroadcast();
    }
}
