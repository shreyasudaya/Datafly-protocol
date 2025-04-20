package main

import (
	"fmt"
	"time"
)

func BaselinePerformanceTest() {
	fmt.Println("\n[Experiment] Baseline Performance Test")
	patient := NewPatient("P1")
	doctor := NewDoctor("D1")
	patient.GrantAccessToDoctor(doctor.ID)
	if !doctor.RequestAccess(patient.ID, patient.Contract) {
		fmt.Println("[Access Denied] Doctor is not authorized.")
		return
	}
	source := NewBlockchain("ChainA")
	dest := NewBlockchain("ChainB")
	source.AddNodes(7)
	dest.AddNodes(7)
	data := "Test Patient Data"
	source.StorePatientData(patient.ID, data)
	start := time.Now()
	DataFlyProtocol(patient, doctor, source, dest, data)
	fmt.Println("[METRICS] Latency:", time.Since(start))
	fmt.Println("[METRICS] Total Transactions:", source.TransactionCount+dest.TransactionCount)
}

func ThroughputTest(recordCount int) {
	fmt.Println("\n[Experiment] Throughput Test with", recordCount, "records")
	patient := NewPatient("P2")
	doctor := NewDoctor("D2")
	patient.GrantAccessToDoctor(doctor.ID)
	source := NewBlockchain("ChainA")
	dest := NewBlockchain("ChainB")
	source.AddNodes(7)
	dest.AddNodes(7)
	start := time.Now()
	for i := 0; i < recordCount; i++ {
		source.StorePatientData(fmt.Sprintf("%s-%d", patient.ID, i), fmt.Sprintf("Data %d", i))
		DataFlyProtocol(patient, doctor, source, dest, fmt.Sprintf("Data %d", i))
	}
	totalTime := time.Since(start)
	fmt.Println("[METRICS] Total Time:", totalTime)
	fmt.Printf("[METRICS] Throughput: %.2f records/sec\n", float64(recordCount)/totalTime.Seconds())
}

func FailureCaseSimulation() {
	fmt.Println("\n[Experiment] Failure Case Simulation")
	patient := NewPatient("P3")
	doctor := NewDoctor("D3")
	patient.GrantAccessToDoctor(doctor.ID)
	source := NewBlockchain("ChainA")
	dest := NewBlockchain("ChainB")
	// Only 2 nodes: not enough for >2/3 consensus
	source.AddNodes(2)
	dest.AddNodes(7)
	source.StorePatientData(patient.ID, "Test Failure")
	DataFlyProtocol(patient, doctor, source, dest, "Test Failure")
}

func MultiPartyConsensusTest(nodeCount int) {
	fmt.Println("\n[Experiment] Multi-Party Consensus Test with", nodeCount, "nodes")
	patient := NewPatient("P4")
	doctor := NewDoctor("D4")
	patient.GrantAccessToDoctor(doctor.ID)
	source := NewBlockchain("ChainA")
	dest := NewBlockchain("ChainB")
	source.AddNodes(nodeCount)
	dest.AddNodes(nodeCount)
	data := "Consensus Test Data"
	source.StorePatientData(patient.ID, data)
	DataFlyProtocol(patient, doctor, source, dest, data)
}
