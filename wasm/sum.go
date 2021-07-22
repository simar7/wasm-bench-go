//go:generate tinygo build -opt=s -o sum.wasm -target wasm sum.go
package main

func main() {}

//export sum
func sum(in1, in2 int) int {
	return in1 + in2
}
