import { mudConfig } from "@latticexyz/world/register";

export default mudConfig({
  systems: {},
  tables: {
    Position: {
      dataStruct: false,
      schema: {
        x: "int32",
        y: "int32",
      },
    },
  },
  enums: {},
  modules: [],
});
