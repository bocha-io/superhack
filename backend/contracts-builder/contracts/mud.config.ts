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
    MatchResult: {
      dataStruct: false,
      schema: {
        winner: "bytes32",
        loser: "bytes32",
      },
    },

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
    MonType: [
      "Cobarett",
      "Flarezael",
      "Firomenis",
      "Baobaffe",
      "Howliage",
      "Sunnydra",
      "Tobishimi",
      "Mobiusk",
      "Ramon",
    ],
    StatusType: ["Walking", "Fighting"],
    ActionType: ["Attack", "Swap", "Surrender"],
    ElementType: ["Fire", "Water", "Grass", "Normal"],
    AttackType: [
      // Fire
      "LavaPlume",
      "FireLash",
      "FieryDance",
      // Grass
      "LeafTornado",
      "RazorLeaf",
      "SolarBlade",
      // Water
      "AquaTail",
      "BubbleBeam",
      "HydroVortex",
      // Normal
      "Bite",
      "Tackle",
      "CrushGrip",
      "Flail",
    ],
  },
  modules: [],
});
