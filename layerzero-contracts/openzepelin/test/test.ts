import { expect } from "chai";
import { ethers } from "hardhat";

describe("BochaCoin", function () {
  it("Test contract", async function () {
    const ContractFactory = await ethers.getContractFactory("BochaCoin");

    const instance = await ContractFactory.deploy();
    await instance.waitForDeployment();

    expect(await instance.name()).to.equal("BochaCoin");
  });
});
