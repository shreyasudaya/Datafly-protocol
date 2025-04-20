const hre = require("hardhat");

async function main() {
    const contract = await hre.ethers.getContractAt("PatientRecord", "0x5FbDB2315678afecb367f032d93F642f64180aa3");
    const tx = await contract.addRecord("Kush", "Asthma");
    await tx.wait();
    console.log("Record added");
}

main();
