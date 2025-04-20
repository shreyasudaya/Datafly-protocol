package main

import (
	"crypto/ecdsa"
	"fmt"
	"sync"
)

type AccessContract struct {
	Authorizations map[string]bool
	PatientPubKey  ecdsa.PublicKey
	PatientID      string
	Mux            sync.Mutex
}

func NewAccessContract(patientID string, pubKey ecdsa.PublicKey) *AccessContract {
	return &AccessContract{
		Authorizations: make(map[string]bool),
		PatientPubKey:  pubKey,
		PatientID:      patientID,
	}
}

func (ac *AccessContract) AuthorizeDoctor(doctorID string) {
	ac.Mux.Lock()
	defer ac.Mux.Unlock()
	ac.Authorizations[doctorID] = true
	fmt.Println("[AccessControl] Doctor", doctorID, "authorized by", ac.PatientID)
}

func (ac *AccessContract) IsAuthorized(doctorID string) bool {
	ac.Mux.Lock()
	defer ac.Mux.Unlock()
	return ac.Authorizations[doctorID]
}
