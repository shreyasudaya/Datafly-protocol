package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
)

type Patient struct {
	ID       string
	PubKey   ecdsa.PublicKey
	PrivKey  ecdsa.PrivateKey
	Contract *AccessContract
}

func NewPatient(id string) *Patient {
	priv, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	pub := priv.PublicKey
	contract := NewAccessContract(id, pub)
	return &Patient{
		ID:       id,
		PubKey:   pub,
		PrivKey:  *priv,
		Contract: contract,
	}
}

func (p *Patient) GrantAccessToDoctor(doctorID string) {
	fmt.Println("[Patient", p.ID, "] Granting access to Doctor", doctorID)
	p.Contract.AuthorizeDoctor(doctorID)
}
