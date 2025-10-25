//go:build js && wasm
// +build js,wasm

package main

import "syscall/js"

func wasmGenerateKeypair(this js.Value, args []js.Value) any {
	output := Generic_GenerateKeypair()

	return output

}

func wasmGetSignature(this js.Value, args []js.Value) any {

	privateKey := args[0]

	jsData := args[1]
	data := make([]byte, jsData.Get("length").Int())

	js.CopyBytesToGo(data, jsData)
	return Generic_GetSignature(privateKey.String(), data)
}

func wasmHashArgon(this js.Value, args []js.Value) any {
	jsData := args[1]
	data := make([]byte, jsData.Get("length").Int())

	js.CopyBytesToGo(data, jsData)

	return Generic_HashArgon(data)
}

func wasmGetTransactionSig(this js.Value, args []js.Value) any {
	privateKey := args[0]
	stringifiedData := args[1]

	return Gereneric_GetTransactionSig(privateKey.String(), stringifiedData.String())

}
