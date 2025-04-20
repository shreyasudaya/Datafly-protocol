// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract PatientRecord {
    struct Patient {
        string name;
        string diagnosis;
        uint256 timestamp;
    }

    mapping(address => Patient) private records;

    function addRecord(string memory _name, string memory _diagnosis) external {
        records[msg.sender] = Patient(_name, _diagnosis, block.timestamp);
    }

    function getRecord(address _addr) external view returns (string memory, string memory, uint256) {
        require(bytes(records[_addr].name).length > 0, "No record found");
        Patient memory p = records[_addr];
        return (p.name, p.diagnosis, p.timestamp);
    }
}
