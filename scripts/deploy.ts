import { ethers } from "hardhat";
import { config } from "dotenv";

async function main() {
  config();

  const Gateway = await ethers.getContractFactory("Gateway");
  const gateway = await Gateway.deploy();

  await gateway.deployed();

  console.log(`Gateway deployed to ${gateway.address}`);
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
