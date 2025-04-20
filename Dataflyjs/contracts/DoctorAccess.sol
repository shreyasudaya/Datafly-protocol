// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

interface IPatientRecord {
    struct Record {
        string name;
        string diagnosis;
        uint256 timestamp;
    }

    function getRecords(address patient) external view returns (Record memory);
}

contract DoctorAccess {
    address public patientContract;

    constructor(address _patientContract) {
        patientContract = _patientContract;
    }

    // Doctor views patient record using address
    function viewPatient(address patient) external view returns (string memory, string memory, uint256) {
        IPatientRecord.Record memory record = IPatientRecord(patientContract).getRecords(patient);
        require(bytes(record.name).length > 0, "No data found");
        return (record.name, record.diagnosis, record.timestamp);
    }
}
