package main

import (
	_ "embed"
	"testing"
)

//go:embed testdata/tinygo/sum.wasm
var sumTinyGoWasm []byte

func BenchmarkWasmtime(b *testing.B) {
	fn, err := createWasmtime(sumTinyGoWasm)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = fn()
	}
}

func BenchmarkWasmer(b *testing.B) {
	fn, err := createWasmer(sumTinyGoWasm)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = fn()
	}
}

func BenchmarkWazero(b *testing.B) {
	fn, err := createWazero(sumTinyGoWasm)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = fn()
	}
}
