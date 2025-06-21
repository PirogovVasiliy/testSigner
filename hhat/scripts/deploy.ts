import { ethers } from "hardhat";
//npx hardhat run scripts/deploy.ts --network localhost

async function main() {

    const [owner] = await ethers.getSigners();

    const Factory = await ethers.getContractFactory("Accum", owner);
    console.log("Развертывание контракта...");

    const contr = await Factory.deploy(owner);
    await contr.waitForDeployment(); 
    console.log("Контракт развернут: ", contr.target);
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});