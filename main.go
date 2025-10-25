//go:build js && wasm
// +build js,wasm

package main

import (
	"fmt"
	"syscall/js"
)

func main() {

	fmt.Println("wasm bindings yay!")

	js.Global().Set("wasmGenerateKeypair", js.FuncOf(wasmGenerateKeypair))
	js.Global().Set("wasmGetSignature", js.FuncOf(wasmGetSignature))
	js.Global().Set("wasmHashArgon", js.FuncOf(wasmHashArgon))
	js.Global().Set("wasmGetTransactionSig", js.FuncOf(wasmGetTransactionSig))

	select {}
}
