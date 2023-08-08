import { mudConfig } from "@latticexyz/world/register";

export default mudConfig({
  systems: {},
  tables: {
    Player: "bool",
    Status: "StatusType",

    Position: {
      dataStruct: false,
      schema: {
        x: "int32",
        y: "int32",
      },
    },

    Match: "bool",
    PlayerOne: "bytes32",
    PlayerTwo: "bytes32",
    PlayerOneCurrentMon: "bytes32",
    PlayerTwoCurrentMon: "bytes32",

    PlayerFirstMon: "bytes32",
    PlayerSecondMon: "bytes32",
    PlayerThirdMon: "bytes32",
    Mon: "bool",
    MonSpecie: "MonType",
    MonHp: "int32",

    // Inventory
    InventoryFirstMon: "MonType",
    InventorySecondMon: "MonType",
    InventoryThirdMon: "MonType",
  },
  enums: {
    MonType: ["Ramon", "Renzok", "Hanchon"],
    StatusType: ["Walking", "Fighting"],
    ActionType: ["Attack", "Swap"],
  },
  modules: [],
});
