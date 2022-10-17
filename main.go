package main

import (
	"context"

	"github.com/bytecodealliance/wasmtime-go"
	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"
	"github.com/wasmerio/wasmer-go/wasmer"
)

func main() {
	panic("this is a benchmark project only")
}

func createWasmtime(in []byte) (func() error, error) {
	store := wasmtime.NewStore(wasmtime.NewEngine())

	module, err := wasmtime.NewModule(store.Engine, in)
	if err != nil {
		return nil, err
	}

	instance, err := wasmtime.NewInstance(store, module, []wasmtime.AsExtern{})
	if err != nil {
		return nil, err
	}

	sum := instance.GetExport(store, "sum").Func()

	return func() error {
		_, err = sum.Call(store, 5, 37)
		return nil
	}, nil
}

func createWasmer(in []byte) (func() error, error) {
	engine := wasmer.NewEngine()

	store := wasmer.NewStore(engine)

	module, err := wasmer.NewModule(store, in)
	if err != nil {
		return nil, err
	}

	importObject := wasmer.NewImportObject()
	instance, err := wasmer.NewInstance(module, importObject)
	if err != nil {
		return nil, err
	}

	sum, err := instance.Exports.GetFunction("sum")
	if err != nil {
		return nil, err
	}

	return func() error {
		_, err := sum(5, 37)

		return err
	}, nil
}

func createWazero(in []byte) (func() error, error) {
	ctx := context.Background()

	r := wazero.NewRuntime(ctx)

	_, err := wasi_snapshot_preview1.Instantiate(ctx, r)
	if err != nil {
		return nil, err
	}

	cc, err := r.CompileModule(ctx, in)
	if err != nil {
		return nil, err
	}

	m, err := r.InstantiateModule(ctx, cc, wazero.NewModuleConfig())
	if err != nil {
		return nil, err
	}

	fn := m.ExportedFunction("sum")

	return func() error {
		_, err := fn.Call(ctx, 5, 37)

		return err
	}, nil
}
