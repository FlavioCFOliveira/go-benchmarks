# go-benchmarks

```shell
go test -benchmem -run=^$ -bench ^Benchmark -benchtime 30s github.com/FlavioCFOliveira/go-benchmarks
goos: windows
goarch: amd64
pkg: github.com/FlavioCFOliveira/go-benchmarks
cpu: Intel(R) Core(TM) i7-10510U CPU @ 1.80GHz
BenchmarkJsonDecode-8              36709           1002144 ns/op          158176 B/op       2248 allocs/op
BenchmarkJsonUnmarshal-8           36458           1065536 ns/op          158473 B/op       2255 allocs/op
PASS
ok      github.com/FlavioCFOliveira/go-benchmarks       95.720s
```
