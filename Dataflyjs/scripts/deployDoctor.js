const hre = require("hardhat");

async function main() {
  const patientContractAddress = "0x5FbDB2315678afecb367f032d93F642f64180aa3";

  const DoctorAccess = await hre.ethers.getContractFactory("DoctorAccess");
  const doctorAccess = await DoctorAccess.deploy(patientContractAddress);

  await doctorAccess.waitForDeployment(); // ✅

  console.log("✅ DoctorAccess deployed at:", await doctorAccess.getAddress());
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
