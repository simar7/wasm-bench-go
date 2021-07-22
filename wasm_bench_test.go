package main

import (
	"bytes"
	_ "embed"
	"testing"

	"github.com/bytecodealliance/wasmtime-go"
	"github.com/wasmerio/wasmer-go/wasmer"

	"github.com/mathetake/gasm/wasi"
	gasm "github.com/mathetake/gasm/wasm"
)

//go:embed wasm/sum.wasm
var sumTinyGoWasm string

const wasmBytecode = `
	(module
	  (type (func (param i32 i32) (result i32)))
	  (func (type 0)
	    local.get 0
	    local.get 1
	    i32.add)
	  (export "sum" (func 0)))
`

func BenchmarkWasmtime(b *testing.B) {
	store := wasmtime.NewStore(wasmtime.NewEngine())
	wasm, _ := wasmtime.Wat2Wasm(wasmBytecode)

	module, _ := wasmtime.NewModule(store.Engine, wasm)
	instance, _ := wasmtime.NewInstance(store, module, []wasmtime.AsExtern{})

	sum := instance.GetExport(store, "sum").Func()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = sum.Call(store, 5, 37)
	}
}

func BenchmarkWasmer(b *testing.B) {
	wasmBytes := []byte(wasmBytecode)

	engine := wasmer.NewEngine()
	store := wasmer.NewStore(engine)

	module, _ := wasmer.NewModule(store, wasmBytes)

	importObject := wasmer.NewImportObject()
	instance, _ := wasmer.NewInstance(module, importObject)

	sum, _ := instance.Exports.GetFunction("sum")
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = sum(5, 37)
	}
}

func BenchmarkGasm(b *testing.B) {
	mod, _ := gasm.DecodeModule(bytes.NewBuffer([]byte(sumTinyGoWasm)))
	vm, _ := gasm.NewVM(mod, wasi.New().Modules())
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _, _ = vm.ExecExportedFunction("sum", uint64(5), uint64(37))
	}
}
