package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("--- DataFly Simulation Start ---")

	// Initialize actors
	patient := NewPatient("Patient123")
	doctor := NewDoctor("DoctorABC")

	// Patient grants access
	patient.GrantAccessToDoctor(doctor.ID)

	// Doctor requests access
	if !doctor.RequestAccess(patient.ID, patient.Contract) {
		fmt.Println("[Access Denied] Doctor is not authorized.")
		return
	}

	// Initialize chains
	sourceChain := NewBlockchain("HospitalChainA")
	destChain := NewBlockchain("HospitalChainB")

	// Add nodes to chains
	sourceChain.AddNodes(7) // 7 nodes for >2/3 consensus (5 needed)
	destChain.AddNodes(7)

	// Simulate data storage and migration
	dataToMigrate := "Patient record: Diabetic - Meds: Metformin"
	sourceChain.StorePatientData(patient.ID, dataToMigrate)

	start := time.Now()
	DataFlyProtocol(patient, doctor, sourceChain, destChain, dataToMigrate)
	elapsed := time.Since(start)

	fmt.Println("[METRICS] Latency:", elapsed)
	fmt.Println("[METRICS] Total Transactions:",
		sourceChain.TransactionCount+destChain.TransactionCount)
}
