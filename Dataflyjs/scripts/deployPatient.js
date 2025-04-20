const hre = require("hardhat");

async function main() {
  const PatientRecord = await hre.ethers.getContractFactory("PatientRecord");
  const patientRecord = await PatientRecord.deploy();

  await patientRecord.waitForDeployment(); // ✅ correct in ethers v6+

  console.log("✅ PatientRecord deployed at:", await patientRecord.getAddress());
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
