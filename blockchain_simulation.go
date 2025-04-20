package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"
	"sync"
)

type Blockchain struct {
	Name             string
	Nodes            []Node
	Ledger           map[string]string
	TransactionCount int
	Mux              sync.Mutex
}

type Node struct {
	ID      string
	PubKey  ecdsa.PublicKey
	PrivKey ecdsa.PrivateKey
}

func NewBlockchain(name string) *Blockchain {
	return &Blockchain{
		Name:   name,
		Ledger: make(map[string]string),
	}
}

func (bc *Blockchain) AddNodes(n int) {
	for i := 0; i < n; i++ {
		priv, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
		node := Node{
			ID:      fmt.Sprintf("%s-Node%d", bc.Name, i),
			PrivKey: *priv,
			PubKey:  priv.PublicKey,
		}
		bc.Nodes = append(bc.Nodes, node)
	}
}

func (bc *Blockchain) StorePatientData(patientID, data string) {
	hash := sha256.Sum256([]byte(data))
	bc.Ledger[patientID] = hex.EncodeToString(hash[:])
	fmt.Println("[", bc.Name, "] Stored hash of patient data.")
	bc.TransactionCount++
}

func (bc *Blockchain) SimulatePegConsensus(data string) (bool, []byte, []Signature) {
	hash := sha256.Sum256([]byte(data))
	sigs := []Signature{}
	required := (len(bc.Nodes) * 2 / 3) + 1
	count := 0
	for _, node := range bc.Nodes {
		r, s, _ := ecdsa.Sign(rand.Reader, &node.PrivKey, hash[:])
		if ecdsa.Verify(&node.PubKey, hash[:], r, s) {
			sigs = append(sigs, Signature{R: r, S: s, PubKey: node.PubKey})
			count++
			if count >= required {
				break
			}
		}
	}
	return count >= required, hash[:], sigs
}

type Signature struct {
	R, S   *big.Int
	PubKey ecdsa.PublicKey
}

func VerifyPegProof(data string, hash []byte, sigs []Signature) bool {
	valid := 0
	for _, sig := range sigs {
		if ecdsa.Verify(&sig.PubKey, hash, sig.R, sig.S) {
			valid++
		}
	}
	return valid >= (len(sigs) * 2 / 3)
}
