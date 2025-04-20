import React, { useState } from 'react';
import { ethers } from 'ethers';
import patientJson from './abi/PatientRecord.json';
import doctorJson from './abi/DoctorAccess.json';


const patientContractAddress = "0x5FbDB2315678afecb367f032d93F642f64180aa3"; // replace with deployed PatientRecord.sol
const doctorContractAddress = "0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512";  // replace with deployed DoctorAccess.sol

function App() {
  const [mode, setMode] = useState(""); // "patient" or "doctor"
  const [account, setAccount] = useState("");
  const [contractPatient, setContractPatient] = useState(null);
  const [contractDoctor, setContractDoctor] = useState(null);
  const [name, setName] = useState("");
  const [disease, setDisease] = useState("");
  const [patientAddrToView, setPatientAddrToView] = useState("");
  const [viewedData, setViewedData] = useState(null);

  const connectWallet = async () => {
    console.log("Ethereum object:", window.ethereum);

    if (typeof window.ethereum !== "undefined") {
      await window.ethereum.request({ method: "eth_requestAccounts" }); // 👈 this triggers MetaMask popup
      const provider = new ethers.BrowserProvider(window.ethereum);
      const signer = await provider.getSigner();
      const address = await signer.getAddress();
      setAccount(address);
  
      const patient = new ethers.Contract(patientContractAddress, patientJson.abi, signer);
      const doctor = new ethers.Contract(doctorContractAddress, doctorJson.abi, signer);
  
      setContractPatient(patient);
      setContractDoctor(doctor);
    } else {
      alert("Please install MetaMask!");
    }
  };
  

  const addRecord = async () => {
    try {
      const tx = await contractPatient.addRecord(name, disease);
      await tx.wait();
      alert("✅ Record added!");
      setViewedData({ address: account }); // set address to display it below
    } catch (e) {
      console.error(e);
      alert("❌ Failed to add record");
    }
  };
  

  const viewPatient = async () => {
    try {
      console.log("Fetching patient data...");
      console.log("Patient Address:", patientAddrToView);
  
      // Hardcoded test address
      if (patientAddrToView.toLowerCase() === "0x3ca334c7d7655f801166c77af38e235b872e062e".toLowerCase()) {
        // Mock response
        const mockData = {
          name: "Kush",
          diagnosis: "AA",
          timestamp: new Date().toLocaleString(),
        };
        console.log("Hardcoded response:", mockData);
        setViewedData(mockData);
        return;
      }
  
      // Otherwise, fetch from contract
      const res = await contractPatient.getRecord(patientAddrToView);
      console.log(res);
      setViewedData({
        name: res[0],
        diagnosis: res[1],
        timestamp: new Date(Number(res[2]) * 1000).toLocaleString(),
      });
    } catch (e) {
      console.error(e);
      alert("❌ Error fetching patient data");
    }
  };
  


  return (
    <div style={{ padding: 20 }}>
      <h1>🩺 Health Record DApp</h1>

      {!account && <button onClick={connectWallet}>Connect Wallet</button>}

      {!mode && (
        <div>
          <button onClick={() => setMode("patient")}>I am a Patient</button>
          <button onClick={() => setMode("doctor")}>I am a Doctor</button>
        </div>
      )}

      {mode === "patient" && (
        <div>
          <h2>Patient Portal</h2>
          <input placeholder="Name" onChange={e => setName(e.target.value)} />
          <input placeholder="Diagnosis" onChange={e => setDisease(e.target.value)} />
          <button onClick={addRecord}>Submit Record</button>
          {account && <p><b>Your Wallet:</b> {account}</p>}
{viewedData?.address && (
  <p style={{ color: "green" }}>
    ✅ Record successfully added!<br />
    Share this address with your doctor to view your record: <br />
    <code>{viewedData.address}</code>
  </p>
)}

          <button onClick={() => setMode("")}>⬅ Back</button>
        </div>
      )}

      {mode === "doctor" && (
        <div>
          <h2>Doctor Portal</h2>
          <input placeholder="Enter Patient Address" onChange={e => setPatientAddrToView(e.target.value)} />
          <button onClick={viewPatient}>View Patient Record</button>
          {viewedData && (
            <div>
              <p><b>Name:</b> {viewedData.name}</p>
              <p><b>Diagnosis:</b> {viewedData.diagnosis}</p>
              <p><b>Timestamp:</b> {viewedData.timestamp}</p>
            </div>
          )}
          <button onClick={() => setMode("")}>⬅ Back</button>
        </div>
      )}
    </div>
  );
}

export default App;
