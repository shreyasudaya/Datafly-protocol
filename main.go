package main

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"time"
)

var tpl = template.Must(template.ParseFiles("templates/index.html"))

type PageData struct {
	Output string
}

func main() {
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/run", runSimulation)

	fmt.Println("🌐 Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, nil)
}

func runSimulation(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Bad form", http.StatusBadRequest)
		return
	}
	data := r.FormValue("data")

	var buf bytes.Buffer

	// Simulate everything
	fmt.Fprintln(&buf, "--- DataFly Simulation Start ---")
	patient := NewPatient("Patient123")
	doctor := NewDoctor("DoctorABC")
	patient.GrantAccessToDoctor(doctor.ID)

	if !doctor.RequestAccess(patient.ID, patient.Contract) {
		fmt.Fprintln(&buf, "[Access Denied] Doctor is not authorized.")
		tpl.Execute(w, PageData{Output: buf.String()})
		return
	}

	sourceChain := NewBlockchain("HospitalChainA")
	destChain := NewBlockchain("HospitalChainB")
	sourceChain.AddNodes(7)
	destChain.AddNodes(7)

	sourceChain.StorePatientData(patient.ID, data)
	start := time.Now()
	CaptureOutput(&buf, func() {
		DataFlyProtocol(patient, doctor, sourceChain, destChain, data)
	})
	elapsed := time.Since(start)

	fmt.Fprintln(&buf, "--- Experiments ---")
	BaselinePerformanceTest()
	ThroughputTest(100)
	FailureCaseSimulation()
	MultiPartyConsensusTest(7)

	fmt.Fprintf(&buf, "[METRICS] Latency: %v\n", elapsed)
	fmt.Fprintf(&buf, "[METRICS] Total Transactions: %d\n",
		sourceChain.TransactionCount+destChain.TransactionCount)

	tpl.Execute(w, PageData{Output: buf.String()})
}
