package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

func CaptureOutput(w io.Writer, fn func()) {
	stdout := os.Stdout
	r, wPipe, _ := os.Pipe()
	os.Stdout = wPipe

	fn()

	wPipe.Close()
	os.Stdout = stdout

	var buf bytes.Buffer
	io.Copy(&buf, r)
	fmt.Fprint(w, buf.String())
}

func DataFlyProtocol(patient *Patient, doctor *Doctor, source *Blockchain, dest *Blockchain, data string) {
	fmt.Println("[DataFly] Starting data migration...")

	encKey := sha256.Sum256(append(doctor.PubKey.X.Bytes(), doctor.PubKey.Y.Bytes()...))
	encrypted := EncryptData(data, encKey[:])

	consensusOk, hash, sigs := source.SimulatePegConsensus(string(encrypted))
	if !consensusOk {
		fmt.Println("[DataFly] Source consensus failed.")
		return
	}

	ok := VerifyPegProof(string(encrypted), hash, sigs)
	if !ok {
		fmt.Println("[DataFly] Destination rejected proof.")
		return
	}

	dest.Ledger[patient.ID] = hex.EncodeToString(encrypted)
	dest.TransactionCount++

	decrypted := DecryptData(encrypted, encKey[:])
	fmt.Println("[DataFly] Data decrypted by Doctor:", decrypted)
}
