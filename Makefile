bench: wasm-bundle
	GOGC=off go test -bench=. -benchmem -count=3 -benchtime=100000x .

wasm-bundle:
	go generate ./...

clean:
	rm -rf wasm/*.wasm

