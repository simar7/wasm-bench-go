# wasm-bench-go
wasm runtime benchmarks for go


```
make bench | prettybench -no-passthrough
benchmark               iter       time/iter   bytes alloc         allocs
---------               ----       ---------   -----------         ------
BenchmarkWasmtime-8   100000   3488.00 ns/op      311 B/op   20 allocs/op
BenchmarkWasmtime-8   100000   3614.00 ns/op      311 B/op   20 allocs/op
BenchmarkWasmtime-8   100000   3543.00 ns/op      311 B/op   20 allocs/op
BenchmarkWasmer-8     100000   1997.00 ns/op      208 B/op   12 allocs/op
BenchmarkWasmer-8     100000   1990.00 ns/op      208 B/op   12 allocs/op
BenchmarkWasmer-8     100000   1950.00 ns/op      208 B/op   12 allocs/op
BenchmarkGasm-8       100000    331.50 ns/op      288 B/op    9 allocs/op
BenchmarkGasm-8       100000    334.40 ns/op      288 B/op    9 allocs/op
BenchmarkGasm-8       100000    332.20 ns/op      288 B/op    9 allocs/op

```
