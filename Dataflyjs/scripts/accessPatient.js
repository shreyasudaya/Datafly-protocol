const hre = require("hardhat");

async function main() {
    const contract = await hre.ethers.getContractAt("DoctorAccess", "0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512");
    const patientAddress = "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266";
    const data = await contract.viewPatient(patientAddress);
    console.log("Patient Data:", data);
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
