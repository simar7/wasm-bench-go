# wasm-bench-go
wasm runtime benchmarks for go


```
make bench | prettybench -no-passthrough
benchmark                iter       time/iter   bytes alloc         allocs
---------                ----       ---------   -----------         ------
BenchmarkWasmtime-12   100000   2929.00 ns/op      304 B/op   20 allocs/op
BenchmarkWasmtime-12   100000   2643.00 ns/op      312 B/op   21 allocs/op
BenchmarkWasmtime-12   100000   2626.00 ns/op      312 B/op   21 allocs/op
BenchmarkWasmer-12     100000   1524.00 ns/op      216 B/op   13 allocs/op
BenchmarkWasmer-12     100000   1478.00 ns/op      208 B/op   12 allocs/op
BenchmarkWasmer-12     100000   1546.00 ns/op      208 B/op   12 allocs/op
BenchmarkWazero-12     100000     52.12 ns/op       16 B/op    1 allocs/op
BenchmarkWazero-12     100000     53.93 ns/op       16 B/op    1 allocs/op
BenchmarkWazero-12     100000     53.91 ns/op       16 B/op    1 allocs/op
```
