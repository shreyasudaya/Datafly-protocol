package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
)

type Doctor struct {
	ID      string
	PubKey  ecdsa.PublicKey
	PrivKey ecdsa.PrivateKey
}

func NewDoctor(id string) *Doctor {
	priv, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	return &Doctor{
		ID:      id,
		PubKey:  priv.PublicKey,
		PrivKey: *priv,
	}
}

func (d *Doctor) RequestAccess(patientID string, ac *AccessContract) bool {
	fmt.Println("[Doctor", d.ID, "] Requesting access to", patientID)
	return ac.IsAuthorized(d.ID)
}
