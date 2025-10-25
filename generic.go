//go:build js && wasm
// +build js,wasm

package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/json"

	nodecrypto "github.com/canavan-a/broom/node/crypto"
	netnode "github.com/canavan-a/broom/node/netnode"
)

func Generic_GenerateKeypair() any {
	priv, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)

	address, _ := nodecrypto.GenerateAddress(priv.PublicKey)
	privateKey := nodecrypto.GeneratePrivateKeyText(priv)

	return map[string]any{
		"private": privateKey,
		"public":  address,
	}
}

func Generic_GetSignature(privateKey string, data []byte) any {

	output := make(map[string]any)
	output["error"] = ""

	private, err := nodecrypto.ParsePrivateKey(privateKey)
	if err != nil {
		output["error"] = err.Error()
		return output
	}
	sig := nodecrypto.Sign(data, private)
	output["sig"] = sig
	return output
}

func Generic_HashArgon(data []byte) any {
	output := make(map[string]any)
	output["error"] = ""

	hash := nodecrypto.HashArgon2d(data)

	output["hash"] = hash
	return output
}

func Gereneric_GetTransactionSig(privateKey string, stringifiedTxn string) any {

	output := make(map[string]any)
	var txn netnode.Transaction

	err := json.Unmarshal([]byte(stringifiedTxn), &txn)
	if err != nil {
		output["error"] = err.Error()
		return output
	}

	err = txn.Sign(privateKey)
	if err != nil {
		output["error"] = err.Error()
		return output
	}

	output["sig"] = txn.Sig

	return output

}
