package main

import (
	_ "embed"
	"testing"

	"github.com/bytecodealliance/wasmtime-go"
	"github.com/wasmerio/wasmer-go/wasmer"
)

const wasmBytecode = `
	(module
	  (type (func (param i32 i32) (result i32)))
	  (func (type 0)
	    local.get 0
	    local.get 1
	    i32.add)
	  (export "sum" (func 0)))
`

func BenchmarkWasmer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		wasmBytes := []byte(wasmBytecode)

		engine := wasmer.NewEngine()
		store := wasmer.NewStore(engine)

		module, _ := wasmer.NewModule(store, wasmBytes)

		importObject := wasmer.NewImportObject()
		instance, _ := wasmer.NewInstance(module, importObject)

		sum, _ := instance.Exports.GetFunction("sum")
		b.StartTimer()

		_, _ = sum(5, 37)
	}
}

func BenchmarkWasmtime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		store := wasmtime.NewStore(wasmtime.NewEngine())
		wasm, _ := wasmtime.Wat2Wasm(wasmBytecode)

		module, _ := wasmtime.NewModule(store.Engine, wasm)
		instance, _ := wasmtime.NewInstance(store, module, []wasmtime.AsExtern{})

		sum := instance.GetExport(store, "sum").Func()
		b.StartTimer()

		_, _ = sum.Call(store, 5, 37)
	}
}
