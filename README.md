# Wasm crypto sdk for broom

The transaction signing and hashing logic must be duplicated in the browser implementation. Instead of reimplementing in js and stringing together third party libraries it is better if we use the exact same function bindings for wasm. This will streamline implementation of other broom browser based utilities such as mining and faucet dispensing.

## Exposed functions are:
`wasmGenerateKeypair`

`wasmGetSignature`

`wasmHashArgon`

`wasmGetTransactionSig`
