bench: wasm-bundle
	go test -bench=. -benchmem -count=3 -benchtime=100000x .

wasm-bundle:
	tinygo build -o wasm/sum.wasm -target wasm wasm/sum.go

clean:
	rm -rf wasm/*.wasm

